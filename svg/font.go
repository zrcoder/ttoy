package svg

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

// woffToTTF returns the sfnt (TrueType) bytes inside a WOFF1 container,
// which is what golang.org/x/image/font/sfnt reads: diagrams ship their
// fonts as WOFF, base64 encoded in their style sheet.
func woffToTTF(src []byte) ([]byte, error) {
	const headerLen = 44
	const entryLen = 20
	if len(src) < headerLen {
		return nil, errors.New("woff: truncated header")
	}
	switch string(src[:4]) {
	case "wOFF":
	case "wOF2":
		return nil, errors.New("woff: WOFF2 needs a Brotli decoder")
	default:
		return nil, errors.New("woff: not a WOFF file")
	}
	num := int(binary.BigEndian.Uint16(src[12:14]))
	total := int(binary.BigEndian.Uint32(src[16:20]))
	if num == 0 || headerLen+entryLen*num > len(src) {
		return nil, errors.New("woff: truncated table directory")
	}

	type table struct {
		tag  [4]byte
		cs   uint32
		data []byte
	}
	tables := make([]table, 0, num)
	for i := range num {
		e := src[headerLen+entryLen*i : headerLen+entryLen*(i+1)]
		off := binary.BigEndian.Uint32(e[4:8])
		comp := binary.BigEndian.Uint32(e[8:12])
		orig := binary.BigEndian.Uint32(e[12:16])
		if uint64(off)+uint64(comp) > uint64(len(src)) {
			return nil, errors.New("woff: table outside the file")
		}
		raw := src[off : off+comp]
		var data []byte
		if comp < orig {
			zr, err := zlib.NewReader(bytes.NewReader(raw))
			if err != nil {
				return nil, fmt.Errorf("woff: %s: %w", e[0:4], err)
			}
			data, err = io.ReadAll(zr)
			zr.Close()
			if err != nil {
				return nil, fmt.Errorf("woff: %s: %w", e[0:4], err)
			}
			if uint32(len(data)) != orig {
				return nil, fmt.Errorf("woff: %s: inflated %d bytes, want %d", e[0:4], len(data), orig)
			}
		} else {
			data = append([]byte(nil), raw...)
		}
		var t table
		copy(t.tag[:], e[0:4])
		t.cs = binary.BigEndian.Uint32(e[16:20])
		t.data = data
		tables = append(tables, t)
	}

	// The sfnt header, the table directory, then the tables it points at.
	var b bytes.Buffer
	if total > 0 {
		b.Grow(total)
	}
	b.Write(src[4:8]) // flavor: the sfnt version the tables came in
	entrySel := uint16(0)
	for 1<<(entrySel+1) <= uint32(num) {
		entrySel++
	}
	searchRange := uint16(16) << entrySel
	binary.Write(&b, binary.BigEndian, uint16(num))
	binary.Write(&b, binary.BigEndian, searchRange)
	binary.Write(&b, binary.BigEndian, entrySel)
	binary.Write(&b, binary.BigEndian, uint16(num*16)-searchRange)

	off := 12 + 16*num
	for _, t := range tables {
		b.Write(t.tag[:])
		binary.Write(&b, binary.BigEndian, t.cs)
		binary.Write(&b, binary.BigEndian, uint32(off))
		binary.Write(&b, binary.BigEndian, uint32(len(t.data)))
		off += (len(t.data) + 3) &^ 3
	}
	for _, t := range tables {
		b.Write(t.data)
		for b.Len()%4 != 0 {
			b.WriteByte(0)
		}
	}
	return b.Bytes(), nil
}

// dataURIBytes returns what a data:image/...;base64,... URL holds, behind
// the url(...) and quotes a font source wraps it in.
func dataURIBytes(s string) ([]byte, error) {
	s = strings.TrimSpace(s)
	if i := strings.Index(s, "url("); i >= 0 {
		s = s[i+4:]
	}
	s = strings.Trim(strings.TrimSpace(s), `"'`)
	if j := strings.IndexByte(s, ')'); j >= 0 {
		s = s[:j]
	}
	s = strings.Trim(strings.TrimSpace(s), `"'`)
	const marker = "base64,"
	i := strings.Index(s, marker)
	if i < 0 {
		return nil, fmt.Errorf("font source is not base64: %q", trunc(s, 40))
	}
	return base64.StdEncoding.DecodeString(strings.TrimSpace(s[i+len(marker):]))
}

// trunc shortens s for an error message.
func trunc(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// svgFonts holds the faces one SVG's fonts name, and system faces for the
// characters those do not cover, as CJK labels in a Latin font.
type svgFonts struct {
	faces map[string]*sfnt.Font // by the font-family the style sheet names
	def   *sfnt.Font            // the regular face, for text naming none

	once     sync.Once
	fallback []*sfnt.Font
}

// newSVGFonts parses each font of a document, by the family name its
// @font-face declares.
func newSVGFonts(fonts map[string][]byte) *svgFonts {
	fs := &svgFonts{faces: make(map[string]*sfnt.Font, len(fonts))}
	for name, ttf := range fonts {
		f, err := sfnt.Parse(ttf)
		if err != nil {
			continue
		}
		key := strings.ToLower(name)
		fs.faces[key] = f
		if strings.HasSuffix(key, "-font-regular") {
			fs.def = f
		}
	}
	if fs.def == nil {
		for _, f := range fs.faces {
			fs.def = f
			break
		}
	}
	return fs
}

// face returns the face of a family name, or the document's regular one.
func (fs *svgFonts) face(family string) *sfnt.Font {
	if f := fs.faces[strings.ToLower(strings.Trim(family, `"'`))]; f != nil {
		return f
	}
	return fs.def
}

// systemFallback lists fonts to reach for when a document's own fonts have
// no glyph for a character, as Source Sans Pro has none for Chinese, and
// for a document that carries no fonts of its own. It spans the three
// systems ttoy runs on: a path of another system fails to read and is
// skipped, so where a machine keeps its fonts decides what shows.
var systemFallback = []string{
	// macOS
	"/System/Library/Fonts/Hiragino Sans GB.ttc",
	"/System/Library/Fonts/STHeiti Medium.ttc",
	"/System/Library/Fonts/Supplemental/Songti.ttc",
	"/System/Library/Fonts/Supplemental/Arial Unicode.ttf",
	"/System/Library/Fonts/Helvetica.ttc",
	"/System/Library/Fonts/Supplemental/Arial.ttf",
	// Linux, whose directories differ by distribution
	"/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc",
	"/usr/share/fonts/truetype/noto/NotoSansCJK-Regular.ttc",
	"/usr/share/fonts/noto-cjk/NotoSansCJK-Regular.ttc",
	"/usr/share/fonts/google-noto-cjk/NotoSansCJK-Regular.ttc",
	"/usr/share/fonts/truetype/wqy/wqy-zenhei.ttc",
	"/usr/share/fonts/truetype/wqy/wqy-microhei.ttc",
	"/usr/share/fonts/truetype/arphic/uming.ttc",
	"/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf",
	"/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf",
	"/usr/share/fonts/truetype/noto/NotoSans-Regular.ttf",
	// Windows
	windowsFont("msyh.ttc"),
	windowsFont("simhei.ttf"),
	windowsFont("simsun.ttc"),
	windowsFont("arial.ttf"),
	windowsFont("segoeui.ttf"),
}

// windowsFont is a font of the Windows install, whose directory is named
// by WINDIR and is not always on C:.
func windowsFont(name string) string {
	if dir := os.Getenv("WINDIR"); dir != "" {
		return filepath.Join(dir, "Fonts", name)
	}
	return filepath.Join(`C:\Windows`, "Fonts", name)
}

// fallbacks loads the system fonts once, on the first character the
// document's fonts cannot draw.
func (fs *svgFonts) fallbacks() []*sfnt.Font {
	fs.once.Do(func() {
		for _, path := range systemFallback {
			data, err := os.ReadFile(path)
			if err != nil {
				continue
			}
			if f, err := sfnt.Parse(data); err == nil {
				fs.fallback = append(fs.fallback, f)
				continue
			}
			c, err := sfnt.ParseCollection(data)
			if err != nil {
				continue
			}
			for i := 0; i < c.NumFonts(); i++ {
				if f, err := c.Font(i); err == nil {
					fs.fallback = append(fs.fallback, f)
				}
			}
		}
	})
	return fs.fallback
}

// glyphFor returns the face drawing r: the primary one when it has the
// character, a system font when not, or nil when nothing has it.
func (fs *svgFonts) glyphFor(primary *sfnt.Font, buf *sfnt.Buffer, r rune) (*sfnt.Font, sfnt.GlyphIndex) {
	if primary != nil {
		if idx, err := primary.GlyphIndex(buf, r); err == nil && idx != 0 {
			return primary, idx
		}
	}
	for _, f := range fs.fallbacks() {
		if idx, err := f.GlyphIndex(buf, r); err == nil && idx != 0 {
			return f, idx
		}
	}
	return nil, 0
}

// shapedGlyph is one character's glyph at a pen position along a line.
type shapedGlyph struct {
	face *sfnt.Font
	idx  sfnt.GlyphIndex
	size float64 // pixels per em
	x    float64 // pen position from the line's origin
}

// shape lays s out along the x axis with primary at size pixels per em,
// and returns its width in pixels. Fallback fonts take over where the
// primary has no glyph.
func (fs *svgFonts) shape(primary *sfnt.Font, size float64, s string) (float64, []shapedGlyph) {
	if size <= 0 {
		size = 16
	}
	ppem := fixed.Int26_6(size * 64)
	var buf sfnt.Buffer
	var glyphs []shapedGlyph
	var x float64
	var prevFace *sfnt.Font
	var prevIdx sfnt.GlyphIndex
	for _, r := range s {
		face, idx := fs.glyphFor(primary, &buf, r)
		if face == nil {
			continue // no font draws it: leave it out rather than a box
		}
		if prevFace == face {
			if kern, err := face.Kern(&buf, prevIdx, idx, ppem, font.HintingNone); err == nil {
				x += float64(kern) / 64
			}
		}
		adv, err := face.GlyphAdvance(&buf, idx, ppem, font.HintingNone)
		if err != nil {
			continue
		}
		glyphs = append(glyphs, shapedGlyph{face: face, idx: idx, size: size, x: x})
		x += float64(adv) / 64
		prevFace, prevIdx = face, idx
	}
	return x, glyphs
}

// appendPath writes the glyphs' outlines into path data, placed with their
// baseline at (ox, oy).
func appendPath(d *strings.Builder, glyphs []shapedGlyph, ox, oy float64) {
	var buf sfnt.Buffer
	for _, g := range glyphs {
		segs, err := g.face.LoadGlyph(&buf, g.idx, fixed.Int26_6(g.size*64), nil)
		if err != nil || len(segs) == 0 {
			continue
		}
		if d.Len() > 0 {
			d.WriteByte('Z')
		}
		for _, s := range segs {
			p := func(i int) (float64, float64) {
				return ox + g.x + float64(s.Args[i].X)/64, oy + float64(s.Args[i].Y)/64
			}
			switch s.Op {
			case sfnt.SegmentOpMoveTo:
				x, y := p(0)
				fmt.Fprintf(d, "M%s%s", coord(x), coord(y))
			case sfnt.SegmentOpLineTo:
				x, y := p(0)
				fmt.Fprintf(d, "L%s%s", coord(x), coord(y))
			case sfnt.SegmentOpQuadTo:
				x1, y1 := p(0)
				x, y := p(1)
				fmt.Fprintf(d, "Q%s%s%s%s", coord(x1), coord(y1), coord(x), coord(y))
			case sfnt.SegmentOpCubeTo:
				x1, y1 := p(0)
				x2, y2 := p(1)
				x, y := p(2)
				fmt.Fprintf(d, "C%s%s%s%s%s%s", coord(x1), coord(y1), coord(x2), coord(y2), coord(x), coord(y))
			}
		}
	}
	if d.Len() > 0 {
		d.WriteByte('Z')
	}
}

// coord formats a path coordinate with the separator SVG path data needs.
func coord(v float64) string {
	return fmt.Sprintf("%.2f ", v)
}

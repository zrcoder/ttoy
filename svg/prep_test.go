package svg

import (
	"image"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/cdor"
)

// d2SVG renders a d2 diagram the way the plot panel does.
func d2SVG(t *testing.T, src string) []byte {
	t.Helper()
	c := cdor.Ctx()
	c.D2(src)
	c.Fill("transparent")
	c.Theme(200)
	data, err := c.Gen()
	if err != nil {
		t.Fatal(err)
	}
	return data
}

// TestPrepareSVG reads a diagram through the whole path: the labels and
// arrows MyGo leaves out are rewritten, and what comes out draws with the
// colors of the theme the document asks for.
func TestPrepareSVG(t *testing.T) {
	in := d2SVG(t, `"a" -> "b": hello
c -> d`)
	prepared, err := Prepare(in, true)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	for _, gone := range []string{"<text", "<tspan", "<marker", "marker-end="} {
		if strings.Contains(string(prepared), gone) {
			t.Errorf("prepared document still holds %s", gone)
		}
	}
	assertSimpleSelectors(t, string(prepared))

	s, err := ui.ParseSVG(prepared)
	if err != nil {
		t.Fatalf("ui.ParseSVG: %v", err)
	}
	img := ui.Render(func(c *ui.Context) { ui.Image(c, s) }, 369, 455, 1)
	// The boxes of the theme in the dark branch, #313244, and its light
	// purple connections, #CBA6F7: both only reach the drawing if the
	// color scheme query was unwrapped and the rules matched.
	if n := countColor(img, 0x31, 0x32, 0x44); n < 100 {
		t.Errorf("boxes filled %d pixels with their color, want 100", n)
	}
	if n := countColor(img, 0xCB, 0xA6, 0xF7); n < 10 {
		t.Errorf("connections drew %d pixels with their color, want 10", n)
	}
}

// TestPrepareSVGColorScheme reads the one query a diagram carries, for
// the system's color scheme, into the branch of the theme it is in.
func TestPrepareSVGColorScheme(t *testing.T) {
	src := `<svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 40 40">
<style type="text/css"><![CDATA[
.box { fill: #EEEEEE; }
@media screen and (prefers-color-scheme:dark) {
	.box { fill: #111111; }
}
]]></style>
<rect class="box" x="0" y="0" width="40" height="40"/>
</svg>`
	for _, tc := range []struct {
		dark bool
		want [3]uint8
	}{{false, [3]uint8{0xEE, 0xEE, 0xEE}}, {true, [3]uint8{0x11, 0x11, 0x11}}} {
		prepared, err := Prepare([]byte(src), tc.dark)
		if err != nil {
			t.Fatalf("Prepare: %v", err)
		}
		assertSimpleSelectors(t, string(prepared))
		s, err := ui.ParseSVG(prepared)
		if err != nil {
			t.Fatalf("ui.ParseSVG: %v", err)
		}
		img := ui.Render(func(c *ui.Context) { ui.Image(c, s) }, 40, 40, 1)
		if n := countColor(img, tc.want[0], tc.want[1], tc.want[2]); n < 1000 {
			t.Errorf("dark=%v: %d pixels of #%02X%02X%02X, want 1000", tc.dark, n, tc.want[0], tc.want[1], tc.want[2])
		}
	}
}

// TestPrepareSVGTspanLines reads a label whose lines are tspans and whose
// size an ancestor sets, as a code block's labels are.
func TestPrepareSVGTspanLines(t *testing.T) {
	src := `<svg xmlns="http://www.w3.org/2000/svg" width="120" height="60" viewBox="0 0 120 60">
<style type="text/css"><![CDATA[
.wrap .row { fill: #123456; }
]]></style>
<g class="wrap" style="font-size:12"><text class="row" x="10" y="20" style="text-anchor:middle">one<tspan x="10" dy="1.5em">two</tspan></text></g>
</svg>`
	prepared, err := Prepare([]byte(src), false)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	out := string(prepared)
	assertSimpleSelectors(t, out)
	if strings.Contains(out, "<text") || strings.Contains(out, "<tspan") {
		t.Fatal("the label kept its text elements")
	}
	if !strings.Contains(out, `.row {`) {
		t.Error("the class rule was dropped")
	}
	// One path per line: the first sits on y=20, the second an em, 12px,
	// below it, which the bottoms of glyphs without descenders show.
	ds := regexp.MustCompile(`d="([^"]+)"`).FindAllStringSubmatch(out, -1)
	if len(ds) != 2 {
		t.Fatalf("drew %d paths, want one per line", len(ds))
	}
	for i, want := range []float64{20, 38} {
		var bottom float64
		for _, y := range pathYs(ds[i][1]) {
			if y > bottom {
				bottom = y
			}
		}
		if bottom < want-1 || bottom > want+1 {
			t.Errorf("line %d ends at y=%.2f, want around %.0f", i+1, bottom, want)
		}
	}
}

// pathYs reads the y coordinates of the points a path is drawn through,
// which come after each x.
func pathYs(d string) []float64 {
	var ys []float64
	for i, s := range regexp.MustCompile(`-?\d+(?:\.\d+)?`).FindAllString(d, -1) {
		if i%2 == 0 {
			continue
		}
		if v, err := strconv.ParseFloat(s, 64); err == nil {
			ys = append(ys, v)
		}
	}
	return ys
}

// TestPrepareSVGMarker reads a path with an arrowhead, which MyGo draws
// no markers for.
func TestPrepareSVGMarker(t *testing.T) {
	src := `<svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 40 40">
<style type="text/css"><![CDATA[
.head { fill: #FF0000; }
]]></style>
<marker id="head" markerWidth="10" markerHeight="10" refX="8" refY="5" viewBox="0 0 10 10" orient="auto" markerUnits="userSpaceOnUse"><polygon class="head" points="0,0 10,5 0,10"/></marker>
<path d="M 0 0 L 40 40" fill="none" marker-end="url(#head)"/>
</svg>`
	prepared, err := Prepare([]byte(src), false)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	out := string(prepared)
	if strings.Contains(out, "<marker") || strings.Contains(out, "marker-end=") {
		t.Fatal("the marker definitions were kept")
	}
	g := regexp.MustCompile(`<g transform="translate\([^"]+\)"`).FindString(out)
	if g == "" {
		t.Fatal("no arrowhead was placed on the path")
	}
	// The path runs down and right, so the head turns 45 degrees.
	if !strings.Contains(g, "rotate(45.00") {
		t.Errorf("arrowhead %q is not turned along the path", g)
	}
	if !strings.Contains(out, `class="head"`) {
		t.Error("the arrowhead lost the class its fill comes from")
	}
}

// TestSVGFonts reads the fonts a diagram carries: they arrive as WOFF,
// and shaping a label with them is what draws it as paths.
func TestSVGFonts(t *testing.T) {
	root, _, err := parseSVGTree(d2SVG(t, `a -> b`))
	if err != nil {
		t.Fatalf("parseSVGTree: %v", err)
	}
	st := &svgStyles{faces: map[string]string{}}
	st.rewrite(root, true)
	if len(st.faces) == 0 {
		t.Fatal("no @font-face was read")
	}
	fonts := newSVGFonts(st.fonts())
	if len(fonts.faces) == 0 || fonts.def == nil {
		t.Fatalf("decoded %d fonts, none usable", len(fonts.faces))
	}
	width, glyphs := fonts.shape(fonts.def, 16, "Hello")
	if width <= 0 || len(glyphs) != 5 {
		t.Errorf("shaped %q to %v in %d glyphs, want a width and 5 glyphs", "Hello", width, len(glyphs))
	}
}

// assertSimpleSelectors checks that every rule a style sheet holds has the
// simple selectors MyGo matches: a selector with a combinator is dropped
// by MyGo, which loses the declaration.
func assertSimpleSelectors(t *testing.T, doc string) {
	t.Helper()
	sheets := regexp.MustCompile(`(?s)<style[^>]*>(.*?)</style>`).FindAllStringSubmatch(doc, -1)
	if len(sheets) == 0 {
		t.Fatal("the document holds no style sheet")
	}
	sel := regexp.MustCompile(`([^{@\n]+)\{`)
	for _, sheet := range sheets {
		for _, m := range sel.FindAllStringSubmatch(sheet[1], -1) {
			for _, s := range strings.Split(m[1], ",") {
				s = strings.TrimSpace(s)
				if s == "" || strings.ContainsAny(s, " \t\r\n") {
					t.Errorf("selector %q has a combinator MyGo would drop", s)
				}
			}
		}
	}
}

// countColor is how many pixels of an image are within a few steps of a
// color, which antialiasing leaves off.
func countColor(img image.Image, r, g, b uint8) int {
	bnd := img.Bounds()
	n := 0
	for y := bnd.Min.Y; y < bnd.Max.Y; y++ {
		for x := bnd.Min.X; x < bnd.Max.X; x++ {
			cr, cg, cb, _ := img.At(x, y).RGBA()
			if abs(int(cr>>8)-int(r)) <= 2 &&
				abs(int(cg>>8)-int(g)) <= 2 &&
				abs(int(cb>>8)-int(b)) <= 2 {
				n++
			}
		}
	}
	return n
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

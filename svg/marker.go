package svg

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// expandMarkers draws each arrowhead where the path it belongs to ends,
// from the graphics its marker holds, and drops the marker definitions:
// MyGo draws no markers.
func expandMarkers(root *svgNode) {
	markers := map[string]*svgNode{}
	walk(root, func(n *svgNode) {
		if n.name == "marker" {
			if id, ok := n.attr("id"); ok {
				markers[id] = n
			}
		}
	})
	placeMarkers(root, markers)
}

// placeMarkers walks the tree, putting the arrowheads of each element
// after it, and dropping the marker definitions it has collected.
func placeMarkers(n *svgNode, markers map[string]*svgNode) {
	if n.name == "#text" {
		return
	}
	var kids []*svgNode
	for _, k := range n.kids {
		if k.name == "marker" {
			continue
		}
		kids = append(kids, k)
		kids = append(kids, markerAt(k, markers)...)
	}
	n.kids = kids
	for _, k := range n.kids {
		placeMarkers(k, markers)
	}
}

// markerAt returns a copy of the marker each of an element's marker
// attributes names, placed where it points, and clears the attributes.
func markerAt(n *svgNode, markers map[string]*svgNode) []*svgNode {
	var out []*svgNode
	for _, name := range []string{"marker-end", "marker-start"} {
		v, ok := n.attr(name)
		if !ok {
			continue
		}
		n.delAttr(name)
		m := markers[urlRef(v)]
		if m == nil {
			continue
		}
		if g, ok := markerCopy(n, m, name == "marker-start"); ok {
			out = append(out, g)
		}
	}
	return out
}

// urlRef is what a url(#id) reference points at.
func urlRef(v string) string {
	v = strings.TrimSpace(v)
	if i := strings.Index(v, "url("); i >= 0 {
		v = v[i+4:]
		if j := strings.IndexByte(v, ')'); j >= 0 {
			v = v[:j]
		}
	}
	return strings.Trim(strings.TrimSpace(v), `"'#`)
}

// markerCopy builds the group of graphics a marker holds, moved to the
// point of n it belongs to and turned the way the path runs there.
func markerCopy(n, m *svgNode, atStart bool) (*svgNode, bool) {
	x, y, deg, ok := elementPoint(n, atStart)
	if !ok {
		return nil, false
	}
	orient, _ := m.attr("orient")
	switch orient {
	case "", "0":
		deg = 0
	case "auto":
	case "auto-start-reverse":
		if atStart {
			deg += 180
		}
	default:
		if v, err := strconv.ParseFloat(strings.TrimSuffix(orient, "deg"), 64); err == nil {
			deg = v
		}
	}

	mw, mh := numberOr(m, "markerWidth", 3), numberOr(m, "markerHeight", 3)
	refX, refY := numberOr(m, "refX", 0), numberOr(m, "refY", 0)
	vx, vy, vw, vh := 0.0, 0.0, mw, mh
	if vb, ok := m.attr("viewBox"); ok {
		if p := strings.Fields(strings.ReplaceAll(vb, ",", " ")); len(p) == 4 {
			if a, b, c, d := parseFour(p); a != 0 || b != 0 || c != 0 || d != 0 {
				vx, vy, vw, vh = a, b, c, d
			}
		}
	}
	scale := 1.0
	if vw > 0 && vh > 0 {
		scale = math.Min(mw/vw, mh/vh)
	}
	// The marker's viewport sits at the point, its reference point on it,
	// in the coordinates its viewBox scales into.
	offX, offY := (mw-vw*scale)/2, (mh-vh*scale)/2
	refCX := (refX-offX)/scale + vx
	refCY := (refY-offY)/scale + vy
	if units, _ := m.attr("markerUnits"); units != "" && units != "userSpaceOnUse" {
		scale *= strokeWidthOf(n)
	}

	g := &svgNode{name: "g"}
	g.setAttr("transform", fmt.Sprintf(
		"translate(%s %s) rotate(%s) scale(%s) translate(%s %s)",
		coord2(x), coord2(y), coord2(deg), coord2(scale), coord2(-refCX), coord2(-refCY),
	))
	for _, a := range m.attrs {
		switch a.name {
		case "id", "markerWidth", "markerHeight", "markerUnits", "marker-start",
			"marker-mid", "marker-end", "orient", "preserveAspectRatio",
			"refX", "refY", "viewBox", "overflow":
			continue
		}
		g.attrs = append(g.attrs, a)
	}
	g.kids = deepCopy(m.kids)
	return g, true
}

// strokeWidthOf is the width a marker of the element scales with.
func strokeWidthOf(n *svgNode) float64 {
	if v, ok := n.styleDecl("stroke-width"); ok {
		return parseLength(v, 1)
	}
	if v, ok := n.attr("stroke-width"); ok {
		return parseLength(v, 1)
	}
	return 1
}

// numberOr is an attribute as a number, or a default.
func numberOr(n *svgNode, name string, def float64) float64 {
	v, ok := n.attr(name)
	if !ok {
		return def
	}
	return parseLength(v, def)
}

// parseFour reads four numbers, as a viewBox holds.
func parseFour(p []string) (float64, float64, float64, float64) {
	var v [4]float64
	for i := range p {
		if i > 3 {
			break
		}
		v[i], _ = strconv.ParseFloat(strings.TrimSpace(p[i]), 64)
	}
	return v[0], v[1], v[2], v[3]
}

// elementPoint is where a marker belongs on an element: its first or last
// point, and the direction the path runs there in degrees.
func elementPoint(n *svgNode, atStart bool) (x, y, deg float64, ok bool) {
	if d, has := n.attr("d"); has {
		g := pathGeometry(d)
		if !g.ok {
			return 0, 0, 0, false
		}
		if atStart {
			return g.startX, g.startY, g.startDeg, true
		}
		return g.endX, g.endY, g.endDeg, true
	}
	if p, has := n.attr("points"); has {
		pts := pathPoints(p)
		if len(pts) < 2 {
			return 0, 0, 0, false
		}
		if atStart {
			return pts[0].x, pts[0].y, angle(pts[0], pts[1]), true
		}
		last, prev := pts[len(pts)-1], pts[len(pts)-2]
		return last.x, last.y, angle(prev, last), true
	}
	if n.name == "line" {
		x1, y1 := numberOr(n, "x1", 0), numberOr(n, "y1", 0)
		x2, y2 := numberOr(n, "x2", 0), numberOr(n, "y2", 0)
		if atStart {
			return x1, y1, angle(point{x1, y1}, point{x2, y2}), true
		}
		return x2, y2, angle(point{x1, y1}, point{x2, y2}), true
	}
	return 0, 0, 0, false
}

// deepCopy copies a subtree of graphics.
func deepCopy(kids []*svgNode) []*svgNode {
	out := make([]*svgNode, 0, len(kids))
	for _, k := range kids {
		c := &svgNode{name: k.name, attrs: append([]svgAttr(nil), k.attrs...), chars: k.chars}
		c.kids = deepCopy(k.kids)
		out = append(out, c)
	}
	return out
}

// coord2 formats a number for an attribute, not for path data.
func coord2(v float64) string {
	return strconv.FormatFloat(v, 'f', 2, 64)
}

type point struct{ x, y float64 }

func angle(from, to point) float64 {
	return math.Atan2(to.y-from.y, to.x-from.x) * 180 / math.Pi
}

// pathPoints reads a points attribute, of a polyline or a polygon.
func pathPoints(s string) []point {
	fields := strings.Fields(strings.ReplaceAll(s, ",", " "))
	var pts []point
	for i := 0; i+1 < len(fields); i += 2 {
		x, err1 := strconv.ParseFloat(fields[i], 64)
		y, err2 := strconv.ParseFloat(fields[i+1], 64)
		if err1 != nil || err2 != nil {
			break
		}
		pts = append(pts, point{x, y})
	}
	return pts
}

// pathGeom is where a path begins and ends, and which way it runs there.
type pathGeom struct {
	startX, startY float64
	endX, endY     float64
	startDeg       float64
	endDeg         float64
	ok             bool
}

// pathGeometry reads the ends of a path's first subpath: enough to put an
// arrowhead where the path stops, turned along it.
func pathGeometry(d string) pathGeom {
	var g pathGeom
	r := &pathReader{s: d}
	var x, y float64       // the point the path has reached
	var sx, sy float64     // where its first subpath began
	var started, seg bool  // saw a subpath, saw a segment in it
	var outX, outY float64 // the direction the first segment leaves with
	var inX, inY float64   // the direction the last segment arrives with
	var prev byte
	for {
		op, ok := r.command(prev)
		if !ok {
			break
		}
		prev = op
		rel := op >= 'a'
		o := op
		if rel {
			o = op - 'a' + 'A'
		}
		switch o {
		case 'M':
			nx, ny, ok := r.pair()
			if !ok {
				return g
			}
			if rel {
				nx, ny = x+nx, y+ny
			}
			if !started {
				sx, sy = nx, ny
				g.startX, g.startY = nx, ny
				started = true
			}
			x, y = nx, ny
		case 'Z':
			if started {
				inX, inY = sx-x, sy-y
				x, y = sx, sy
				if !seg {
					g.startX, g.startY = x, y
				}
				seg = true
			}
		default:
			args, ok := r.args(o)
			if !ok {
				return g
			}
			nx, ny := x, y
			var c1x, c1y, c2x, c2y float64
			switch o {
			case 'H':
				nx = args[0]
				if rel {
					nx = x + args[0]
				}
			case 'V':
				ny = args[0]
				if rel {
					ny = y + args[0]
				}
			case 'L', 'T', 'A':
				nx, ny = args[0], args[1]
				if rel {
					nx, ny = x+nx, y+ny
				}
			case 'Q', 'S':
				c1x, c1y = args[0], args[1]
				nx, ny = args[2], args[3]
				if rel {
					c1x, c1y = x+c1x, y+c1y
					nx, ny = x+nx, y+ny
				}
			case 'C':
				c1x, c1y = args[0], args[1]
				c2x, c2y = args[2], args[3]
				nx, ny = args[4], args[5]
				if rel {
					c1x, c1y = x+c1x, y+c1y
					c2x, c2y = x+c2x, y+c2y
					nx, ny = x+nx, y+ny
				}
			}
			if !seg {
				switch o {
				case 'Q', 'S', 'C':
					outX, outY = c1x-x, c1y-y
				default:
					outX, outY = nx-x, ny-y
				}
			}
			switch o {
			case 'Q', 'S':
				inX, inY = nx-c1x, ny-c1y
			case 'C':
				inX, inY = nx-c2x, ny-c2y
			default:
				inX, inY = nx-x, ny-y
			}
			x, y = nx, ny
			seg = true
		}
	}
	if !started {
		return g
	}
	if seg {
		g.endX, g.endY = x, y
		if outX != 0 || outY != 0 {
			g.startDeg = math.Atan2(outY, outX) * 180 / math.Pi
		}
		if inX != 0 || inY != 0 {
			g.endDeg = math.Atan2(inY, inX) * 180 / math.Pi
		}
	} else {
		g.endX, g.endY = g.startX, g.startY
	}
	g.ok = true
	return g
}

// pathReader reads the commands of path data.
type pathReader struct {
	s string
	i int
}

func (r *pathReader) skipSpace() {
	for r.i < len(r.s) && isSVGSpace(r.s[r.i]) {
		r.i++
	}
}

// command returns the next command, repeating the previous one where the
// path writes only numbers, as SVG allows.
func (r *pathReader) command(prev byte) (byte, bool) {
	for {
		r.skipSpace()
		if r.i >= len(r.s) {
			return 0, false
		}
		c := r.s[r.i]
		if isPathOp(c) {
			r.i++
			return c, true
		}
		if prev != 0 && r.startsNumber() {
			if prev == 'M' {
				return 'L', true
			}
			if prev == 'm' {
				return 'l', true
			}
			return prev, true
		}
		r.i++
	}
}

func (r *pathReader) startsNumber() bool {
	if r.i >= len(r.s) {
		return false
	}
	c := r.s[r.i]
	return c == '+' || c == '-' || c == '.' || c >= '0' && c <= '9'
}

// args reads the numbers a command takes.
func (r *pathReader) args(o byte) ([]float64, bool) {
	var want int
	switch o {
	case 'H', 'V':
		want = 1
	case 'L', 'T', 'M', 'A':
		want = 2
	case 'Q', 'S':
		want = 4
	case 'C':
		want = 6
	default:
		return nil, true
	}
	args := make([]float64, want)
	for i := range args {
		v, ok := r.number()
		if !ok {
			return nil, false
		}
		args[i] = v
	}
	return args, true
}

func (r *pathReader) pair() (float64, float64, bool) {
	x, ok := r.number()
	if !ok {
		return 0, 0, false
	}
	y, ok := r.number()
	if !ok {
		return 0, 0, false
	}
	return x, y, true
}

func (r *pathReader) number() (float64, bool) {
	r.skipSpace()
	start := r.i
	if r.i < len(r.s) && (r.s[r.i] == '+' || r.s[r.i] == '-') {
		r.i++
	}
	seen := false
	for r.i < len(r.s) && r.s[r.i] >= '0' && r.s[r.i] <= '9' {
		r.i++
		seen = true
	}
	if r.i < len(r.s) && r.s[r.i] == '.' {
		r.i++
		for r.i < len(r.s) && r.s[r.i] >= '0' && r.s[r.i] <= '9' {
			r.i++
			seen = true
		}
	}
	if r.i < len(r.s) && (r.s[r.i] == 'e' || r.s[r.i] == 'E') {
		j := r.i + 1
		if j < len(r.s) && (r.s[j] == '+' || r.s[j] == '-') {
			j++
		}
		if j < len(r.s) && r.s[j] >= '0' && r.s[j] <= '9' {
			r.i = j
			for r.i < len(r.s) && r.s[r.i] >= '0' && r.s[r.i] <= '9' {
				r.i++
			}
			seen = true
		}
	}
	if !seen {
		r.i = start
		return 0, false
	}
	v, err := strconv.ParseFloat(r.s[start:r.i], 64)
	if err != nil {
		r.i = start
		return 0, false
	}
	return v, true
}

func isPathOp(c byte) bool {
	switch c {
	case 'M', 'm', 'L', 'l', 'H', 'h', 'V', 'v', 'C', 'c',
		'S', 's', 'Q', 'q', 'T', 't', 'A', 'a', 'Z', 'z':
		return true
	}
	return false
}

func isSVGSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n' || c == '\f'
}

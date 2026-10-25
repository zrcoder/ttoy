package svg

import (
	"strconv"
	"strings"
)

// textEnv is the text an element inherits: the size its labels are drawn
// at, how they sit along a line, and the family its classes name.
type textEnv struct {
	size   float64
	anchor string
	family string
}

// convertText replaces every label with paths of the glyphs it is made of,
// drawn with the fonts the document embeds: MyGo draws paths, but no text.
func convertText(n *svgNode, st *svgStyles, fonts *svgFonts, env textEnv) {
	if n.name == "#text" || n.name == "style" {
		return
	}
	own := env
	if v, ok := n.styleDecl("font-size"); ok {
		own.size = parseLength(v, env.size)
	}
	if v, ok := n.styleDecl("text-anchor"); ok {
		own.anchor = v
	}
	if v, ok := n.styleDecl("font-family"); ok {
		own.family = strings.Trim(v, `"'`)
	}
	if family := st.familyOf(n.classes()); family != "" {
		own.family = family
	}
	if n.name == "text" {
		labelToPaths(n, st, fonts, own)
		return
	}
	for _, k := range n.kids {
		convertText(k, st, fonts, own)
	}
}

// labelRun is one stretch of a label's characters: either the element's
// own, or a tspan's, which may carry a color of its own.
type labelRun struct {
	text    string
	tspan   *svgNode
	newLine bool
	width   float64
	glyphs  []shapedGlyph
}

// labelLines splits a label into the lines it holds, each line with the
// runs to place along it.
type labelLine struct {
	x, y  float64
	runs  []labelRun
	width float64
}

// labelToPaths turns one label into a group of paths, one per run, so
// that a tspan's own color survives.
func labelToPaths(n *svgNode, st *svgStyles, fonts *svgFonts, env textEnv) {
	runs := labelRuns(n)
	if len(runs) == 0 {
		n.name = "g"
		n.delAttr("x")
		n.delAttr("y")
		n.kids = nil
		return
	}

	x, y := lengthOf(n, "x", env.size), lengthOf(n, "y", env.size)
	var lines []*labelLine
	for i := range runs {
		r := &runs[i]
		size, family := env.size, env.family
		if r.tspan != nil {
			if v, ok := r.tspan.styleDecl("font-size"); ok {
				size = parseLength(v, env.size)
			}
			if v, ok := r.tspan.styleDecl("font-family"); ok {
				family = strings.Trim(v, `"'`)
			}
			if f := st.familyOf(r.tspan.classes()); f != "" {
				family = f
			}
		}
		r.width, r.glyphs = fonts.shape(fonts.face(family), size, r.text)

		if len(lines) == 0 || r.newLine {
			lx, ly := x, y
			if len(lines) > 0 {
				lx, ly = lines[len(lines)-1].x, lines[len(lines)-1].y
				if r.tspan != nil {
					if v, ok := r.tspan.attr("dy"); ok {
						ly += parseLength(v, env.size)
					}
					if v, ok := r.tspan.attr("y"); ok {
						ly = parseLength(v, env.size)
					}
					if v, ok := r.tspan.attr("x"); ok {
						lx = parseLength(v, env.size)
					}
				}
			}
			lines = append(lines, &labelLine{x: lx, y: ly})
		}
		line := lines[len(lines)-1]
		line.runs = append(line.runs, *r)
		line.width += r.width
	}

	var factor float64
	switch env.anchor {
	case "middle":
		factor = 0.5
	case "end":
		factor = 1
	}
	paths := make([]*svgNode, 0, len(runs))
	for _, line := range lines {
		pen := line.x - line.width*factor
		for _, r := range line.runs {
			if len(r.glyphs) == 0 {
				pen += r.width
				continue
			}
			var d strings.Builder
			appendPath(&d, r.glyphs, pen, line.y)
			pen += r.width
			if d.Len() == 0 {
				continue
			}
			p := &svgNode{name: "path", attrs: labelPathAttrs(n, r.tspan)}
			p.setAttr("d", d.String())
			paths = append(paths, p)
		}
	}

	n.name = "g"
	n.delAttr("x")
	n.delAttr("y")
	n.kids = paths
}

// labelRuns reads a label's characters in order: the element's own as
// runs of their own, and each tspan as one, a tspan starting a new line
// when it says where it is.
func labelRuns(n *svgNode) []labelRun {
	var runs []labelRun
	for _, k := range n.kids {
		switch k.name {
		case "#text":
			s := k.chars
			if s == "" || (strings.TrimSpace(s) == "" && strings.ContainsAny(s, "\r\n")) {
				continue // the whitespace a pretty printer put between tags
			}
			runs = append(runs, labelRun{text: s})
		case "tspan":
			s := k.ownText()
			if s == "" {
				continue
			}
			newLine := false
			for _, at := range []string{"x", "y", "dy"} {
				if _, ok := k.attr(at); ok {
					newLine = true
				}
			}
			runs = append(runs, labelRun{text: s, tspan: k, newLine: newLine})
		}
	}
	if len(runs) > 0 {
		runs[0].newLine = true
	}
	return runs
}

// labelPathAttrs is what a path of a run carries: the label's attributes,
// its class and style merged with the tspan's, and no geometry.
func labelPathAttrs(label, tspan *svgNode) []svgAttr {
	out := append([]svgAttr(nil), label.attrs...)
	out = removeAttr(removeAttr(out, "x"), "y")
	class, style := attrValue(out, "class"), attrValue(out, "style")
	out = removeAttr(removeAttr(out, "class"), "style")
	if tspan != nil {
		for _, a := range tspan.attrs {
			switch a.name {
			case "x", "y", "dx", "dy":
				continue
			case "class":
				class = strings.TrimSpace(class + " " + a.value)
				continue
			case "style":
				style = joinStyles(style, a.value)
				continue
			}
			out = setAttrIn(out, a.name, a.value)
		}
	}
	if class != "" {
		out = setAttrIn(out, "class", class)
	}
	if style != "" {
		out = setAttrIn(out, "style", style)
	}
	return out
}

// attrValue is an attribute's value, or none.
func attrValue(attrs []svgAttr, name string) string {
	for _, a := range attrs {
		if a.name == name {
			return a.value
		}
	}
	return ""
}

// setAttrIn gives a name a value in a list of attributes.
func setAttrIn(attrs []svgAttr, name, value string) []svgAttr {
	for i := range attrs {
		if attrs[i].name == name {
			attrs[i].value = value
			return attrs
		}
	}
	return append(attrs, svgAttr{name, value})
}

// joinStyles puts two style attributes together, the second winning.
func joinStyles(a, b string) string {
	a, b = strings.TrimSuffix(strings.TrimSpace(a), ";"), strings.TrimSpace(b)
	switch {
	case a == "":
		return b
	case b == "":
		return a
	}
	return a + ";" + b
}

// lengthOf is an x or y attribute as a number, in the units it is written
// with: em counts the font size, as the y of a code line does.
func lengthOf(n *svgNode, name string, base float64) float64 {
	v, ok := n.attr(name)
	if !ok {
		return 0
	}
	return parseLength(v, base)
}

// parseLength reads a length in the units SVG allows, against base, the
// font size an em counts.
func parseLength(s string, base float64) float64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return base
	}
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return v
	}
	for _, u := range []struct {
		suffix string
		to     float64
	}{
		{"px", 1},
		{"rem", base},
		{"em", base},
		{"pt", 96.0 / 72},
		{"cm", 96.0 / 2.54},
		{"mm", 96.0 / 25.4},
		{"in", 96},
	} {
		if strings.HasSuffix(s, u.suffix) {
			if v, err := strconv.ParseFloat(strings.TrimSuffix(s, u.suffix), 64); err == nil {
				return v * u.to
			}
		}
	}
	if strings.HasSuffix(s, "%") {
		if v, err := strconv.ParseFloat(strings.TrimSuffix(s, "%"), 64); err == nil {
			return v / 100 * base
		}
	}
	return base
}

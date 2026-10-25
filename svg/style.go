package svg

import (
	"strings"
)

// cssDecl is one property of a rule or of a style attribute.
type cssDecl struct{ name, value string }

// parseDecls reads the declarations of a block or of a style attribute,
// splitting them where a semicolon falls outside a value: a font source
// carries one in its base64.
func parseDecls(s string) []cssDecl {
	s = stripCSSComments(s)
	var out []cssDecl
	depth := 0
	var quote byte
	start := 0
	flush := func(i int) {
		if d, ok := declOf(s[start:i]); ok {
			out = append(out, d)
		}
		start = i + 1
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case quote != 0:
			if c == '\\' {
				i++
			} else if c == quote {
				quote = 0
			}
		case c == '"' || c == '\'':
			quote = c
		case c == '(':
			depth++
		case c == ')':
			if depth > 0 {
				depth--
			}
		case c == ';' && depth == 0:
			flush(i)
		}
	}
	flush(len(s))
	return out
}

// declOf reads one name and value pair.
func declOf(s string) (cssDecl, bool) {
	i := strings.IndexByte(s, ':')
	if i < 0 {
		return cssDecl{}, false
	}
	name := strings.ToLower(strings.TrimSpace(s[:i]))
	if name == "" {
		return cssDecl{}, false
	}
	return cssDecl{name, strings.TrimSpace(s[i+1:])}, true
}

// declValue returns the last declaration of a name, which wins.
func declValue(decls []cssDecl, name string) (string, bool) {
	var v string
	var ok bool
	for _, d := range decls {
		if d.name == name {
			v, ok = d.value, true
		}
	}
	return v, ok
}

// cssRule is a rule whose selectors MyGo matches: a type, an id and
// classes, without combinators.
type cssRule struct {
	classes []string
	decls   []cssDecl
}

// svgStyles is what a document's style sheets say: which classes name a
// font, and which fonts the document carries.
type svgStyles struct {
	rules []cssRule
	faces map[string]string // font-family name -> data URL
}

// rewrite reads every style sheet of the document, records the rules and
// @font-face fonts it holds, and puts back what MyGo matches: MyGo drops
// a selector with a combinator, so the ancestors before the class a rule
// really wants are dropped here, leaving the class itself. The one query
// a diagram carries, for the system's color scheme, is unwrapped into the
// branch of the current theme, as MyGo runs no queries.
func (st *svgStyles) rewrite(root *svgNode, dark bool) {
	walk(root, func(n *svgNode) {
		if n.name != "style" {
			return
		}
		css, rules, faces := rewriteCSS(n.ownText(), dark)
		st.rules = append(st.rules, rules...)
		for family, src := range faces {
			st.faces[family] = src
		}
		kids := n.kids[:0]
		for _, k := range n.kids {
			if k.name != "#text" {
				kids = append(kids, k)
			}
		}
		n.kids = append([]*svgNode{{name: "#text", chars: css}}, kids...)
	})
}

// familyOf returns the font-family the style sheet gives a class of the
// element, or none when the rules name none.
func (st *svgStyles) familyOf(classes []string) string {
	if len(classes) == 0 {
		return ""
	}
	var family string
	for _, r := range st.rules {
		if !intersects(r.classes, classes) {
			continue
		}
		if v, ok := declValue(r.decls, "font-family"); ok {
			family = v
		}
	}
	return strings.Trim(family, `"'`)
}

// fonts decodes the WOFF the document's @font-face rules carry into the
// sfnt bytes a font package reads.
func (st *svgStyles) fonts() map[string][]byte {
	out := make(map[string][]byte, len(st.faces))
	for family, src := range st.faces {
		data, err := dataURIBytes(src)
		if err != nil {
			continue
		}
		ttf, err := woffToTTF(data)
		if err != nil {
			continue
		}
		out[family] = ttf
	}
	return out
}

// intersects tells whether two sets of names share one.
func intersects(a, b []string) bool {
	for _, x := range a {
		for _, y := range b {
			if x == y {
				return true
			}
		}
	}
	return false
}

// rewriteCSS flattens the selectors of a style sheet, and returns the
// sheet to write back, the rules to match elements against, and the fonts
// its @font-face rules name.
func rewriteCSS(src string, dark bool) (string, []cssRule, map[string]string) {
	o := &cssOut{faces: map[string]string{}}
	rewriteCSSInto(o, stripCSSComments(src), dark)
	return o.b.String(), o.rules, o.faces
}

// cssOut is what a style sheet comes out as.
type cssOut struct {
	b     strings.Builder
	rules []cssRule
	faces map[string]string
}

// rewriteCSSInto reads one style sheet into o: at-rules pass through as
// they are, save for the color scheme query, whose branch the theme is in
// is written out in its place.
func rewriteCSSInto(o *cssOut, s string, dark bool) {
	for {
		s = strings.TrimSpace(s)
		if s == "" {
			return
		}
		switch {
		case strings.HasPrefix(s, "<!--"):
			s = s[4:]
			continue
		case strings.HasPrefix(s, "-->"):
			s = s[3:]
			continue
		case strings.HasPrefix(s, "@"):
			semi, brace := strings.IndexByte(s, ';'), strings.IndexByte(s, '{')
			if semi >= 0 && (brace < 0 || semi < brace) {
				o.b.WriteString(s[:semi+1])
				o.b.WriteByte('\n')
				s = s[semi+1:]
				continue
			}
			if brace < 0 {
				o.b.WriteString(s)
				return
			}
			end := cssBlockEnd(s, brace)
			raw, pre := strings.TrimSpace(s[:end]), strings.TrimSpace(s[:brace])
			body := s[brace+1 : end-1]
			s = s[end:]
			if isSchemeQuery(pre) {
				if schemeApplies(pre, dark) {
					rewriteCSSInto(o, stripCSSComments(body), dark)
				}
				continue
			}
			o.b.WriteString(raw)
			o.b.WriteByte('\n')
			if pre == "@font-face" {
				decls := parseDecls(body)
				family, ok := declValue(decls, "font-family")
				src, srcOK := declValue(decls, "src")
				if ok && srcOK {
					o.faces[strings.Trim(family, `"'`)] = src
				}
			}
			continue
		}
		open := strings.IndexByte(s, '{')
		if open < 0 {
			return
		}
		end := cssBlockEnd(s, open)
		decls := parseDecls(s[open+1 : end-1])
		var sel []string
		for _, part := range strings.Split(stripCSSComments(s[:open]), ",") {
			if f, ok := flattenSelector(strings.TrimSpace(part)); ok {
				sel = append(sel, f)
			}
		}
		s = s[end:]
		if len(sel) == 0 {
			continue
		}
		o.rules = append(o.rules, cssRule{classes: classesOf(sel), decls: decls})
		o.b.WriteString(strings.Join(sel, ", "))
		o.b.WriteString(" {\n")
		for _, d := range decls {
			o.b.WriteString("\t")
			o.b.WriteString(d.name)
			o.b.WriteString(": ")
			o.b.WriteString(d.value)
			o.b.WriteString(";\n")
		}
		o.b.WriteString("}\n")
	}
}

// isSchemeQuery tells whether an at-rule selects on the system's color
// scheme: the one query a diagram carries, and one MyGo runs no query for.
func isSchemeQuery(pre string) bool {
	return strings.HasPrefix(pre, "@media") && strings.Contains(pre, "prefers-color-scheme")
}

// schemeApplies tells whether a color scheme query holds in the theme.
func schemeApplies(pre string, dark bool) bool {
	q := strings.ReplaceAll(strings.ToLower(pre), " ", "")
	switch {
	case strings.Contains(q, "prefers-color-scheme:dark"):
		return dark
	case strings.Contains(q, "prefers-color-scheme:light"):
		return !dark
	}
	return true
}

// flattenSelector keeps the last simple compound of a selector: the part
// after the combinators, which is the one MyGo matches on its own. It
// reports false for a selector MyGo would not match at all, as one with a
// pseudo-class.
func flattenSelector(sel string) (string, bool) {
	fields := strings.Fields(sel)
	if len(fields) == 0 {
		return "", false
	}
	last := fields[len(fields)-1]
	if last == "" || strings.HasPrefix(last, "@") {
		return "", false
	}
	for i := 0; i < len(last); i++ {
		switch c := last[i]; c {
		case ':', '[', ']', '(', ')':
			return "", false
		}
	}
	return last, true
}

// classesOf is the class names of each selector of a rule.
func classesOf(sel []string) []string {
	var out []string
	for _, s := range sel {
		for {
			i := strings.IndexByte(s, '.')
			if i < 0 {
				break
			}
			s = s[i+1:]
			j := 0
			for j < len(s) && isNameChar(s[j]) {
				j++
			}
			if j == 0 {
				break
			}
			out = append(out, s[:j])
			s = s[j:]
		}
	}
	return out
}

// isNameChar tells whether c may be part of a CSS name.
func isNameChar(c byte) bool {
	return c == '-' || c == '_' ||
		c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' ||
		c >= 0x80
}

// stripCSSComments removes /* ... */ sections.
func stripCSSComments(s string) string {
	for {
		i := strings.Index(s, "/*")
		if i < 0 {
			return s
		}
		j := strings.Index(s[i+2:], "*/")
		if j < 0 {
			return s[:i]
		}
		s = s[:i] + " " + s[i+2+j+2:]
	}
}

// cssBlockEnd returns the index after the block opened at s[open].
func cssBlockEnd(s string, open int) int {
	depth := 0
	for i := open; i < len(s); i++ {
		switch s[i] {
		case '{':
			depth++
		case '}':
			if depth--; depth == 0 {
				return i + 1
			}
		}
	}
	return len(s)
}

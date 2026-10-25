// Package svg prepares a diagram SVG for MyGo's ui.SVG, which draws shapes
// but no text, no markers, and only style sheets of simple selectors. The
// one entry point is Prepare.
package svg

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// Prepare rewrites a diagram SVG into the subset of SVG that MyGo's
// ui.SVG draws. MyGo draws shapes but leaves out text and markers, and
// matches style sheets of simple selectors only, so this turns labels into
// paths drawn with the fonts the document embeds, arrows into graphics
// where markers were, and descendant selectors into simple ones. What
// comes out renders as vectors, at any size, with no rasterizing step.
// The document's color scheme query is unwrapped into the branch dark
// tells, as MyGo runs no queries.
func Prepare(data []byte, dark bool) ([]byte, error) {
	root, ns, err := parseSVGTree(data)
	if err != nil {
		return nil, err
	}
	st := &svgStyles{faces: map[string]string{}}
	st.rewrite(root, dark)
	convertText(root, st, newSVGFonts(st.fonts()), textEnv{size: 16, anchor: "start"})
	expandMarkers(root)
	var b bytes.Buffer
	b.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	writeSVG(&b, root, ns)
	return b.Bytes(), nil
}

// svgNode is an element of the document. Character data is a child named
// "#text", so that what a label holds keeps its place among its tspans.
type svgNode struct {
	name  string
	attrs []svgAttr
	kids  []*svgNode
	chars string // of a "#text" child
}

type svgAttr struct{ name, value string }

func (n *svgNode) attr(name string) (string, bool) {
	for _, a := range n.attrs {
		if a.name == name {
			return a.value, true
		}
	}
	return "", false
}

func (n *svgNode) setAttr(name, value string) {
	for i := range n.attrs {
		if n.attrs[i].name == name {
			n.attrs[i].value = value
			return
		}
	}
	n.attrs = append(n.attrs, svgAttr{name, value})
}

func (n *svgNode) delAttr(name string) {
	n.attrs = removeAttr(n.attrs, name)
}

func removeAttr(attrs []svgAttr, name string) []svgAttr {
	out := attrs[:0]
	for _, a := range attrs {
		if a.name != name {
			out = append(out, a)
		}
	}
	return out
}

// classes is the class attribute split up, as the style sheet names them.
func (n *svgNode) classes() []string {
	v, ok := n.attr("class")
	if !ok {
		return nil
	}
	return strings.Fields(v)
}

// ownText is the character data the element holds directly, between its
// children.
func (n *svgNode) ownText() string {
	var b strings.Builder
	for _, k := range n.kids {
		if k.name == "#text" {
			b.WriteString(k.chars)
		}
	}
	return b.String()
}

// styleDecl returns a declaration of the element's style attribute.
func (n *svgNode) styleDecl(name string) (string, bool) {
	v, ok := n.attr("style")
	if !ok {
		return "", false
	}
	return declValue(parseDecls(v), name)
}

// parseSVGTree reads an SVG document into a tree, and returns the
// prefixes the document declared for its namespaces.
func parseSVGTree(data []byte) (root *svgNode, ns map[string]string, err error) {
	dec := xml.NewDecoder(bytes.NewReader(data))
	ns = map[string]string{}
	var stack []*svgNode
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("svg: %w", err)
		}
		switch t := tok.(type) {
		case xml.StartElement:
			n := &svgNode{name: t.Name.Local, attrs: make([]svgAttr, 0, len(t.Attr))}
			for _, a := range t.Attr {
				if a.Name.Space == "xmlns" {
					ns[a.Value] = a.Name.Local
				}
			}
			for _, a := range t.Attr {
				n.attrs = append(n.attrs, svgAttr{name: attrRaw(a, ns), value: a.Value})
			}
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				p.kids = append(p.kids, n)
			} else if root == nil {
				root = n
			}
			stack = append(stack, n)
		case xml.EndElement:
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			if len(stack) == 0 {
				continue
			}
			top := stack[len(stack)-1]
			if n := len(top.kids); n > 0 && top.kids[n-1].name == "#text" {
				top.kids[n-1].chars += string(t)
				continue
			}
			top.kids = append(top.kids, &svgNode{name: "#text", chars: string(t)})
		}
	}
	if root == nil || root.name != "svg" {
		return nil, nil, fmt.Errorf("svg: no svg element")
	}
	return root, ns, nil
}

// attrRaw writes an attribute name as the document did: with the prefix
// of the namespace it belongs to when the document declared one.
func attrRaw(a xml.Attr, ns map[string]string) string {
	switch {
	case a.Name.Space == "", a.Name.Local == "xmlns":
		return a.Name.Local
	case a.Name.Space == "xmlns":
		return "xmlns:" + a.Name.Local
	}
	if p, ok := ns[a.Name.Space]; ok {
		return p + ":" + a.Name.Local
	}
	return a.Name.Local
}

// walk visits every element of the tree, character data aside.
func walk(n *svgNode, fn func(*svgNode)) {
	if n.name == "#text" {
		return
	}
	fn(n)
	for _, k := range n.kids {
		walk(k, fn)
	}
}

// writeSVG serializes the tree: style sheets keep their CDATA section,
// which is the one text MyGo reads as it is written.
func writeSVG(b *bytes.Buffer, n *svgNode, ns map[string]string) {
	if n.name == "#text" {
		xmlEscape(b, n.chars)
		return
	}
	b.WriteByte('<')
	b.WriteString(n.name)
	for _, a := range n.attrs {
		b.WriteByte(' ')
		b.WriteString(a.name)
		b.WriteString(`="`)
		xmlEscape(b, a.value)
		b.WriteByte('"')
	}
	if len(n.kids) == 0 {
		b.WriteString("/>")
		return
	}
	b.WriteByte('>')
	if n.name == "style" {
		b.WriteString("<![CDATA[")
		for _, k := range n.kids {
			if k.name == "#text" {
				b.WriteString(k.chars)
			}
		}
		b.WriteString("]]>")
	} else {
		for _, k := range n.kids {
			writeSVG(b, k, ns)
		}
	}
	b.WriteString("</")
	b.WriteString(n.name)
	b.WriteByte('>')
}

// xmlEscape writes s with the characters XML escapes.
func xmlEscape(b *bytes.Buffer, s string) {
	_ = xml.EscapeText(b, []byte(s))
}

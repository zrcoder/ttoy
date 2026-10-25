package svg

import (
	"testing"

	"golang.org/x/image/font/sfnt"
)

// TestSystemFallback checks that this machine has a font to reach for when
// a document brings none of its own, which is how a label of a plain SVG
// is drawn. It skips where the list of candidates finds none, as a bare
// system may well have none of them.
func TestSystemFallback(t *testing.T) {
	fs := &svgFonts{}
	if len(fs.fallbacks()) == 0 {
		t.Skip("no font of the fallback list is installed here")
	}
	var buf sfnt.Buffer
	if _, idx := fs.glyphFor(nil, &buf, 'A'); idx == 0 {
		t.Error("no fallback font draws A")
	}
}

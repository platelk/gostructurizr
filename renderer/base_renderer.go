package renderer

import (
	"fmt"
	"io"
	"strings"

	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
)

// BaseRenderer provides common functionality for all renderers
type BaseRenderer struct {
	w     io.Writer
	level int
}

// WriteLine writes a line with proper indentation
func (r *BaseRenderer) WriteLine(format string, args ...interface{}) {
	indent := strings.Repeat("    ", r.level)
	fmt.Fprintf(r.w, "%s%s\n", indent, fmt.Sprintf(format, args...))
}

// Render is a placeholder for implementation by derived renderers
func (r *BaseRenderer) Render(interface{}) error {
	return fmt.Errorf("render not implemented")
}

// RenderProperties renders properties of an element
func renderProperties(w io.Writer, properties *gostructurizr.Properties, level int) {
	if properties == nil || len(properties.Properties) == 0 {
		return
	}

	indent := strings.Repeat("    ", level)
	fmt.Fprintf(w, "%s%s %s\n", indent, dsl.Properties, dsl.OpenBracket)
	
	for key, value := range properties.Properties {
		indentInner := strings.Repeat("    ", level+1)
		fmt.Fprintf(w, "%s%s %q\n", indentInner, key, value)
	}
	
	fmt.Fprintf(w, "%s%s\n", indent, dsl.CloseBracket)
}

// RenderTags renders tags of an element
func renderTags(w io.Writer, tags *gostructurizr.TagsNode, level int) {
	if tags == nil || len(tags.Tags) == 0 {
		return
	}

	indent := strings.Repeat("    ", level)
	tagList := strings.Join(tags.Tags, ", ")
	fmt.Fprintf(w, "%s%s %q\n", indent, dsl.Tags, tagList)
}

// RenderNamer renders a name for a Namer interface
func renderNamer(namer gostructurizr.Namer) string {
	return fmt.Sprintf("%q", namer.Name())
}
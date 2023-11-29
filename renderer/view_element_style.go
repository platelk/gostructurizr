package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderViewElementStyle(e *gostructurizr.ElementStyleNode, renderer *strings.Builder, level int) error {
	writeLine(renderer, level, dsl.Element, dsl.Space, generateStringIdentifier(e.Tag().String()), dsl.Space, dsl.OpenBracket)
	if e.Shape() != nil {
		writeLine(renderer, level+1, dsl.Shape, dsl.Space, e.Shape().String())
	}
	if e.Icon() != nil {
		writeLine(renderer, level+1, dsl.Icon, dsl.Space, generateStringIdentifier(*e.Icon()))
	}
	if e.Width() != nil {
		writeLine(renderer, level+1, dsl.Width, dsl.Space, fmt.Sprintf("%d", e.Width()))
	}
	if e.Height() != nil {
		writeLine(renderer, level+1, dsl.Height, dsl.Space, fmt.Sprintf("%d", e.Height()))
	}
	if e.Background() != nil {
		writeLine(renderer, level+1, dsl.Background, dsl.Space, *e.Background())
	}
	if e.Color() != nil {
		writeLine(renderer, level+1, dsl.Color, dsl.Space, *e.Color())
	}
	if e.Stroke() != nil {
		writeLine(renderer, level+1, dsl.Stroke, dsl.Space, *e.Stroke())
	}
	if e.Opacity() != nil {
		writeLine(renderer, level+1, dsl.Opacity, dsl.Space, fmt.Sprintf("%d", e.Opacity()))
	}
	if e.FontSize() != nil {
		writeLine(renderer, level+1, dsl.FontSize, dsl.Space, fmt.Sprintf("%d", e.FontSize()))
	}
	if e.Metadata() != nil {
		writeLine(renderer, level+1, dsl.Metadata, dsl.Space, fmt.Sprintf("%s", *e.Metadata()))
	}
	if e.Description() != nil {
		writeLine(renderer, level+1, dsl.Description, dsl.Space, fmt.Sprintf("%s", *e.Description()))
	}
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

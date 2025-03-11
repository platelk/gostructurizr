package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderViewElementStyle(e *gostructurizr.ElementStyleNode, renderer *strings.Builder, level int) error {
	writeLine(renderer, level, dsl.Element, dsl.Space, generateStringIdentifier(e.Tag().String()), dsl.Space, dsl.OpenBracket)
	
	// Basic properties
	if e.Shape() != nil {
		writeLine(renderer, level+1, dsl.Shape, dsl.Space, e.Shape().String())
	}
	if e.Icon() != nil {
		writeLine(renderer, level+1, dsl.Icon, dsl.Space, generateStringIdentifier(*e.Icon()))
	}
	if e.Width() != nil {
		writeLine(renderer, level+1, dsl.Width, dsl.Space, fmt.Sprintf("%d", *e.Width()))
	}
	if e.Height() != nil {
		writeLine(renderer, level+1, dsl.Height, dsl.Space, fmt.Sprintf("%d", *e.Height()))
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
		writeLine(renderer, level+1, dsl.Opacity, dsl.Space, fmt.Sprintf("%d", *e.Opacity()))
	}
	if e.FontSize() != nil {
		writeLine(renderer, level+1, dsl.FontSize, dsl.Space, fmt.Sprintf("%d", *e.FontSize()))
	}
	if e.Metadata() != nil {
		writeLine(renderer, level+1, dsl.Metadata, dsl.Space, fmt.Sprintf("%v", *e.Metadata()))
	}
	if e.Description() != nil {
		writeLine(renderer, level+1, dsl.Description, dsl.Space, fmt.Sprintf("%v", *e.Description()))
	}
	
	// Advanced properties
	if e.Border() != nil {
		writeLine(renderer, level+1, dsl.Border, dsl.Space, fmt.Sprintf("%d", *e.Border()))
	}
	if e.BorderStyle() != nil {
		writeLine(renderer, level+1, dsl.BorderStyle, dsl.Space, generateStringIdentifier(string(*e.BorderStyle())))
	}
	if e.StrokeWidth() != nil {
		writeLine(renderer, level+1, dsl.StrokeWidth, dsl.Space, fmt.Sprintf("%d", *e.StrokeWidth()))
	}
	if e.FontFamily() != nil {
		writeLine(renderer, level+1, dsl.FontFamily, dsl.Space, generateStringIdentifier(string(*e.FontFamily())))
	}
	if e.FontStyle() != nil {
		writeLine(renderer, level+1, dsl.FontStyle, dsl.Space, generateStringIdentifier(*e.FontStyle()))
	}
	if e.Shadow() != nil {
		writeLine(renderer, level+1, dsl.Shadow, dsl.Space, fmt.Sprintf("%t", *e.Shadow()))
	}
	if e.ZIndex() != nil {
		writeLine(renderer, level+1, dsl.ZIndex, dsl.Space, fmt.Sprintf("%d", *e.ZIndex()))
	}
	if e.Rotation() != nil {
		writeLine(renderer, level+1, dsl.Rotation, dsl.Space, fmt.Sprintf("%d", *e.Rotation()))
	}
	if e.Position() != nil {
		writeLine(renderer, level+1, dsl.Position, dsl.Space, fmt.Sprintf("%d,%d", (*e.Position())[0], (*e.Position())[1]))
	}
	if len(e.MultipleIcons()) > 0 {
		icons := strings.Join(e.MultipleIcons(), "\", \"")
		writeLine(renderer, level+1, dsl.Icons, dsl.Space, fmt.Sprintf("[\"%s\"]", icons))
	}
	
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

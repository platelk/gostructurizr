package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderViewStyles(s *gostructurizr.StylesNode, renderer *strings.Builder, level int) error {
	writeLine(renderer, level, dsl.Styles, dsl.Space, dsl.OpenBracket)
	
	// Render element styles
	for _, e := range s.ElementsStyle() {
		if err := renderViewElementStyle(e, renderer, level+1); err != nil {
			return fmt.Errorf("can't render element style: %w", err)
		}
	}
	
	// Render advanced relationship styles
	for _, r := range s.AdvancedRelationships() {
		if err := renderAdvancedRelationshipStyle(r, renderer, level+1); err != nil {
			return fmt.Errorf("can't render advanced relationship style: %w", err)
		}
	}
	
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

// renderAdvancedRelationshipStyle renders an advanced relationship style to DSL
func renderAdvancedRelationshipStyle(style *gostructurizr.AdvancedRelationshipStyleNode, renderer *strings.Builder, level int) error {
	writeLine(renderer, level, dsl.Element, dsl.Space, generateStringIdentifier(style.Tag().String()), dsl.Space, dsl.OpenBracket)

	// Basic properties
	if style.Color() != nil {
		writeLine(renderer, level+1, dsl.Color, dsl.Space, generateStringIdentifier(*style.Color()))
	}
	
	if style.Opacity() != nil {
		writeLine(renderer, level+1, dsl.Opacity, dsl.Space, fmt.Sprintf("%d", *style.Opacity()))
	}
	
	if style.Width() != nil {
		writeLine(renderer, level+1, dsl.Thickness, dsl.Space, fmt.Sprintf("%d", *style.Width()))
	}
	
	// Advanced properties
	if style.LineStyle() != nil {
		writeLine(renderer, level+1, dsl.Style, dsl.Space, generateStringIdentifier(string(*style.LineStyle())))
	}
	
	if style.FontSize() != nil {
		writeLine(renderer, level+1, dsl.FontSize, dsl.Space, fmt.Sprintf("%d", *style.FontSize()))
	}
	
	if style.FontColor() != nil {
		writeLine(renderer, level+1, dsl.FontColor, dsl.Space, generateStringIdentifier(*style.FontColor()))
	}
	
	if style.FontFamily() != nil {
		writeLine(renderer, level+1, dsl.FontFamily, dsl.Space, generateStringIdentifier(string(*style.FontFamily())))
	}
	
	if style.FontStyle() != nil {
		writeLine(renderer, level+1, dsl.FontStyle, dsl.Space, generateStringIdentifier(*style.FontStyle()))
	}
	
	if style.Routing() != nil {
		writeLine(renderer, level+1, dsl.Routing, dsl.Space, generateStringIdentifier(string(*style.Routing())))
	}
	
	if style.Position() != nil {
		writeLine(renderer, level+1, dsl.Position, dsl.Space, fmt.Sprintf("%d", *style.Position()))
	}
	
	if style.StartTerminator() != nil {
		writeLine(renderer, level+1, dsl.SourceTerminator, dsl.Space, generateStringIdentifier(string(*style.StartTerminator())))
	}
	
	if style.EndTerminator() != nil {
		writeLine(renderer, level+1, dsl.DestTerminator, dsl.Space, generateStringIdentifier(string(*style.EndTerminator())))
	}

	writeLine(renderer, level, dsl.CloseBracket)
	
	return nil
}

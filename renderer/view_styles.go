package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderViewStyles(s *gostructurizr.StylesNode, renderer *strings.Builder, level int) error {
	writeLine(renderer, level, dsl.Styles, dsl.Space, dsl.OpenBracket)
	for _, e := range s.ElementsStyle() {
		if err := renderViewElementStyle(e, renderer, level+1); err != nil {
			return fmt.Errorf("can't render element style: %w", err)
		}
	}
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderView(v *gostructurizr.ViewsNode, renderer *strings.Builder, level int) error {
	writeLine(renderer, level, dsl.Views, dsl.Space, dsl.OpenBracket)
	for _, s := range v.SystemContextViews() {
		if err := renderSystemContext(s, renderer, level+1); err != nil {
			return fmt.Errorf("can't generate system context view: %w", err)
		}
	}
	if err := renderViewConfiguration(v.Configuration(), renderer, level+1); err != nil {
		return fmt.Errorf("can't render view configuration: %w", err)
	}
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

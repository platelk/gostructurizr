package renderer

import (
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderInclude(e *gostructurizr.ExpressionViewNode, renderer *strings.Builder, level int) error {
	line := []string{dsl.Include, dsl.Space}
	if e.From() != nil {
		line = append(line, dsl.Space, generateVarName(e.From().Name()))
	}
	if e.Afferent() {
		if e.From() != nil {
			line = append(line, dsl.Space)
		}
		line = append(line, dsl.Arrow)
	}

	line = append(line, generateVarName(e.On().Name()))
	if e.Efferent() {
		if e.To() != nil {
			line = append(line, dsl.Space)
		}
		line = append(line, dsl.Arrow)
	}
	if e.To() != nil {
		line = append(line, dsl.Space, generateVarName(e.To().Name()))
	}
	writeLine(renderer, level, line...)
	return nil
}

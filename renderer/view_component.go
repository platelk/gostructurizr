package renderer

import (
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderViewComponent(c *gostructurizr.ComponentsViewNode, renderer *strings.Builder, level int) error {
	var line []string
	line = append(line, dsl.Container, dsl.Space, generateVarName(c.Container().Name()))
	if c.Key() != nil && *c.Key() != "" {
		line = append(line, dsl.Space, generateStringIdentifier(*c.Key()))
	}
	if c.Description() != nil && *c.Description() != "" {
		line = append(line, dsl.Space, generateStringIdentifier(*c.Description()))
	}
	line = append(line, dsl.OpenBracket)
	writeLine(renderer, level, line...)
	if c.IsAllElements() && c.IsAllPeople() {
		writeLine(renderer, level+1, dsl.Include, dsl.Space, dsl.All)
	}
	if c.AutoLayout() {
		writeLine(renderer, level+1, dsl.AutoLayout)
	}
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

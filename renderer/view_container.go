package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderViewContainer(c *gostructurizr.ContainersViewNode, renderer *strings.Builder, level int) error {
	var line []string
	line = append(line, dsl.Container, dsl.Space, generateVarName(c.SoftwareSystem().Name()))
	if c.Key() != nil && *c.Key() != "" {
		line = append(line, dsl.Space, generateStringIdentifier(*c.Key()))
	}
	if c.Description() != nil && *c.Description() != "" {
		line = append(line, dsl.Space, generateStringIdentifier(*c.Description()))
	}
	line = append(line, dsl.Space, dsl.OpenBracket)
	writeLine(renderer, level, line...)
	if c.IsAllElements() {
		writeLine(renderer, level+1, dsl.Include, dsl.Space, dsl.All)
	}
	for _, e := range c.Includes() {
		if err := renderInclude(e, renderer, level+1); err != nil {
			return fmt.Errorf("can't render include: %w", err)
		}
	}
	if c.AutoLayout() {
		writeLine(renderer, level+1, dsl.AutoLayout)
	}
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderContainer(c *gostructurizr.ContainerNode, renderer *strings.Builder, level int) error {
	var line []string
	line = append(line, generateVarName(c.Name()), dsl.Space, dsl.Equal, dsl.Space, dsl.Container, dsl.Space, generateStringIdentifier(c.Name()))
	if c.Description() != nil {
		line = append(line, dsl.Space, generateStringIdentifier(*c.Description()))
	}
	if c.Technology() != nil {
		if c.Description() == nil {
			line = append(line, dsl.Space, dsl.EmptyIdentifier)
		}
		line = append(line, dsl.Space, generateStringIdentifier(*c.Technology()))
	}
	components := c.Components()
	if (c.Tags() == nil || len(c.Tags().List()) == 0) && (components == nil || len(components) == 0) {
		writeLine(renderer, level, line...)
		return nil
	}
	line = append(line, dsl.Space, dsl.OpenBracket)
	writeLine(renderer, level, line...)
	if c.Tags() != nil && len(c.Tags().List()) > 0 {
		indent := strings.Repeat("    ", level+1)
		tagList := strings.Join(c.Tags().List(), ", ")
		fmt.Fprintf(renderer, "%s%s %q\n", indent, dsl.Tags, tagList)
	}
	for _, component := range components {
		if err := renderComponent(component, renderer, level+1); err != nil {
			return fmt.Errorf("can't render component: %w", err)
		}
	}
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

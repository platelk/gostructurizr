package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderComponent(c *gostructurizr.ComponentNode, renderer *strings.Builder, level int) error {
	var line []string
	line = append(line, generateVarName(c.Name()), dsl.Space, dsl.Equal, dsl.Space, dsl.Component, dsl.Space, generateStringIdentifier(c.Name()))
	if c.Description() != nil {
		line = append(line, dsl.Space, generateStringIdentifier(*c.Description()))
	}
	if c.Tags() == nil || len(c.Tags().List()) == 0 {
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
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

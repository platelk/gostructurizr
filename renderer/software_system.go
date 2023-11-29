package renderer

import (
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderSoftwareSystem(s *gostructurizr.SoftwareSystemNode, renderer *strings.Builder, level int) error {
	var line []string
	line = append(line, generateVarName(s.Name()), dsl.Space, dsl.Equal, dsl.Space, dsl.SoftwareSystem, dsl.Space, generateStringIdentifier(s.Name()))
	if s.Description() != nil {
		line = append(line, dsl.Space, generateStringIdentifier(*s.Description()))
	}
	writeLine(renderer, level, line...)
	return nil
}

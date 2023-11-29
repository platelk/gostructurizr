package renderer

import (
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderSystemContext(s *gostructurizr.SystemContextViewNode, renderer *strings.Builder, level int) error {
	var line []string
	line = append(line, dsl.SystemContext, dsl.Space, generateVarName(s.SoftwareSystem().Name()))
	if s.Key() != nil && *s.Key() != "" {
		line = append(line, dsl.Space, generateStringIdentifier(*s.Key()))
	}
	if s.Description() != nil && *s.Description() != "" {
		line = append(line, dsl.Space, generateStringIdentifier(*s.Description()))
	}
	line = append(line, dsl.OpenBracket)
	writeLine(renderer, level, line...)
	if s.IsAllSoftwareSystem() && s.IsAllPeople() {
		writeLine(renderer, level+1, dsl.Include, dsl.Space, dsl.All)
	}
	if s.AutoLayout() {
		writeLine(renderer, level+1, dsl.AutoLayout)
	}
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

package renderer

import (
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderPerson(p *gostructurizr.PersonNode, renderer *strings.Builder, level int) error {
	var line []string
	line = append(line, generateVarName(p.Name()), dsl.Space, dsl.Equal, dsl.Space, dsl.Person, dsl.Space, generateStringIdentifier(p.Name()))
	if p.Description() != nil {
		line = append(line, dsl.Space, generateStringIdentifier(*p.Description()))
	}
	if p.Tags() != nil && p.Tags().String() != "" {
		line = append(line, dsl.Space, generateStringIdentifier(p.Tags().String()))
	}
	writeLine(renderer, level, line...)
	return nil
}

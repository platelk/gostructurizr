package renderer

import (
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderRelationShip(r *gostructurizr.RelationShipNode, renderer *strings.Builder, level int) error {
	var line []string

	line = append(line, generateVarName(r.From().Name()), dsl.Space, dsl.Arrow, dsl.Space, generateVarName(r.To().Name()))
	if r.Description() != nil {
		line = append(line, dsl.Space, generateStringIdentifier(*r.Description()))
	}
	writeLine(renderer, level, line...)

	return nil
}

package renderer

import (
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderTags(p *gostructurizr.TagsNode, renderer *strings.Builder, level int) error {
	line := []string{dsl.Tags, dsl.Space}
	for _, tag := range p.Values() {
		line = append(line, generateStringIdentifier(tag))
	}
	writeLine(renderer, level, line...)
	return nil
}

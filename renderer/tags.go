package renderer

import (
	"strings"

	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
)

// RenderTagsV1 is the legacy renderer for tags
func RenderTagsV1(p *gostructurizr.TagsNode, renderer *strings.Builder, level int) error {
	line := []string{dsl.Tags, dsl.Space}
	for _, tag := range p.List() {
		line = append(line, generateStringIdentifier(tag))
	}
	writeLine(renderer, level, line...)
	return nil
}

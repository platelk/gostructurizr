package gostructurizr

import (
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

type TagsNode struct {
	values []string
}

func Tags() *TagsNode {
	return &TagsNode{}
}

func (t *TagsNode) Add(s string) *TagsNode {
	t.values = append(t.values, s)
	return t
}

func (t *TagsNode) String() string {
	return strings.Join(t.values, dsl.TagSeparator)
}

func (t *TagsNode) Values() []string {
	return t.values
}

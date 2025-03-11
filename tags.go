package gostructurizr

import (
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

// TagsNode represents a collection of tags for an element
type TagsNode struct {
	Tags []string
}

// NewTags creates a new TagsNode
func NewTags() TagsNode {
	return TagsNode{
		Tags: []string{},
	}
}

// Add adds a tag to the node
func (t *TagsNode) Add(s string) *TagsNode {
	t.Tags = append(t.Tags, s)
	return t
}

// String returns a string representation of all tags
func (t *TagsNode) String() string {
	return strings.Join(t.Tags, dsl.TagSeparator)
}

// List returns all tags as a slice
func (t *TagsNode) List() []string {
	return t.Tags
}

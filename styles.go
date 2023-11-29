package gostructurizr

import (
	"github.com/platelk/gostructurizr/tags"
)

type StylesNode struct {
	elements []*ElementStyleNode
}

func styles() *StylesNode {
	return &StylesNode{}
}

func (s *StylesNode) AddElementStyle(tag tags.Tag) *ElementStyleNode {
	e := elementStyle(tag)
	s.elements = append(s.elements, e)
	return e
}

func (s *StylesNode) ElementsStyle() []*ElementStyleNode {
	return s.elements
}

func (s *StylesNode) AddRelationshipStyle(ship tags.Tag) *RelationShipStyleNode {
	return &RelationShipStyleNode{}
}

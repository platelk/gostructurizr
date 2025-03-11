package gostructurizr

import (
	"github.com/platelk/gostructurizr/tags"
)

type StylesNode struct {
	elements               []*ElementStyleNode
	advancedRelationships  []*AdvancedRelationshipStyleNode
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

// AddAdvancedRelationshipStyle adds an advanced relationship style for a tag
func (s *StylesNode) AddAdvancedRelationshipStyle(tag tags.Tag) *AdvancedRelationshipStyleNode {
	r := AdvancedRelationshipStyle(tag)
	s.advancedRelationships = append(s.advancedRelationships, r)
	return r
}

// AdvancedRelationships returns all advanced relationship styles
func (s *StylesNode) AdvancedRelationships() []*AdvancedRelationshipStyleNode {
	return s.advancedRelationships
}

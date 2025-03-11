package gostructurizr

import (
	"github.com/platelk/gostructurizr/tags"
)

// InfrastructureNodeNode represents an infrastructure node in the architecture
type InfrastructureNodeNode struct {
	name       string
	desc       string
	technology string
	tags       TagsNode
	properties Properties
	model      *ModelNode
	parent     *DeploymentNodeNode
}

// InfrastructureNode creates a new InfrastructureNodeNode
func InfrastructureNode(name, desc, technology string) *InfrastructureNodeNode {
	node := &InfrastructureNodeNode{
		name:       name,
		desc:       desc,
		technology: technology,
		tags:       TagsNode{Tags: []string{}},
		properties: Properties{Properties: make(map[string]string)},
	}
	node.tags.Add(tags.InfrastructureNode.String())
	return node
}

// Name returns the name of the infrastructure node
func (i *InfrastructureNodeNode) Name() string {
	return i.name
}

// WithDesc sets the description of the infrastructure node
func (i *InfrastructureNodeNode) WithDesc(desc string) *InfrastructureNodeNode {
	i.desc = desc
	return i
}

// Uses creates a relationship from this infrastructure node to another element
func (i *InfrastructureNodeNode) Uses(toNode Namer, desc string) *RelationShipNode {
	return i.model.addRelationShip(i, toNode, desc)
}

// Tags returns the tags of the infrastructure node
func (i *InfrastructureNodeNode) Tags() *TagsNode {
	return &i.tags
}

// WithTag adds a tag to the infrastructure node
func (i *InfrastructureNodeNode) WithTag(tag string) *InfrastructureNodeNode {
	i.tags.Add(tag)
	return i
}

// Properties returns the properties of the infrastructure node
func (i *InfrastructureNodeNode) Properties() *Properties {
	return &i.properties
}

// Technology returns the technology of the infrastructure node
func (i *InfrastructureNodeNode) Technology() string {
	return i.technology
}

// WithTechnology sets the technology of the infrastructure node
func (i *InfrastructureNodeNode) WithTechnology(technology string) *InfrastructureNodeNode {
	i.technology = technology
	return i
}

// Description returns the description of the infrastructure node
func (i *InfrastructureNodeNode) Description() string {
	return i.desc
}
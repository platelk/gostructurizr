package gostructurizr

import (
	"github.com/platelk/gostructurizr/tags"
)

// DeploymentEnvironment represents the environment type for deployment nodes
type DeploymentEnvironment string

const (
	DevelopmentEnvironment DeploymentEnvironment = "Development"
	TestEnvironment        DeploymentEnvironment = "Test"
	StagingEnvironment     DeploymentEnvironment = "Staging"
	ProductionEnvironment  DeploymentEnvironment = "Production"
)

// Location represents the location type for elements
type Location string

const (
	InternalLocation Location = "Internal"
	ExternalLocation Location = "External"
)

// DeploymentNodeNode represents a deployment node in the architecture
type DeploymentNodeNode struct {
	name                string
	desc                string
	technology          string
	environment         DeploymentEnvironment
	location            Location
	tags                TagsNode
	properties          Properties
	model               *ModelNode
	parent              *DeploymentNodeNode
	children            []*DeploymentNodeNode
	infrastructureNodes []*InfrastructureNodeNode
	containerInstances  []*ContainerInstanceNode
}

// Description returns the description of the deployment node
func (d *DeploymentNodeNode) Description() string {
	return d.desc
}

// Technology returns the technology of the deployment node
func (d *DeploymentNodeNode) Technology() string {
	return d.technology
}

// Environment returns the environment of the deployment node
func (d *DeploymentNodeNode) Environment() DeploymentEnvironment {
	return d.environment
}

// WithEnv sets the deployment environment of the deployment node and all its children
func (d *DeploymentNodeNode) WithEnv(environment DeploymentEnvironment) *DeploymentNodeNode {
	d.environment = environment
	
	// Update environment for all children
	for _, child := range d.children {
		child.WithEnv(environment)
	}
	
	return d
}

// WithTechnology sets the technology of the deployment node
func (d *DeploymentNodeNode) WithTechnology(technology string) *DeploymentNodeNode {
	d.technology = technology
	return d
}

// DeploymentNode creates a new DeploymentNodeNode
func DeploymentNode(name, desc, technology string, environment DeploymentEnvironment) *DeploymentNodeNode {
	node := &DeploymentNodeNode{
		name:        name,
		desc:        desc,
		technology:  technology,
		environment: environment,
		location:    InternalLocation,
		tags:        TagsNode{Tags: []string{}},
		properties:  Properties{Properties: make(map[string]string)},
	}
	node.tags.Add(tags.DeploymentNode.String())
	return node
}

// Name returns the name of the deployment node
func (d *DeploymentNodeNode) Name() string {
	return d.name
}

// WithDesc sets the description of the deployment node
func (d *DeploymentNodeNode) WithDesc(desc string) *DeploymentNodeNode {
	d.desc = desc
	return d
}

// WithLocation sets the location of the deployment node
func (d *DeploymentNodeNode) WithLocation(location Location) *DeploymentNodeNode {
	d.location = location
	return d
}

// Uses creates a relationship from this deployment node to another element
func (d *DeploymentNodeNode) Uses(toNode Namer, desc string) *RelationShipNode {
	return d.model.addRelationShip(d, toNode, desc)
}

// Tags returns the tags of the deployment node
func (d *DeploymentNodeNode) Tags() *TagsNode {
	return &d.tags
}

// WithTag adds a tag to the deployment node
func (d *DeploymentNodeNode) WithTag(tag string) *DeploymentNodeNode {
	d.tags.Add(tag)
	return d
}

// Properties returns the properties of the deployment node
func (d *DeploymentNodeNode) Properties() *Properties {
	return &d.properties
}

// Add adds a new child deployment node
func (d *DeploymentNodeNode) Add(child *DeploymentNodeNode) *DeploymentNodeNode {
	d.children = append(d.children, child)
	child.parent = d
	child.model = d.model
	return d
}

// Children returns the child deployment nodes
func (d *DeploymentNodeNode) Children() []*DeploymentNodeNode {
	return d.children
}

// AddInfrastructureNode adds a new infrastructure node to this deployment node
func (d *DeploymentNodeNode) AddInfrastructureNode(name, desc, technology string) *InfrastructureNodeNode {
	infra := InfrastructureNode(name, desc, technology)
	d.infrastructureNodes = append(d.infrastructureNodes, infra)
	infra.parent = d
	infra.model = d.model
	return infra
}

// InfrastructureNodes returns the infrastructure nodes
func (d *DeploymentNodeNode) InfrastructureNodes() []*InfrastructureNodeNode {
	return d.infrastructureNodes
}

// AddDeploymentNode adds a new child deployment node with specified environment
func (d *DeploymentNodeNode) AddDeploymentNode(name, desc, technology string, environment DeploymentEnvironment) *DeploymentNodeNode {
	child := DeploymentNode(name, desc, technology, environment)
	d.Add(child)
	return child
}

// AddChildNode adds a new child deployment node with parent's environment
func (d *DeploymentNodeNode) AddChildNode(name, desc, technology string) *DeploymentNodeNode {
	child := DeploymentNode(name, desc, technology, d.environment)
	d.Add(child)
	return child
}

// AddContainerInstance adds a container instance to this deployment node
func (d *DeploymentNodeNode) AddContainerInstance(container *ContainerNode) *ContainerInstanceNode {
	instance := ContainerInstance(container)
	d.containerInstances = append(d.containerInstances, instance)
	instance.parent = d
	instance.model = d.model
	return instance
}

// ContainerInstances returns the container instances
func (d *DeploymentNodeNode) ContainerInstances() []*ContainerInstanceNode {
	return d.containerInstances
}
package gostructurizr

import (
	"github.com/platelk/gostructurizr/tags"
)

// ContainerInstanceNode represents an instance of a container in a deployment environment
type ContainerInstanceNode struct {
	container      *ContainerNode
	instanceId     int
	tags           TagsNode
	properties     Properties
	model          *ModelNode
	parent         *DeploymentNodeNode
	healthChecks   []*HealthCheckNode
}

// ContainerInstance creates a new ContainerInstanceNode
func ContainerInstance(container *ContainerNode) *ContainerInstanceNode {
	node := &ContainerInstanceNode{
		container:  container,
		instanceId: 1, // Default instance ID
		tags:       TagsNode{Tags: []string{}},
		properties: Properties{Properties: make(map[string]string)},
	}
	node.tags.Add(tags.ContainerInstance.String())
	return node
}

// Name returns the name of the container instance
func (c *ContainerInstanceNode) Name() string {
	return c.container.Name()
}

// Container returns the referenced container
func (c *ContainerInstanceNode) Container() *ContainerNode {
	return c.container
}

// WithInstanceId sets the instance ID of the container instance
func (c *ContainerInstanceNode) WithInstanceId(id int) *ContainerInstanceNode {
	c.instanceId = id
	return c
}

// InstanceId returns the instance ID
func (c *ContainerInstanceNode) InstanceId() int {
	return c.instanceId
}

// Uses creates a relationship from this container instance to another element
func (c *ContainerInstanceNode) Uses(toNode Namer, desc string) *RelationShipNode {
	return c.model.addRelationShip(c, toNode, desc)
}

// Tags returns the tags of the container instance
func (c *ContainerInstanceNode) Tags() *TagsNode {
	return &c.tags
}

// WithTag adds a tag to the container instance
func (c *ContainerInstanceNode) WithTag(tag string) *ContainerInstanceNode {
	c.tags.Add(tag)
	return c
}

// Properties returns the properties of the container instance
func (c *ContainerInstanceNode) Properties() *Properties {
	return &c.properties
}

// AddHealthCheck adds a health check to this container instance
func (c *ContainerInstanceNode) AddHealthCheck(name, url string) *HealthCheckNode {
	healthCheck := HealthCheck(name, url)
	c.healthChecks = append(c.healthChecks, healthCheck)
	healthCheck.parent = c
	return healthCheck
}

// HealthChecks returns the health checks for this container instance
func (c *ContainerInstanceNode) HealthChecks() []*HealthCheckNode {
	return c.healthChecks
}
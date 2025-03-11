package gostructurizr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeploymentNode(t *testing.T) {
	// Create a deployment node
	node := DeploymentNode("Amazon Web Services", "Cloud platform", "AWS", ProductionEnvironment)
	
	// Test basic properties
	assert.Equal(t, "Amazon Web Services", node.Name())
	assert.Equal(t, "Cloud platform", node.Description())
	assert.Equal(t, "AWS", node.Technology())
	assert.Equal(t, ProductionEnvironment, node.Environment())
	assert.Equal(t, InternalLocation, node.location)
	
	// Test fluent interface
	node.WithDesc("AWS Cloud platform")
	assert.Equal(t, "AWS Cloud platform", node.Description())
	
	node.WithLocation(ExternalLocation)
	assert.Equal(t, ExternalLocation, node.location)
	
	node.WithTechnology("Amazon Web Services")
	assert.Equal(t, "Amazon Web Services", node.Technology())
	
	// Test tags
	assert.Contains(t, node.Tags().List(), "Deployment Node")
	
	node.WithTag("custom-tag")
	assert.Contains(t, node.Tags().List(), "custom-tag")
	
	// Test properties
	node.Properties().Add("region", "us-west-1")
	assert.Equal(t, "us-west-1", node.Properties().Get("region"))
}

func TestDeploymentNodeHierarchy(t *testing.T) {
	// Create a model
	m := Model()
	
	// Add a top-level deployment node
	aws := m.AddDeploymentNode("AWS", "Amazon Web Services", "Amazon Web Services", ProductionEnvironment)
	
	// Add a child deployment node
	usEast := aws.AddDeploymentNode("US-East", "US East Region", "AWS Region", ProductionEnvironment)
	
	// Test parent-child relationship
	assert.Equal(t, aws, usEast.parent)
	assert.Equal(t, m, aws.model)
	assert.Equal(t, m, usEast.model)
	assert.Contains(t, aws.Children(), usEast)
	
	// Add infrastructure node
	rds := usEast.AddInfrastructureNode("RDS", "Relational Database Service", "Amazon RDS")
	
	// Test infrastructure node
	assert.Equal(t, "RDS", rds.Name())
	assert.Equal(t, usEast, rds.parent)
	assert.Equal(t, m, rds.model)
	assert.Contains(t, usEast.InfrastructureNodes(), rds)
	
	// Create software system and container
	system := m.AddSoftwareSystem("Banking System", "Online Banking Application")
	webapp := system.AddContainer("Web Application", "Web frontend", "Java and Spring MVC")
	
	// Add container instance
	webappInstance := usEast.AddContainerInstance(webapp)
	
	// Test container instance
	assert.Equal(t, webapp, webappInstance.Container())
	assert.Equal(t, usEast, webappInstance.parent)
	assert.Equal(t, m, webappInstance.model)
	assert.Contains(t, usEast.ContainerInstances(), webappInstance)
}
package gostructurizr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironmentHandling(t *testing.T) {
	// Create a model
	m := Model()
	system := m.AddSoftwareSystem("Test System", "A test system")
	
	// Test environment creation through specialized methods
	devNode := m.AddDevNode("Dev Node", "Development Environment", "Docker")
	testNode := m.AddTestNode("Test Node", "Test Environment", "Kubernetes")
	stagingNode := m.AddStagingNode("Staging Node", "Staging Environment", "AWS")
	prodNode := m.AddProdNode("Prod Node", "Production Environment", "AWS")
	
	// Test environment setting
	assert.Equal(t, DevelopmentEnvironment, devNode.Environment())
	assert.Equal(t, TestEnvironment, testNode.Environment())
	assert.Equal(t, StagingEnvironment, stagingNode.Environment())
	assert.Equal(t, ProductionEnvironment, prodNode.Environment())
	
	// Test WithEnv method
	devNode.WithEnv(StagingEnvironment)
	assert.Equal(t, StagingEnvironment, devNode.Environment())
	
	// Test environment inheritance for child nodes
	child := devNode.AddChildNode("Child Node", "Child of Dev", "Docker")
	assert.Equal(t, StagingEnvironment, child.Environment())
	
	// Test environment propagation to children
	devNode.WithEnv(ProductionEnvironment)
	assert.Equal(t, ProductionEnvironment, child.Environment())
	
	// Test deployment view environment methods
	workspace := Workspace()
	views := workspace.Views()
	
	devView := views.CreateDevView(system)
	testView := views.CreateTestView(system)
	stagingView := views.CreateStagingView(system)
	prodView := views.CreateProdView(system)
	
	assert.Equal(t, DevelopmentEnvironment, devView.Environment())
	assert.Equal(t, TestEnvironment, testView.Environment())
	assert.Equal(t, StagingEnvironment, stagingView.Environment())
	assert.Equal(t, ProductionEnvironment, prodView.Environment())
	
	// Test environment change methods
	devView.ForProduction()
	assert.Equal(t, ProductionEnvironment, devView.Environment())
	
	testView.WithEnv(StagingEnvironment)
	assert.Equal(t, StagingEnvironment, testView.Environment())
}
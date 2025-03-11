package gostructurizr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeploymentView(t *testing.T) {
	// Create a workspace
	workspace := Workspace().WithName("Deployment Test")
	model := workspace.Model()
	
	// Create a software system
	system := model.AddSoftwareSystem("Banking System", "Online Banking Application")
	webapp := system.AddContainer("Web Application", "Web frontend", "Java and Spring MVC")
	api := system.AddContainer("API Application", "JSON API", "Java and Spring Boot")
	db := system.AddContainer("Database", "Stores user data", "Oracle")
	
	// Create deployments
	awsCloud := model.AddDeploymentNode("AWS", "Amazon Web Services", "Amazon Web Services", ProductionEnvironment)
	
	// Add web tier
	webTier := awsCloud.AddDeploymentNode("Web Tier", "Web and API Tier", "Amazon EC2", ProductionEnvironment)
	webAppInstance := webTier.AddContainerInstance(webapp)
	apiInstance := webTier.AddContainerInstance(api)
	
	// Add database tier
	dbTier := awsCloud.AddDeploymentNode("DB Tier", "Database Tier", "Amazon RDS", ProductionEnvironment)
	dbInstance := dbTier.AddContainerInstance(db)
	
	// Create relationships
	webAppInstance.Uses(apiInstance, "Calls API")
	apiInstance.Uses(dbInstance, "Reads from and writes to")
	
	// Create deployment view
	views := workspace.Views()
	deploymentView := views.CreateDeploymentView(system, ProductionEnvironment)
	deploymentView.WithKey("AWSDeployment")
	deploymentView.WithDescription("AWS Deployment for the Banking System")
	deploymentView.WithAutoLayout()
	
	// Add nodes to the view
	deploymentView.AddDeploymentNode(awsCloud)
	
	// Add all relationships
	deploymentView.AddAllRelationships()
	
	// Test deployment view properties
	assert.Equal(t, system, deploymentView.SoftwareSystem())
	assert.Equal(t, ProductionEnvironment, deploymentView.Environment())
	assert.Equal(t, "AWSDeployment", deploymentView.GetKey())
	assert.Equal(t, "AWS Deployment for the Banking System", deploymentView.GetDescription())
	assert.True(t, deploymentView.IsAutoLayout())
}
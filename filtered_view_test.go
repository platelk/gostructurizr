package gostructurizr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilteredView(t *testing.T) {
	// Create a model
	workspace := Workspace().WithName("Filtered View Test")
	model := workspace.Model()
	
	// Create a software system
	system := model.AddSoftwareSystem("Banking System", "Core banking system")
	system.WithTag("Banking")
	
	// Create containers with different tags
	webApp := system.AddContainer("Web Application", "Web frontend", "Java and Spring MVC")
	webApp.WithTag("Web")
	
	apiApp := system.AddContainer("API Application", "REST API", "Java and Spring Boot")
	apiApp.WithTag("API")
	apiApp.WithTag("Service")
	
	database := system.AddContainer("Database", "Stores user data", "Oracle")
	database.WithTag("Database")
	
	// Create a base container view
	views := workspace.Views()
	containerView := views.CreateContainerView(system)
	containerView.WithKey("AllContainers")
	containerView.WithDescription("All containers")
	containerView.AddAllContainers()
	
	// Create a filtered view based on the container view
	filteredView := views.CreateFilteredView(containerView, "API Services Only")
	filteredView.WithKey("APIServices")
	filteredView.WithDescription("Shows only API services")
	filteredView.Include("API")
	filteredView.Exclude("Database")
	filteredView.WithAutoLayout()
	
	// Test filtered view properties
	assert.Equal(t, "API Services Only", filteredView.Title())
	assert.Equal(t, "APIServices", filteredView.Key())
	assert.Equal(t, "Shows only API services", filteredView.Description())
	assert.True(t, filteredView.IsAutoLayout())
	
	// Test filter criteria
	criteria := filteredView.FilterCriteria()
	assert.Equal(t, 2, len(criteria))
	
	assert.Equal(t, Include, criteria[0].Mode)
	assert.Equal(t, TagFilter, criteria[0].Type)
	assert.Equal(t, "API", criteria[0].Value)
	
	assert.Equal(t, Exclude, criteria[1].Mode)
	assert.Equal(t, TagFilter, criteria[1].Type)
	assert.Equal(t, "Database", criteria[1].Value)
	
	// Test chained methods
	filteredView2 := views.CreateFilteredView(containerView, "Web Only")
	filteredView2.WithKey("WebOnly").
		WithDescription("Shows only web containers").
		Include("Web").
		WithAutoLayout()
	
	assert.Equal(t, "Web Only", filteredView2.Title())
	assert.Equal(t, "WebOnly", filteredView2.Key())
	assert.Equal(t, "Shows only web containers", filteredView2.Description())
	
	// Test advanced filtering with name filter
	filteredView3 := views.CreateFilteredView(containerView, "API Name Filter")
	filteredView3.WithNameFilter("API", Include)
	
	criteria3 := filteredView3.FilterCriteria()
	assert.Equal(t, NameFilter, criteria3[0].Type)
	assert.Equal(t, "API", criteria3[0].Value)
}
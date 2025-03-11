package gostructurizr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnterprise(t *testing.T) {
	// Create a model
	m := Model()
	
	// Set enterprise
	enterprise := m.SetEnterprise("ACME Corporation")
	
	// Test enterprise properties
	assert.Equal(t, "ACME Corporation", enterprise.Name())
	assert.Equal(t, m, enterprise.model)
	
	// Test fluent interface
	enterprise.WithName("New Corp Name")
	assert.Equal(t, "New Corp Name", enterprise.Name())
	
	// Test getter
	assert.Equal(t, enterprise, m.Enterprise())
	
	// Test properties
	enterprise.Properties().Add("domain", "acme.com")
	assert.Equal(t, "acme.com", enterprise.Properties().Get("domain"))
}

func TestLocationConcepts(t *testing.T) {
	// Create a model with software systems
	m := Model()
	_ = m.AddSoftwareSystem("Internal CRM", "Customer Relationship Management")
	_ = m.AddSoftwareSystem("External Payment Gateway", "Processes payments")
	
	// Create deployment nodes with locations
	internalDC := m.AddDeploymentNode("Internal DC", "Internal datacenter", "On-premises", ProductionEnvironment)
	internalDC.WithLocation(InternalLocation)
	
	externalDC := m.AddDeploymentNode("External Cloud", "Cloud provider", "AWS", ProductionEnvironment)
	externalDC.WithLocation(ExternalLocation)
	
	// Test location properties
	assert.Equal(t, InternalLocation, internalDC.location)
	assert.Equal(t, ExternalLocation, externalDC.location)
}
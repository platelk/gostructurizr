package gostructurizr

// DeploymentViewNode represents a deployment view in the architecture
type DeploymentViewNode struct {
	ViewNode
	softwareSystem *SoftwareSystemNode
	environment    DeploymentEnvironment
}

// DeploymentView creates a new DeploymentViewNode
func DeploymentView(softwareSystem *SoftwareSystemNode, environment DeploymentEnvironment) *DeploymentViewNode {
	return &DeploymentViewNode{
		ViewNode:       *NewViewNode(),
		softwareSystem: softwareSystem,
		environment:    environment,
	}
}

// SoftwareSystem returns the software system for this deployment view
func (d *DeploymentViewNode) SoftwareSystem() *SoftwareSystemNode {
	return d.softwareSystem
}

// WithSoftwareSystem sets the software system for this deployment view
func (d *DeploymentViewNode) WithSoftwareSystem(softwareSystem *SoftwareSystemNode) *DeploymentViewNode {
	d.softwareSystem = softwareSystem
	return d
}

// Environment returns the environment for this deployment view
func (d *DeploymentViewNode) Environment() DeploymentEnvironment {
	return d.environment
}

// WithEnvironment sets the environment for this deployment view
func (d *DeploymentViewNode) WithEnvironment(environment DeploymentEnvironment) *DeploymentViewNode {
	d.environment = environment
	return d
}

// WithEnv is an alias for WithEnvironment - sets the environment for this deployment view
func (d *DeploymentViewNode) WithEnv(environment DeploymentEnvironment) *DeploymentViewNode {
	return d.WithEnvironment(environment)
}

// ForDev sets the environment to Development
func (d *DeploymentViewNode) ForDev() *DeploymentViewNode {
	return d.WithEnvironment(DevelopmentEnvironment)
}

// ForTest sets the environment to Test
func (d *DeploymentViewNode) ForTest() *DeploymentViewNode {
	return d.WithEnvironment(TestEnvironment)
}

// ForStaging sets the environment to Staging
func (d *DeploymentViewNode) ForStaging() *DeploymentViewNode {
	return d.WithEnvironment(StagingEnvironment)
}

// ForProduction sets the environment to Production
func (d *DeploymentViewNode) ForProduction() *DeploymentViewNode {
	return d.WithEnvironment(ProductionEnvironment)
}

// Add methods from the parent ViewNode
func (d *DeploymentViewNode) WithKey(key string) *DeploymentViewNode {
	d.ViewNode.WithKey(key)
	return d
}

// WithDescription sets the description of the deployment view
func (d *DeploymentViewNode) WithDescription(desc string) *DeploymentViewNode {
	d.ViewNode.WithDescription(desc)
	return d
}

// WithAutoLayout sets the auto layout of the deployment view
func (d *DeploymentViewNode) WithAutoLayout() *DeploymentViewNode {
	d.ViewNode.WithAutoLayout()
	return d
}

// AddDeploymentNode adds a deployment node to the view
func (d *DeploymentViewNode) AddDeploymentNode(deploymentNode *DeploymentNodeNode) *DeploymentViewNode {
	d.AddElement(deploymentNode)
	return d
}

// AddAll adds all deployment nodes from the model
func (d *DeploymentViewNode) AddAll() *DeploymentViewNode {
	// This would require model to have a deploymentNodes field
	// Just a stub for now
	return d
}

// Add adds a deployment node to the view
func (d *DeploymentViewNode) Add(deploymentNode *DeploymentNodeNode) *DeploymentViewNode {
	return d.AddDeploymentNode(deploymentNode)
}

// AddRelationship adds a relationship to the view
func (d *DeploymentViewNode) AddRelationship(relationship *RelationShipNode) *DeploymentViewNode {
	d.ViewNode.AddRelationship(relationship)
	return d
}

// AddAllRelationships adds all relationships from the model
func (d *DeploymentViewNode) AddAllRelationships() *DeploymentViewNode {
	// This would add all relationships that connect elements in this view
	// Just a stub for now
	return d
}
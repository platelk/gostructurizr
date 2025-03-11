package gostructurizr

// EnterpriseNode represents an enterprise boundary in the architecture
type EnterpriseNode struct {
	name       string
	properties Properties
	model      *ModelNode
}

// Enterprise creates a new EnterpriseNode
func Enterprise(name string) *EnterpriseNode {
	return &EnterpriseNode{
		name:       name,
		properties: Properties{},
	}
}

// Name returns the name of the enterprise
func (e *EnterpriseNode) Name() string {
	return e.name
}

// WithName sets the name of the enterprise
func (e *EnterpriseNode) WithName(name string) *EnterpriseNode {
	e.name = name
	return e
}

// Properties returns the properties of the enterprise
func (e *EnterpriseNode) Properties() *Properties {
	return &e.properties
}
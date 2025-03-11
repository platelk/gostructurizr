package gostructurizr

// ViewNode contains common fields and methods for all view types
type ViewNode struct {
	key           string
	description   string
	autoLayout    bool
	elements      []Namer
	relationships []*RelationShipNode
}

// NewViewNode creates a new base ViewNode
func NewViewNode() *ViewNode {
	return &ViewNode{
		elements:      []Namer{},
		relationships: []*RelationShipNode{},
	}
}

// WithKey sets the key for this view
func (v *ViewNode) WithKey(key string) *ViewNode {
	v.key = key
	return v
}

// GetKey returns the key for this view
func (v *ViewNode) GetKey() string {
	return v.key
}

// WithDescription sets the description for this view
func (v *ViewNode) WithDescription(desc string) *ViewNode {
	v.description = desc
	return v
}

// GetDescription returns the description for this view
func (v *ViewNode) GetDescription() string {
	return v.description
}

// WithAutoLayout sets auto layout for this view
func (v *ViewNode) WithAutoLayout() *ViewNode {
	v.autoLayout = true
	return v
}

// IsAutoLayout returns whether auto layout is enabled
func (v *ViewNode) IsAutoLayout() bool {
	return v.autoLayout
}

// AddElement adds an element to this view
func (v *ViewNode) AddElement(element Namer) *ViewNode {
	v.elements = append(v.elements, element)
	return v
}

// Elements returns all elements in this view
func (v *ViewNode) Elements() []Namer {
	return v.elements
}

// AddRelationship adds a relationship to this view
func (v *ViewNode) AddRelationship(relationship *RelationShipNode) *ViewNode {
	v.relationships = append(v.relationships, relationship)
	return v
}

// RelationShips returns all relationships in this view
func (v *ViewNode) RelationShips() []*RelationShipNode {
	return v.relationships
}
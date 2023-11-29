package gostructurizr

type PersonNode struct {
	name        string
	description *string
	tags        *TagsNode
	model       *ModelNode
}

func Person(name, description string) *PersonNode {
	return &PersonNode{
		name:        name,
		description: &description,
		tags:        Tags(),
	}
}

func (p *PersonNode) Name() string {
	return p.name
}

func (p *PersonNode) Description() *string {
	return p.description
}

func (p *PersonNode) Tags() *TagsNode {
	return p.tags
}

func (p *PersonNode) Uses(to Namer, desc string) *RelationShipNode {
	return p.model.addRelationShip(p, to, desc)
}

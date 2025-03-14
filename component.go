package gostructurizr

type ComponentNode struct {
	node *ContainerNode
	name string
	desc *string
	tech *string
	tags *TagsNode
}

func Component(name string) *ComponentNode {
	return &ComponentNode{
		name: name,
		tags: &TagsNode{Tags: []string{}},
	}
}

func (c *ComponentNode) Name() string {
	return c.name
}

func (c *ComponentNode) WithDesc(desc string) *ComponentNode {
	c.desc = &desc
	return c
}

func (c *ComponentNode) Description() *string {
	return c.desc
}

func (c *ComponentNode) WithTechnology(tech string) *ComponentNode {
	c.tech = &tech
	return c
}

func (c *ComponentNode) Technology() *string {
	return c.tech
}

func (c *ComponentNode) WithTag(t string) *ComponentNode {
	c.tags.Add(t)
	return c
}

func (c *ComponentNode) Tags() *TagsNode {
	return c.tags
}

func (c *ComponentNode) Uses(to Namer, desc string) *RelationShipNode {
	return c.node.sys.model.addRelationShip(c, to, desc)
}

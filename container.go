package gostructurizr

type ContainerNode struct {
	sys        *SoftwareSystemNode
	name       string
	desc       *string
	tech       *string
	tags       *TagsNode
	components []*ComponentNode
}

func Container(name string) *ContainerNode {
	return &ContainerNode{
		name: name,
		tags: Tags(),
	}
}

func (c *ContainerNode) Name() string {
	return c.name
}

func (c *ContainerNode) WithDesc(desc string) *ContainerNode {
	c.desc = &desc
	return c
}

func (c *ContainerNode) Description() *string {
	return c.desc
}

func (c *ContainerNode) WithTechnology(tech string) *ContainerNode {
	c.tech = &tech
	return c
}

func (c *ContainerNode) Technology() *string {
	return c.tech
}

func (c *ContainerNode) WithTag(t string) *ContainerNode {
	c.tags.Add(t)
	return c
}

func (c *ContainerNode) Tags() *TagsNode {
	return c.tags
}

func (c *ContainerNode) AddComponent(name string) *ComponentNode {
	component := Component(name)
	component.node = c
	c.components = append(c.components, component)

	return component
}

func (c *ContainerNode) Components() []*ComponentNode {
	return c.components
}

func (c *ContainerNode) Uses(to Namer, desc string) *RelationShipNode {
	return c.sys.model.addRelationShip(c, to, desc)
}

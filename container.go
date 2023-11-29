package gostructurizr

type ContainerNode struct {
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

func (c *ContainerNode) AddComponent(name string) *ComponentNode {
	component := Component(name)
	c.components = append(c.components, component)

	return component
}

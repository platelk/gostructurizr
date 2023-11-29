package gostructurizr

type SoftwareSystemNode struct {
	model      *ModelNode
	name       string
	desc       *string
	containers []*ContainerNode
	tags       *TagsNode
}

func SoftwareSystem(name, desc string) *SoftwareSystemNode {
	return &SoftwareSystemNode{
		name: name,
		desc: &desc,
	}
}

func (s *SoftwareSystemNode) Name() string {
	return s.name
}

func (s *SoftwareSystemNode) Description() *string {
	return s.desc
}

func (s *SoftwareSystemNode) Uses(to Namer, desc string) *RelationShipNode {
	return s.model.addRelationShip(s, to, desc)
}

func (s *SoftwareSystemNode) AddContainer(name string) *ContainerNode {
	c := Container(name)
	c.sys = s
	s.containers = append(s.containers, c)

	return c
}

func (s *SoftwareSystemNode) Containers() []*ContainerNode {
	return s.containers
}

func (s *SoftwareSystemNode) WithTag(t string) *SoftwareSystemNode {
	s.tags.Add(t)
	return s
}

func (s *SoftwareSystemNode) Tags() *TagsNode {
	return s.tags
}

package gostructurizr

type SoftwareSystemNode struct {
	model      *ModelNode
	name       string
	desc       *string
	containers []*ContainerNode
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
	s.containers = append(s.containers, c)

	return c
}

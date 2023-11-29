package gostructurizr

type ModelNode struct {
	properties      Properties
	softwareGroups  []GroupNode[*SoftwareSystemNode]
	softwareSystems []*SoftwareSystemNode
	persons         []*PersonNode
	uses            []*RelationShipNode
}

func Model() *ModelNode {
	return &ModelNode{}
}

func (m *ModelNode) AddPerson(name, desc string) *PersonNode {
	p := Person(name, desc)
	m.persons = append(m.persons, p)
	p.model = m
	return p
}

func (m *ModelNode) Persons() []*PersonNode {
	return m.persons
}

func (m *ModelNode) AddSoftwareSystem(name, desc string) *SoftwareSystemNode {
	s := SoftwareSystem(name, desc)
	m.softwareSystems = append(m.softwareSystems, s)
	s.model = m
	return s
}

func (m *ModelNode) SoftwareSystems() []*SoftwareSystemNode {
	return m.softwareSystems
}

func (m *ModelNode) RelationShip() []*RelationShipNode {
	return m.uses
}

func (m *ModelNode) addRelationShip(from, to Namer, desc string) *RelationShipNode {
	r := Uses(from, to, desc)
	m.uses = append(m.uses, r)
	return r
}

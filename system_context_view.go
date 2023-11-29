package gostructurizr

type SystemContextViewNode struct {
	softwareSystem       *SoftwareSystemNode
	key, description     *string
	addAllSoftwareSystem bool
	addAllPeople         bool
	autoLayout           bool
}

func systemContextView(softwareSystem *SoftwareSystemNode) *SystemContextViewNode {
	return &SystemContextViewNode{
		softwareSystem: softwareSystem,
	}
}

func (s *SystemContextViewNode) WithKey(key string) *SystemContextViewNode {
	s.key = &key
	return s
}

func (s *SystemContextViewNode) Key() *string {
	return s.key
}

func (s *SystemContextViewNode) WithDescription(desc string) *SystemContextViewNode {
	s.description = &desc
	return s
}

func (s *SystemContextViewNode) Description() *string {
	return s.description
}

func (s *SystemContextViewNode) SoftwareSystem() *SoftwareSystemNode {
	return s.softwareSystem
}

func (s *SystemContextViewNode) AddAllSoftwareSystem() *SystemContextViewNode {
	s.addAllSoftwareSystem = true
	return s
}

func (s *SystemContextViewNode) WithAutoLayout() *SystemContextViewNode {
	s.autoLayout = true

	return s
}

func (s *SystemContextViewNode) AutoLayout() bool {
	return s.autoLayout
}

func (s *SystemContextViewNode) IsAllSoftwareSystem() bool {
	return s.addAllSoftwareSystem
}

func (s *SystemContextViewNode) AddAllPeople() *SystemContextViewNode {
	s.addAllPeople = true
	return s
}

func (s *SystemContextViewNode) IsAllPeople() bool {
	return s.addAllPeople
}

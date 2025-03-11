package gostructurizr

type ContainersViewNode struct {
	softwareSystem   *SoftwareSystemNode
	key, description *string
	addAllElement    bool
	addAllPeople     bool
	autoLayout       bool
	includes         []*ExpressionViewNode
	softwareSystems  []*SoftwareSystemNode
}

func containersView(softwareSystem *SoftwareSystemNode) *ContainersViewNode {
	return &ContainersViewNode{
		softwareSystem: softwareSystem,
	}
}

func (s *ContainersViewNode) WithKey(key string) *ContainersViewNode {
	s.key = &key
	return s
}

func (s *ContainersViewNode) Key() *string {
	return s.key
}

func (s *ContainersViewNode) WithDescription(desc string) *ContainersViewNode {
	s.description = &desc
	return s
}

func (s *ContainersViewNode) Description() *string {
	return s.description
}

func (s *ContainersViewNode) SoftwareSystem() *SoftwareSystemNode {
	return s.softwareSystem
}

func (s *ContainersViewNode) AddAllElements() *ContainersViewNode {
	s.addAllElement = true
	return s
}

func (s *ContainersViewNode) WithAutoLayout() *ContainersViewNode {
	s.autoLayout = true

	return s
}

func (s *ContainersViewNode) AutoLayout() bool {
	return s.autoLayout
}

func (s *ContainersViewNode) IsAllElements() bool {
	return s.addAllElement
}

func (s *ContainersViewNode) AddAllPeople() *ContainersViewNode {
	s.addAllPeople = true
	return s
}

func (s *ContainersViewNode) IsAllPeople() bool {
	return s.addAllPeople
}

func (s *ContainersViewNode) WithInclude(e *ExpressionViewNode) {
	s.includes = append(s.includes, e)
}

func (s *ContainersViewNode) Includes() []*ExpressionViewNode {
	return s.includes
}

// AddAllContainers adds all containers to the view
func (s *ContainersViewNode) AddAllContainers() *ContainersViewNode {
	s.addAllElement = true
	return s
}

// AddSoftwareSystem adds a software system to the view
func (s *ContainersViewNode) AddSoftwareSystem(system *SoftwareSystemNode) *ContainersViewNode {
	s.softwareSystems = append(s.softwareSystems, system)
	return s
}

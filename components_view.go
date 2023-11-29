package gostructurizr

type ComponentsViewNode struct {
	container        *ContainerNode
	key, description *string
	addAllElement    bool
	addAllPeople     bool
	autoLayout       bool
}

func componentsView(container *ContainerNode) *ComponentsViewNode {
	return &ComponentsViewNode{
		container: container,
	}
}

func (s *ComponentsViewNode) WithKey(key string) *ComponentsViewNode {
	s.key = &key
	return s
}

func (s *ComponentsViewNode) Key() *string {
	return s.key
}

func (s *ComponentsViewNode) WithDescription(desc string) *ComponentsViewNode {
	s.description = &desc
	return s
}

func (s *ComponentsViewNode) Description() *string {
	return s.description
}

func (s *ComponentsViewNode) Container() *ContainerNode {
	return s.container
}

func (s *ComponentsViewNode) AddAllElements() *ComponentsViewNode {
	s.addAllElement = true
	return s
}

func (s *ComponentsViewNode) WithAutoLayout() *ComponentsViewNode {
	s.autoLayout = true

	return s
}

func (s *ComponentsViewNode) AutoLayout() bool {
	return s.autoLayout
}

func (s *ComponentsViewNode) IsAllElements() bool {
	return s.addAllElement
}

func (s *ComponentsViewNode) AddAllPeople() *ComponentsViewNode {
	s.addAllPeople = true
	return s
}

func (s *ComponentsViewNode) IsAllPeople() bool {
	return s.addAllPeople
}

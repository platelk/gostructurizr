package gostructurizr

type ViewsNode struct {
	configuration      *ViewConfiguration
	systemContextViews []*SystemContextViewNode
	containersView     []*ContainersViewNode
	dynamicView        []*DynamicViewNode
	componentViews     []*ComponentsViewNode
}

func views() *ViewsNode {
	return &ViewsNode{
		configuration: NewViewConfiguration(),
	}
}

func (v *ViewsNode) CreateSystemContextView(node *SoftwareSystemNode) *SystemContextViewNode {
	view := systemContextView(node)
	v.systemContextViews = append(v.systemContextViews, view)

	return view
}

func (v *ViewsNode) CreateContainerView(node *SoftwareSystemNode) *ContainersViewNode {
	c := containersView(node)
	v.containersView = append(v.containersView, c)
	return c
}

func (v *ViewsNode) CreateComponentView(node *ContainerNode) *ComponentsViewNode {
	c := componentsView(node)
	v.componentViews = append(v.componentViews, c)
	return c
}

func (v *ViewsNode) CreateDynamicView(identifier Namer) *DynamicViewNode {
	d := dynamicView()
	v.dynamicView = append(v.dynamicView, d)
	return d
}

func (v *ViewsNode) SystemContextViews() []*SystemContextViewNode {
	return v.systemContextViews
}

func (v *ViewsNode) ContainerViews() []*ContainersViewNode {
	return v.containersView
}

func (v *ViewsNode) ComponentViews() []*ComponentsViewNode {
	return v.componentViews
}

func (v *ViewsNode) Configuration() *ViewConfiguration {
	return v.configuration
}

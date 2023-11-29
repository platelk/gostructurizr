package gostructurizr

type ViewsNode struct {
	configuration      *ViewConfiguration
	systemContextViews []*SystemContextViewNode
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

func (v *ViewsNode) SystemContextViews() []*SystemContextViewNode {
	return v.systemContextViews
}

func (v *ViewsNode) Configuration() *ViewConfiguration {
	return v.configuration
}

package gostructurizr

type WorkspaceNode struct {
	name, description *string
	extends           *string
	model             *ModelNode
	views             *ViewsNode
}

func Workspace() *WorkspaceNode {
	return &WorkspaceNode{
		model: Model(),
		views: views(),
	}
}

func (w *WorkspaceNode) WithName(n string) *WorkspaceNode {
	w.name = &n
	return w
}

func (w *WorkspaceNode) WithDesc(d string) *WorkspaceNode {
	w.description = &d
	return w
}

func (w *WorkspaceNode) WithExtend(source string) *WorkspaceNode {
	w.extends = &source
	return w
}

func (w *WorkspaceNode) Name() *string {
	return w.name
}

func (w *WorkspaceNode) Desc() *string {
	return w.description
}

func (w *WorkspaceNode) Extend() *string {
	return w.extends
}

func (w *WorkspaceNode) Model() *ModelNode {
	return w.model
}

func (w *WorkspaceNode) Views() *ViewsNode {
	return w.views
}

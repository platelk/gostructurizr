package gostructurizr

type ViewsNode struct {
	configuration      *ViewConfiguration
	systemContextViews []*SystemContextViewNode
	containersView     []*ContainersViewNode
	dynamicView        []*DynamicViewNode
	componentViews     []*ComponentsViewNode
	deploymentViews    []*DeploymentViewNode
	filteredViews      []*FilteredViewNode
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

// CreateDeploymentView creates a new deployment view for a software system and environment
func (v *ViewsNode) CreateDeploymentView(softwareSystem *SoftwareSystemNode, environment DeploymentEnvironment) *DeploymentViewNode {
	d := DeploymentView(softwareSystem, environment)
	v.deploymentViews = append(v.deploymentViews, d)
	return d
}

// CreateDevView creates a deployment view for Development environment
func (v *ViewsNode) CreateDevView(softwareSystem *SoftwareSystemNode) *DeploymentViewNode {
	return v.CreateDeploymentView(softwareSystem, DevelopmentEnvironment)
}

// CreateTestView creates a deployment view for Test environment
func (v *ViewsNode) CreateTestView(softwareSystem *SoftwareSystemNode) *DeploymentViewNode {
	return v.CreateDeploymentView(softwareSystem, TestEnvironment)
}

// CreateStagingView creates a deployment view for Staging environment
func (v *ViewsNode) CreateStagingView(softwareSystem *SoftwareSystemNode) *DeploymentViewNode {
	return v.CreateDeploymentView(softwareSystem, StagingEnvironment)
}

// CreateProdView creates a deployment view for Production environment
func (v *ViewsNode) CreateProdView(softwareSystem *SoftwareSystemNode) *DeploymentViewNode {
	return v.CreateDeploymentView(softwareSystem, ProductionEnvironment)
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

func (v *ViewsNode) DeploymentViews() []*DeploymentViewNode {
	return v.deploymentViews
}

// CreateFilteredView creates a filtered view based on an existing view
func (v *ViewsNode) CreateFilteredView(baseView Viewable, title string) *FilteredViewNode {
	filteredView := FilteredView(baseView)
	filteredView.WithTitle(title)
	v.filteredViews = append(v.filteredViews, filteredView)
	return filteredView
}

// FilteredViews returns all filtered views
func (v *ViewsNode) FilteredViews() []*FilteredViewNode {
	return v.filteredViews
}

func (v *ViewsNode) Configuration() *ViewConfiguration {
	return v.configuration
}

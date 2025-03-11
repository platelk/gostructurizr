package dsl

const (
	Workspace          = "workspace"
	Extends            = "extends"
	Model              = "model"
	Space              = " "
	OpenBracket        = "{"
	CloseBracket       = "}"
	NewLine            = "\n"
	DoubleQuotes       = "\""
	Equal              = "="
	Person             = "person"
	Comma              = ","
	TagSeparator       = Comma
	SoftwareSystem     = "softwareSystem"
	Arrow              = "->"
	Views              = "views"
	SystemContext      = "systemContext"
	Container          = "container"
	Component          = "component"
	Tags               = "tags"
	Include            = "include"
	All                = "*"
	AutoLayout         = "autoLayout"
	Element            = "element"
	Shape              = "shape"
	Height             = "height"
	Width              = "width"
	Icon               = "icon"
	Background         = "background"
	Color              = "color"
	Stroke             = "stroke"
	StrokeWidth        = "strokeWidth"
	FontSize           = "fontSize"
	Border             = "border"
	Opacity            = "opacity"
	Metadata           = "metadata"
	Description        = "description"
	Properties         = "properties"
	Styles             = "styles"
	DeploymentNode     = "deploymentNode"
	DeploymentView     = "deploymentView"
	InfrastructureNode = "infrastructureNode"
	ContainerInstance  = "containerInstance"
	InstanceId         = "instanceId"
	HealthCheck        = "healthCheck"
	Environment        = "environment"
	Technology         = "technology"
	Url                = "url"
	Name               = "name"
	Interval           = "interval"
	Timeout            = "timeout"
	Location           = "location"
	Enterprise         = "enterprise"
	Key                = "key"
	Group              = "group"
	Dynamic            = "dynamic"
	
	// Advanced styling
	BorderStyle        = "borderStyle"
	Shadow             = "shadow"
	FontFamily         = "fontFamily"
	FontStyle          = "fontStyle"
	Icons              = "icons"
	ZIndex             = "zIndex"
	Rotation           = "rotation"
	Position           = "position"
	
	// Border styles
	BorderSolid        = "solid"
	BorderDashed       = "dashed"
	BorderDotted       = "dotted"
	
	// Advanced relationship styling
	Style              = "style"
	Thickness          = "thickness"
	FontColor          = "fontColor"
	Routing            = "routing"
	SourceTerminator   = "sourceTerminator"
	DestTerminator     = "destinationTerminator"
	
	// Routing styles
	RoutingDirect      = "direct"
	RoutingCurved      = "curved"
	RoutingOrthogonal  = "orthogonal"
	
	// Filtered view
	FilteredView       = "filteredView"
	BaseView           = "baseView"
	Title              = "title"
	IncludeTag         = "includeTag"
	ExcludeTag         = "excludeTag"
	
	// Documentation
	Documentation      = "documentation"
	Decision           = "decision"
	Format             = "format"
	Content            = "content"
	Status             = "status"
)

const (
	EmptyIdentifier = DoubleQuotes + DoubleQuotes
)

package tags

type Tag string

const (
	SoftwareSystem      Tag = "Software system"
	Person              Tag = "Person"
	Element             Tag = "Element"
	Container           Tag = "Container"
	Component           Tag = "Component"
	RelationShip        Tag = "relationship"
	Synchronous         Tag = "synchronous"
	Asynchronous        Tag = "asynchronous"
	DeploymentNode      Tag = "Deployment Node"
	InfrastructureNode  Tag = "Infrastructure Node"
	ContainerInstance   Tag = "Container Instance"
	Enterprise          Tag = "Enterprise"
	Internal            Tag = "Internal"
	External            Tag = "External"
	Database            Tag = "Database"
	Queue               Tag = "Queue"
	Group               Tag = "Group"
	Dynamic             Tag = "Dynamic"
	HealthCheck         Tag = "Health Check"
)

func (t Tag) String() string {
	return string(t)
}

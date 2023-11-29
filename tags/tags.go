package tags

type Tag string

const (
	SoftwareSystem Tag = "Software system"
	Person         Tag = "Person"
	Element        Tag = "Element"
	Container      Tag = "Container"
	RelationShip   Tag = "relationship"
	Synchronous    Tag = "synchronous"
	Asynchronous   Tag = "asynchronous"
)

func (t Tag) String() string {
	return string(t)
}

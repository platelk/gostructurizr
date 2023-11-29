package tags

type Tag string

const (
	SoftwareSystem Tag = "Software system"
	Person         Tag = "Person"
)

func (t Tag) String() string {
	return string(t)
}

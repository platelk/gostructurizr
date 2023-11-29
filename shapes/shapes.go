package shapes

type Shape string

const (
	Person Shape = "person"
)

func (s Shape) String() string {
	return string(s)
}

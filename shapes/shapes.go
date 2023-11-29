package shapes

type Shape string

const (
	Person   Shape = "person"
	Pipe     Shape = "pipe"
	Hexagon  Shape = "hexagon"
	Cylinder Shape = "cylinder"
)

func (s Shape) String() string {
	return string(s)
}

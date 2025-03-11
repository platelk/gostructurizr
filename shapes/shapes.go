package shapes

type Shape string

const (
	// Basic shapes
	Person   Shape = "person"
	Pipe     Shape = "pipe"
	Hexagon  Shape = "hexagon"
	Cylinder Shape = "cylinder"
	
	// Device shapes
	WebBrowser          Shape = "WebBrowser"
	MobileDevicePortrait Shape = "MobileDevicePortrait"
	MobileDeviceLandscape Shape = "MobileDeviceLandscape"
	
	// Component shapes
	Component           Shape = "Component"
	
	// Container shapes
	Database            Shape = "Cylinder"
	Queue               Shape = "Pipe"
	
	// Additional shapes
	Box                 Shape = "Box"
	RoundedBox          Shape = "RoundedBox"
	Circle              Shape = "Circle"
	Ellipse             Shape = "Ellipse"
	Folder              Shape = "Folder"
	Robot               Shape = "Robot"
)

func (s Shape) String() string {
	return string(s)
}

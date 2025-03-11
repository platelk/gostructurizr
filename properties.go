package gostructurizr

// Properties represents a map of key-value properties for elements
type Properties struct {
	Properties map[string]string
}

// NewProperties creates a new Properties
func NewProperties() Properties {
	return Properties{
		Properties: make(map[string]string),
	}
}

// Add adds a property with key and value
func (p *Properties) Add(key, value string) *Properties {
	if p.Properties == nil {
		p.Properties = make(map[string]string)
	}
	p.Properties[key] = value
	return p
}

// Get retrieves a property value by key
func (p *Properties) Get(key string) string {
	if p.Properties == nil {
		return ""
	}
	return p.Properties[key]
}

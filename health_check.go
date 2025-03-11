package gostructurizr

// HealthCheckNode represents a health check for a container instance
type HealthCheckNode struct {
	name       string
	url        string
	interval   int  // interval in seconds
	timeout    int  // timeout in milliseconds
	parent     *ContainerInstanceNode
	properties Properties
}

// HealthCheck creates a new HealthCheckNode
func HealthCheck(name, url string) *HealthCheckNode {
	return &HealthCheckNode{
		name:       name,
		url:        url,
		interval:   60,    // default 60 seconds
		timeout:    1000,  // default 1000 milliseconds
		properties: Properties{Properties: make(map[string]string)},
	}
}

// Name returns the name of the health check
func (h *HealthCheckNode) Name() string {
	return h.name
}

// Url returns the URL of the health check
func (h *HealthCheckNode) Url() string {
	return h.url
}

// WithUrl sets the URL of the health check
func (h *HealthCheckNode) WithUrl(url string) *HealthCheckNode {
	h.url = url
	return h
}

// WithInterval sets the interval in seconds for the health check
func (h *HealthCheckNode) WithInterval(seconds int) *HealthCheckNode {
	h.interval = seconds
	return h
}

// Interval returns the interval in seconds
func (h *HealthCheckNode) Interval() int {
	return h.interval
}

// WithTimeout sets the timeout in milliseconds for the health check
func (h *HealthCheckNode) WithTimeout(milliseconds int) *HealthCheckNode {
	h.timeout = milliseconds
	return h
}

// Timeout returns the timeout in milliseconds
func (h *HealthCheckNode) Timeout() int {
	return h.timeout
}

// Properties returns the properties of the health check
func (h *HealthCheckNode) Properties() *Properties {
	return &h.properties
}
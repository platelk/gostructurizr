package renderer

import (
	"fmt"
	"io"

	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
)

// HealthCheckRenderer renders a health check to DSL
type HealthCheckRenderer struct {
	BaseRenderer
}

// NewHealthCheckRenderer creates a new health check renderer
func NewHealthCheckRenderer(w io.Writer, level int) *HealthCheckRenderer {
	return &HealthCheckRenderer{
		BaseRenderer: BaseRenderer{
			w:     w,
			level: level,
		},
	}
}

// Render renders a health check to DSL
func (r *HealthCheckRenderer) Render(healthCheck *gostructurizr.HealthCheckNode) error {
	r.WriteLine(dsl.HealthCheck + dsl.Space + dsl.OpenBracket)
	r.level++

	// Name
	r.WriteLine(fmt.Sprintf("%s %q", dsl.Name, healthCheck.Name()))

	// URL
	r.WriteLine(fmt.Sprintf("%s %q", dsl.Url, healthCheck.Url()))

	// Interval if not default
	if healthCheck.Interval() != 60 {
		r.WriteLine(fmt.Sprintf("%s %d", dsl.Interval, healthCheck.Interval()))
	}

	// Timeout if not default
	if healthCheck.Timeout() != 1000 {
		r.WriteLine(fmt.Sprintf("%s %d", dsl.Timeout, healthCheck.Timeout()))
	}

	// Properties
	renderProperties(r.w, healthCheck.Properties(), r.level)

	// End health check
	r.level--
	r.WriteLine(dsl.CloseBracket)

	return nil
}
package renderer

import (
	"fmt"
	"io"

	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
)

// ContainerInstanceRenderer renders a container instance to DSL
type ContainerInstanceRenderer struct {
	BaseRenderer
}

// NewContainerInstanceRenderer creates a new container instance renderer
func NewContainerInstanceRenderer(w io.Writer, level int) *ContainerInstanceRenderer {
	return &ContainerInstanceRenderer{
		BaseRenderer: BaseRenderer{
			w:     w,
			level: level,
		},
	}
}

// Render renders a container instance to DSL
func (r *ContainerInstanceRenderer) Render(instance *gostructurizr.ContainerInstanceNode) error {
	r.WriteLine(dsl.ContainerInstance + dsl.Space + dsl.OpenBracket)
	r.level++

	// Reference to container
	containerRef := fmt.Sprintf("%s.%s", 
		instance.Container().Parent().Name(), 
		instance.Container().Name())
	r.WriteLine(fmt.Sprintf("%s %q", dsl.Container, containerRef))

	// Instance ID if not default
	if instance.InstanceId() != 1 {
		r.WriteLine(fmt.Sprintf("%s %d", dsl.InstanceId, instance.InstanceId()))
	}

	// Properties
	renderProperties(r.w, instance.Properties(), r.level)

	// Tags
	renderTags(r.w, instance.Tags(), r.level)

	// Health checks
	for _, healthCheck := range instance.HealthChecks() {
		healthCheckRenderer := NewHealthCheckRenderer(r.w, r.level)
		if err := healthCheckRenderer.Render(healthCheck); err != nil {
			return err
		}
	}

	// End container instance
	r.level--
	r.WriteLine(dsl.CloseBracket)

	return nil
}
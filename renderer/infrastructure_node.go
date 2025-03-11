package renderer

import (
	"fmt"
	"io"

	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
)

// InfrastructureNodeRenderer renders an infrastructure node to DSL
type InfrastructureNodeRenderer struct {
	BaseRenderer
}

// NewInfrastructureNodeRenderer creates a new infrastructure node renderer
func NewInfrastructureNodeRenderer(w io.Writer, level int) *InfrastructureNodeRenderer {
	return &InfrastructureNodeRenderer{
		BaseRenderer: BaseRenderer{
			w:     w,
			level: level,
		},
	}
}

// Render renders an infrastructure node to DSL
func (r *InfrastructureNodeRenderer) Render(node *gostructurizr.InfrastructureNodeNode) error {
	r.WriteLine(dsl.InfrastructureNode + dsl.Space + dsl.OpenBracket)
	r.level++

	// Name
	r.WriteLine(fmt.Sprintf("%s %q", dsl.Name, node.Name()))

	// Description if available
	if node.Description() != "" {
		r.WriteLine(fmt.Sprintf("%s %q", dsl.Description, node.Description()))
	}

	// Technology
	if node.Technology() != "" {
		r.WriteLine(fmt.Sprintf("%s %q", dsl.Technology, node.Technology()))
	}

	// Properties
	renderProperties(r.w, node.Properties(), r.level)

	// Tags
	renderTags(r.w, node.Tags(), r.level)

	// End infrastructure node
	r.level--
	r.WriteLine(dsl.CloseBracket)

	return nil
}
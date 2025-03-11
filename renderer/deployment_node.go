package renderer

import (
	"fmt"
	"io"

	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
)

// DeploymentNodeRenderer renders a deployment node to DSL
type DeploymentNodeRenderer struct {
	BaseRenderer
}

// NewDeploymentNodeRenderer creates a new deployment node renderer
func NewDeploymentNodeRenderer(w io.Writer, level int) *DeploymentNodeRenderer {
	return &DeploymentNodeRenderer{
		BaseRenderer: BaseRenderer{
			w:     w,
			level: level,
		},
	}
}

// Render renders a deployment node to DSL
func (r *DeploymentNodeRenderer) Render(node *gostructurizr.DeploymentNodeNode) error {
	r.WriteLine(dsl.DeploymentNode + dsl.Space + dsl.OpenBracket)
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

	// Environment (implicitly set from parent usually)
	r.WriteLine(fmt.Sprintf("%s %q", dsl.Environment, string(node.Environment())))

	// Properties
	renderProperties(r.w, node.Properties(), r.level)

	// Tags
	renderTags(r.w, node.Tags(), r.level)

	// Render child deployment nodes
	for _, child := range node.Children() {
		childRenderer := NewDeploymentNodeRenderer(r.w, r.level)
		if err := childRenderer.Render(child); err != nil {
			return err
		}
	}

	// Render infrastructure nodes
	for _, infra := range node.InfrastructureNodes() {
		infraRenderer := NewInfrastructureNodeRenderer(r.w, r.level)
		if err := infraRenderer.Render(infra); err != nil {
			return err
		}
	}

	// Render container instances
	for _, instance := range node.ContainerInstances() {
		instanceRenderer := NewContainerInstanceRenderer(r.w, r.level)
		if err := instanceRenderer.Render(instance); err != nil {
			return err
		}
	}

	// End deployment node
	r.level--
	r.WriteLine(dsl.CloseBracket)

	return nil
}
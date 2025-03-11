package renderer

import (
	"fmt"
	"io"

	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
)

// DeploymentViewRenderer renders a deployment view to DSL
type DeploymentViewRenderer struct {
	BaseRenderer
}

// NewDeploymentViewRenderer creates a new deployment view renderer
func NewDeploymentViewRenderer(w io.Writer, level int) *DeploymentViewRenderer {
	return &DeploymentViewRenderer{
		BaseRenderer: BaseRenderer{
			w:     w,
			level: level,
		},
	}
}

// Render renders a deployment view to DSL
func (r *DeploymentViewRenderer) Render(view *gostructurizr.DeploymentViewNode) error {
	r.WriteLine(dsl.DeploymentView + dsl.Space + dsl.OpenBracket)
	r.level++

	// Generate software system
	if view.SoftwareSystem() != nil {
		r.WriteLine(fmt.Sprintf("%s %s", dsl.SoftwareSystem, renderNamer(view.SoftwareSystem())))
	}

	// Generate environment
	environment := string(view.Environment())
	if environment != "" {
		r.WriteLine(fmt.Sprintf("%s %q", dsl.Environment, environment))
	}

	// Generate key
	if view.GetKey() != "" {
		r.WriteLine(fmt.Sprintf("%s %q", dsl.Key, view.GetKey()))
	}

	// Generate description
	if view.GetDescription() != "" {
		r.WriteLine(fmt.Sprintf("%s %q", dsl.Description, view.GetDescription()))
	}

	// Generate auto layout
	if view.IsAutoLayout() {
		r.WriteLine(dsl.AutoLayout)
	}

	// Include specific elements
	for _, element := range view.Elements() {
		r.WriteLine(fmt.Sprintf("%s %s", dsl.Include, renderNamer(element)))
	}

	// Include specific relationships
	for _, rs := range view.RelationShips() {
		from := rs.From().Name()
		to := rs.To().Name()
		r.WriteLine(fmt.Sprintf("%s %s %s %s", dsl.Include, from, dsl.Arrow, to))
	}

	// End deployment view
	r.level--
	r.WriteLine(dsl.CloseBracket)

	return nil
}
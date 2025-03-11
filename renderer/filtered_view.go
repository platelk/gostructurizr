package renderer

import (
	"fmt"
	"io"

	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
)

// FilteredViewRenderer renders a filtered view to DSL
type FilteredViewRenderer struct {
	BaseRenderer
}

// NewFilteredViewRenderer creates a new filtered view renderer
func NewFilteredViewRenderer(w io.Writer, level int) *FilteredViewRenderer {
	return &FilteredViewRenderer{
		BaseRenderer: BaseRenderer{
			w:     w,
			level: level,
		},
	}
}

// Render renders a filtered view to DSL
func (r *FilteredViewRenderer) Render(view *gostructurizr.FilteredViewNode) error {
	r.WriteLine(dsl.FilteredView + dsl.Space + dsl.OpenBracket)
	r.level++

	// Render base view if available
	baseView := view.BaseView()
	if baseView != nil && baseView.Key() != nil {
		// Use the key from the base view
		r.WriteLine(fmt.Sprintf("%s %q", dsl.BaseView, *baseView.Key()))
	}

	// Render title
	if view.Title() != "" {
		r.WriteLine(fmt.Sprintf("%s %q", dsl.Title, view.Title()))
	}

	// Render key
	if view.Key() != "" {
		r.WriteLine(fmt.Sprintf("%s %q", dsl.Key, view.Key()))
	}

	// Render description
	if view.Description() != "" {
		r.WriteLine(fmt.Sprintf("%s %q", dsl.Description, view.Description()))
	}

	// Render filter criteria
	for _, filter := range view.FilterCriteria() {
		mode := string(filter.Mode)
		filterType := string(filter.Type)
		value := filter.Value

		r.WriteLine(fmt.Sprintf("%s %s %q", mode, filterType, value))
	}

	// Render auto layout
	if view.IsAutoLayout() {
		r.WriteLine(dsl.AutoLayout)
	}

	r.level--
	r.WriteLine(dsl.CloseBracket)

	return nil
}
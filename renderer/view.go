package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderView(v *gostructurizr.ViewsNode, renderer *strings.Builder, level int) error {
	writeLine(renderer, level, dsl.Views, dsl.Space, dsl.OpenBracket)
	for _, s := range v.SystemContextViews() {
		if err := renderSystemContext(s, renderer, level+1); err != nil {
			return fmt.Errorf("can't generate system context view: %w", err)
		}
	}
	for _, c := range v.ContainerViews() {
		if err := renderViewContainer(c, renderer, level+1); err != nil {
			return fmt.Errorf("can't generate container view: %w", err)
		}
	}
	for _, c := range v.ComponentViews() {
		if err := renderViewComponent(c, renderer, level+1); err != nil {
			return fmt.Errorf("can't generate component view: %w", err)
		}
	}
	for _, d := range v.DeploymentViews() {
		if err := renderDeploymentView(d, renderer, level+1); err != nil {
			return fmt.Errorf("can't generate deployment view: %w", err)
		}
	}
	for _, f := range v.FilteredViews() {
		if err := renderFilteredView(f, renderer, level+1); err != nil {
			return fmt.Errorf("can't generate filtered view: %w", err)
		}
	}
	if err := renderViewConfiguration(v.Configuration(), renderer, level+1); err != nil {
		return fmt.Errorf("can't render view configuration: %w", err)
	}
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

func renderFilteredView(f *gostructurizr.FilteredViewNode, renderer *strings.Builder, level int) error {
	writeLine(renderer, level, dsl.FilteredView, dsl.Space, dsl.OpenBracket)
	
	// Render base view if available
	baseView := f.BaseView()
	if baseView != nil && baseView.Key() != nil {
		// Use the key from the base view
		writeLine(renderer, level+1, dsl.BaseView, dsl.Space, generateStringIdentifier(*baseView.Key()))
	}
	
	// Render title
	if f.Title() != "" {
		writeLine(renderer, level+1, dsl.Title, dsl.Space, generateStringIdentifier(f.Title()))
	}
	
	// Render key
	if f.Key() != "" {
		writeLine(renderer, level+1, dsl.Key, dsl.Space, generateStringIdentifier(f.Key()))
	}
	
	// Render description
	if f.Description() != "" {
		writeLine(renderer, level+1, dsl.Description, dsl.Space, generateStringIdentifier(f.Description()))
	}
	
	// Render filter criteria
	for _, filter := range f.FilterCriteria() {
		mode := string(filter.Mode)
		filterType := string(filter.Type)
		value := filter.Value
		
		writeLine(renderer, level+1, mode, dsl.Space, filterType, dsl.Space, generateStringIdentifier(value))
	}
	
	// Render auto layout
	if f.IsAutoLayout() {
		writeLine(renderer, level+1, dsl.AutoLayout)
	}
	
	writeLine(renderer, level, dsl.CloseBracket)
	
	return nil
}

func renderDeploymentView(d *gostructurizr.DeploymentViewNode, renderer *strings.Builder, level int) error {
	writeLine(renderer, level, dsl.DeploymentView, dsl.Space, dsl.OpenBracket)
	
	// Software system
	if d.SoftwareSystem() != nil {
		writeLine(renderer, level+1, dsl.SoftwareSystem, dsl.Space, generateVarName(d.SoftwareSystem().Name()))
	}
	
	// Environment
	environment := string(d.Environment())
	if environment != "" {
		writeLine(renderer, level+1, dsl.Environment, dsl.Space, generateStringIdentifier(environment))
	}
	
	// Key
	if d.GetKey() != "" {
		writeLine(renderer, level+1, dsl.Key, dsl.Space, generateStringIdentifier(d.GetKey()))
	}
	
	// Description
	if d.GetDescription() != "" {
		writeLine(renderer, level+1, dsl.Description, dsl.Space, generateStringIdentifier(d.GetDescription()))
	}
	
	// Auto layout
	if d.IsAutoLayout() {
		writeLine(renderer, level+1, dsl.AutoLayout)
	}
	
	// Elements
	for _, element := range d.Elements() {
		writeLine(renderer, level+1, dsl.Include, dsl.Space, generateVarName(element.Name()))
	}
	
	// Relationships
	for _, rs := range d.RelationShips() {
		from := generateVarName(rs.From().Name())
		to := generateVarName(rs.To().Name())
		writeLine(renderer, level+1, dsl.Include, dsl.Space, from, dsl.Space, dsl.Arrow, dsl.Space, to)
	}
	
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}

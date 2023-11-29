package main

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/renderer"
	"github.com/platelk/gostructurizr/shapes"
	"github.com/platelk/gostructurizr/tags"
	"strings"
)

func main() {
	// all software architecture models belong to a workspace
	workspace := gostructurizr.Workspace().WithName("Getting Started").WithDesc("This is a model of my software system.")
	model := workspace.Model()

	// create a model to describe a user using a software system
	user := model.AddPerson("User", "A user of my software system.")
	softwareSystem := model.AddSoftwareSystem("Software System", "My software system.")
	user.Uses(softwareSystem, "Uses")

	// create a system context diagram showing people and software systems
	views := workspace.Views()
	contextView := views.CreateSystemContextView(softwareSystem).WithKey("SystemContext").WithDescription("An example of a System Context diagram.")
	contextView.AddAllSoftwareSystem()
	contextView.AddAllPeople()
	contextView.WithAutoLayout()

	// add some styling to the diagram elements
	styles := views.Configuration().Styles()
	styles.AddElementStyle(tags.SoftwareSystem).WithBackground("#1168bd").WithColor("#ffffff")
	styles.AddElementStyle(tags.Person).WithBackground("#08427b").WithColor("#ffffff").WithShape(shapes.Person)

	var b strings.Builder
	r := renderer.NewDSLRenderer(&b)
	if err := r.Render(workspace); err != nil {
		panic(fmt.Sprintf("renderer didn't succeed: %s", err))
	}
	fmt.Println(b.String())
}

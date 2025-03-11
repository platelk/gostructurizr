package main

import (
	"fmt"
	"os"

	strukt "github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/renderer"
	"github.com/platelk/gostructurizr/shapes"
	"github.com/platelk/gostructurizr/tags"
)

func main() {
	// Create a new workspace
	workspace := strukt.Workspace().
		WithName("Advanced Styling Example").
		WithDesc("This is an example of advanced element and relationship styling in Structurizr")

	// Create the model
	model := workspace.Model()

	// Create users and systems
	user := model.AddPerson("User", "A user of the system")
	user.Tags().Add(string(tags.External))

	webApp := model.AddSoftwareSystem("Web Application", "The main web application")
	webApp.Tags().Add("WebApp")

	database := model.AddSoftwareSystem("Database", "The primary database") 
	database.Tags().Add(string(tags.Database))

	cache := model.AddSoftwareSystem("Cache", "Redis cache")
	cache.Tags().Add("Cache")

	messaging := model.AddSoftwareSystem("Messaging", "Kafka messaging platform")
	messaging.Tags().Add(string(tags.Queue))

	api := model.AddSoftwareSystem("API", "External REST API")
	api.Tags().Add("API")

	// Add relationships
	user.Uses(webApp, "Uses")
	
	// There's no direct API to add tags to relationships in the current implementation
	// This would typically be done with Tagged relationships in Structurizr
	webApp.Uses(database, "Reads from and writes to").WithInteractionStyle(strukt.Synchronous)
	webApp.Uses(cache, "Reads from and writes to")
	webApp.Uses(messaging, "Publishes events to").WithInteractionStyle(strukt.Asynchronous)
	webApp.Uses(api, "Makes API calls to").WithInteractionStyle(strukt.Synchronous)

	// Create views
	views := workspace.Views()

	// Create a System Context view
	contextView := views.CreateSystemContextView(webApp).
		WithKey("SystemContext").
		WithDescription("System Context diagram")
	contextView.WithAutoLayout()
	contextView.AddAllElements()
	contextView.AddAllPeople()

	// Set up the styling
	styles := views.Configuration().Styles()

	// Element styles
	styles.AddElementStyle(tags.Person).
		WithShape(shapes.Person).
		WithBackground("#08427B").
		WithColor("#ffffff").
		WithFontFamily("Arial").
		WithFontSize(24).
		WithBorder(20).
		WithShadow(true)

	styles.AddElementStyle(tags.SoftwareSystem).
		WithShape(shapes.RoundedBox).
		WithBackground("#1168BD").
		WithColor("#ffffff")

	styles.AddElementStyle(tags.Database).
		WithShape(shapes.Cylinder).
		WithBackground("#1168BD").
		WithColor("#ffffff")

	styles.AddElementStyle("WebApp").
		WithBackground("#62A420").
		WithStrokeWidth(2).
		WithBorderStyle(strukt.Dashed)

	styles.AddElementStyle("API").
		WithBackground("#85BBF0").
		WithBorder(4).
		WithShape(shapes.Hexagon)

	styles.AddElementStyle("Cache").
		WithBackground("#D4A017").
		WithRotation(15)

	styles.AddElementStyle(tags.Queue).
		WithBackground("#E62D2D").
		WithShape(shapes.Pipe)

	styles.AddElementStyle(tags.External).
		WithBackground("#999999").
		WithFontStyle("italic")

	// Relationship styles
	styles.AddElementStyle(tags.RelationShip).
		WithBackground("#707070") // Use element style as workaround

	// Sync style
	syncTag := tags.Tag("Sync")
	styles.AddAdvancedRelationshipStyle(syncTag).
		WithColor("#289CE1").
		WithFontColor("#289CE1").
		WithFontSize(12).
		WithWidth(2).
		WithLineStyle(strukt.SolidLine).
		WithDirectRouting()

	// Async style
	asyncTag := tags.Tag("Async")
	styles.AddAdvancedRelationshipStyle(asyncTag).
		WithColor("#E62D2D").
		WithFontColor("#E62D2D").
		WithFontSize(12).
		WithWidth(2).
		WithDashed().
		WithCurvedRouting()

	// Cache style
	cacheTag := tags.Tag("Cache")
	styles.AddAdvancedRelationshipStyle(cacheTag).
		WithColor("#D4A017").
		WithFontColor("#D4A017").
		WithWidth(2).
		WithDotted().
		WithEndTerminator(strukt.Arrow)

	// Database style
	dbTag := tags.Tag("Database")
	styles.AddAdvancedRelationshipStyle(dbTag).
		WithColor("#1168BD").
		WithWidth(2).
		WithOrthogonalRouting()

	// Write the DSL to stdout
	renderer := renderer.NewDSLRenderer(os.Stdout)
	err := renderer.Render(workspace)
	if err != nil {
		fmt.Printf("Error rendering: %v\n", err)
		os.Exit(1)
	}
}
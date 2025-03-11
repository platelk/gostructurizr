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
		WithName("Filtered Views Example").
		WithDesc("This is an example of filtered views in Structurizr")

	// Create the model
	model := workspace.Model()

	// Create enterprise boundary
	model.SetEnterprise("Example Corp")

	// Create users and systems
	customerA := model.AddPerson("Customer A", "A premium customer")
	customerA.Tags().Add("Customer")
	customerA.Tags().Add("Premium")
	
	customerB := model.AddPerson("Customer B", "A regular customer")
	customerB.Tags().Add("Customer")
	customerB.Tags().Add("Regular")
	
	administrator := model.AddPerson("Administrator", "System administrator")
	administrator.Tags().Add("Staff")
	administrator.Tags().Add("Admin")
	
	supportStaff := model.AddPerson("Support Staff", "Customer support")
	supportStaff.Tags().Add("Staff")
	supportStaff.Tags().Add("Support")

	// Create internal systems
	webApp := model.AddSoftwareSystem("Web Application", "The main web application")
	webApp.Tags().Add("Internal")
	webApp.Tags().Add("WebApp")
	
	customerDB := model.AddSoftwareSystem("Customer Database", "Stores customer information")
	customerDB.Tags().Add("Internal")
	customerDB.Tags().Add("Database")
	customerDB.Tags().Add("Critical")
	
	reportingSystem := model.AddSoftwareSystem("Reporting System", "Generates business reports")
	reportingSystem.Tags().Add("Internal")
	reportingSystem.Tags().Add("Reporting")
	
	adminPortal := model.AddSoftwareSystem("Admin Portal", "Admin management interface")
	adminPortal.Tags().Add("Internal")
	adminPortal.Tags().Add("AdminTool")

	// Create external systems
	paymentProvider := model.AddSoftwareSystem("Payment Provider", "Processes payments")
	paymentProvider.Tags().Add("External")
	paymentProvider.Tags().Add("Payment")
	
	emailSystem := model.AddSoftwareSystem("Email System", "Sends emails to customers")
	emailSystem.Tags().Add("External")
	emailSystem.Tags().Add("Communication")
	
	monitoringSystem := model.AddSoftwareSystem("Monitoring System", "Monitors system health")
	monitoringSystem.Tags().Add("External")
	monitoringSystem.Tags().Add("Monitoring")

	// Add relationships
	customerA.Uses(webApp, "Uses")
	customerB.Uses(webApp, "Uses")
	administrator.Uses(adminPortal, "Manages system using")
	supportStaff.Uses(adminPortal, "Views customer info using")

	webApp.Uses(customerDB, "Reads from and writes to")
	webApp.Uses(paymentProvider, "Makes payments using")
	webApp.Uses(emailSystem, "Sends emails using")

	adminPortal.Uses(customerDB, "Reads from")
	adminPortal.Uses(reportingSystem, "Generates reports using")

	reportingSystem.Uses(customerDB, "Reads from")
	monitoringSystem.Uses(webApp, "Monitors")
	monitoringSystem.Uses(customerDB, "Monitors")

	// Create views
	views := workspace.Views()

	// Create a System Context view - this will be our base view
	contextView := views.CreateSystemContextView(webApp).
		WithKey("SystemContext").
		WithDescription("The system context diagram")
	contextView.WithAutoLayout()
	contextView.AddAllElements()
	contextView.AddAllPeople()

	// Create filtered views based on the context view
	// Customer view - showing only customer-related elements
	customerView := views.CreateFilteredView(contextView, "Shows customer interaction").
		WithKey("CustomerView")
	customerView.WithAutoLayout()
	customerView.AddIncludeTag("Customer")
	customerView.AddIncludeTag("WebApp")
	customerView.AddIncludeTag("Payment")
	customerView.AddExcludeTag("Admin")
	customerView.AddExcludeTag("Monitoring")
	customerView.AddExcludeTag("Reporting")

	// Admin view - showing admin-focused elements
	adminView := views.CreateFilteredView(contextView, "Shows admin capabilities").
		WithKey("AdminView")
	adminView.WithAutoLayout()
	adminView.AddIncludeTag("Admin")
	adminView.AddIncludeTag("Staff")
	adminView.AddIncludeTag("AdminTool")
	adminView.AddIncludeTag("Reporting")
	adminView.AddExcludeTag("Customer")
	adminView.AddExcludeTag("Payment")

	// Critical systems view
	criticalView := views.CreateFilteredView(contextView, "Shows critical systems only").
		WithKey("CriticalView")
	criticalView.WithAutoLayout()
	criticalView.AddIncludeTag("Critical")
	criticalView.AddIncludeTag("WebApp")

	// External systems view
	externalView := views.CreateFilteredView(contextView, "Shows only external integrations").
		WithKey("ExternalView")
	externalView.WithAutoLayout()
	externalView.AddIncludeTag("External")
	externalView.AddIncludeTag("WebApp")

	// Set up the styling
	styles := views.Configuration().Styles()

	// Element styles
	styles.AddElementStyle(tags.Person).
		WithShape(shapes.Person).
		WithBackground("#08427B").
		WithColor("#ffffff")

	styles.AddElementStyle(tags.SoftwareSystem).
		WithShape(shapes.RoundedBox).
		WithBackground("#1168BD").
		WithColor("#ffffff")

	styles.AddElementStyle("Customer").
		WithBackground("#3498DB")

	styles.AddElementStyle("Premium").
		WithBackground("#2E86C1").
		WithFontStyle("bold")

	styles.AddElementStyle("Regular").
		WithBackground("#5DADE2")

	styles.AddElementStyle("Staff").
		WithBackground("#16A085")

	styles.AddElementStyle("Database").
		WithShape(shapes.Cylinder).
		WithBackground("#9B59B6")

	styles.AddElementStyle("AdminTool").
		WithBackground("#2C3E50")

	styles.AddElementStyle("Critical").
		WithBorder(4).
		WithBorderStyle(strukt.Solid).
		WithBackground("#E74C3C")

	styles.AddElementStyle("External").
		WithBackground("#95A5A6").
		WithFontStyle("italic")

	// Write the DSL to stdout
	renderer := renderer.NewDSLRenderer(os.Stdout)
	err := renderer.Render(workspace)
	if err != nil {
		fmt.Printf("Error rendering: %v\n", err)
		os.Exit(1)
	}
}
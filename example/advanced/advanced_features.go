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
	// Create a workspace
	workspace := gostructurizr.Workspace().
		WithName("Advanced Features Example").
		WithDesc("An example demonstrating filtered views and advanced styling")
	
	model := workspace.Model()
	
	// Set enterprise
	model.SetEnterprise("ACME Corporation")
	
	// Create people, systems, and containers
	customer := model.AddPerson("Customer", "A customer of ACME Corporation")
	
	internetBankingSystem := model.AddSoftwareSystem("Internet Banking System", "Allows customers to view account information and make transactions")
	internetBankingSystem.WithTag("Web")
	
	mainframeBankingSystem := model.AddSoftwareSystem("Mainframe Banking System", "Stores all core banking information")
	mainframeBankingSystem.WithTag("Mainframe")
	
	emailSystem := model.AddSoftwareSystem("E-mail System", "Sends e-mails to customers")
	emailSystem.WithTag("External")
	
	// Define relationships
	customer.Uses(internetBankingSystem, "Uses")
	internetBankingSystem.Uses(mainframeBankingSystem, "Gets account information from")
	internetBankingSystem.Uses(emailSystem, "Sends e-mail using")
	
	// Add containers to the Internet Banking System
	webApplication := internetBankingSystem.AddContainer("Web Application", "Provides Internet banking functionality via the web", "Java and Spring MVC")
	webApplication.WithTag("WebApp")
	
	mobileApp := internetBankingSystem.AddContainer("Mobile App", "Provides Internet banking functionality via a mobile device", "Xamarin")
	mobileApp.WithTag("MobileApp")
	
	apiApplication := internetBankingSystem.AddContainer("API Application", "Provides an API for use by the mobile app and other systems", "Java and Spring Boot")
	apiApplication.WithTag("API")
	
	database := internetBankingSystem.AddContainer("Database", "Stores user registration information, hashed auth credentials, access logs, etc.", "Oracle")
	database.WithTag("Database")
	
	// Define relationships between containers
	webApplication.Uses(apiApplication, "Uses")
	mobileApp.Uses(apiApplication, "Uses")
	apiApplication.Uses(database, "Reads from and writes to")
	apiApplication.Uses(mainframeBankingSystem, "Uses")
	apiApplication.Uses(emailSystem, "Uses")
	
	// Create deployment elements
	
	// Production environment
	// Create AWS cloud deployment node
	awsCloud := model.AddDeploymentNode("Amazon Web Services", "Cloud services platform", "Amazon Web Services", gostructurizr.ProductionEnvironment)
	awsCloud.WithLocation(gostructurizr.ExternalLocation)
	
	// Create a web server instance
	webServer := awsCloud.AddDeploymentNode("Web Server", "Web server nodes", "Amazon EC2", gostructurizr.ProductionEnvironment)
	webAppInstance := webServer.AddContainerInstance(webApplication)
	
	// Create an app server instance
	appServer := awsCloud.AddDeploymentNode("App Server", "Application server nodes", "Amazon EC2", gostructurizr.ProductionEnvironment)
	apiAppInstance := appServer.AddContainerInstance(apiApplication)
	
	// Create a database server
	dbServer := awsCloud.AddDeploymentNode("Database Server", "Database server nodes", "Amazon RDS", gostructurizr.ProductionEnvironment)
	dbInstance := dbServer.AddContainerInstance(database)
	
	// Create relationships between deployment nodes
	webAppInstance.Uses(apiAppInstance, "Makes API calls to").WithTechnology("HTTPS")
	apiAppInstance.Uses(dbInstance, "Reads from and writes to").WithTechnology("JDBC")
	
	// Create views
	views := workspace.Views()
	
	// System context view
	contextView := views.CreateSystemContextView(internetBankingSystem)
	contextView.WithKey("SystemContext")
	contextView.WithDescription("The system context view for the Internet Banking System")
	contextView.AddAllElements()
	contextView.AddAllPeople()
	contextView.WithAutoLayout()
	
	// Container view
	containerView := views.CreateContainerView(internetBankingSystem)
	containerView.WithKey("Containers")
	containerView.WithDescription("The containers within the Internet Banking System")
	containerView.AddAllContainers()
	containerView.AddSoftwareSystem(mainframeBankingSystem)
	containerView.AddSoftwareSystem(emailSystem)
	containerView.AddAllPeople()
	containerView.WithAutoLayout()
	
	// Deployment view
	deploymentView := views.CreateDeploymentView(internetBankingSystem, gostructurizr.ProductionEnvironment)
	deploymentView.WithKey("Deployment")
	deploymentView.WithDescription("The production deployment scenario for the Internet Banking System")
	deploymentView.AddDeploymentNode(awsCloud)
	deploymentView.WithAutoLayout()
	
	// Filtered views
	// Create a filtered view showing only API components
	apiFilteredView := views.CreateFilteredView(containerView, "API Components")
	apiFilteredView.WithKey("APIComponents")
	apiFilteredView.WithDescription("Shows only API components")
	apiFilteredView.Include("API")
	apiFilteredView.WithAutoLayout()
	
	// Create a filtered view excluding database components
	noDatabaseFilteredView := views.CreateFilteredView(containerView, "No Database Components")
	noDatabaseFilteredView.WithKey("NoDatabases")
	noDatabaseFilteredView.WithDescription("Excludes database components")
	noDatabaseFilteredView.Exclude("Database")
	noDatabaseFilteredView.WithAutoLayout()
	
	// Styling
	styles := views.Configuration().Styles()
	
	// Basic styling
	styles.AddElementStyle(tags.Person).
		WithBackground("#08427b").
		WithColor("#ffffff").
		WithShape(shapes.Person)
	
	styles.AddElementStyle(tags.SoftwareSystem).
		WithBackground("#1168bd").
		WithColor("#ffffff")
	
	styles.AddElementStyle(tags.Container).
		WithBackground("#438dd5").
		WithColor("#ffffff")
	
	// Custom tag styling
	styles.AddElementStyle(tags.Tag("Database")).
		WithBackground("#1168bd").
		WithColor("#ffffff").
		WithShape(shapes.Cylinder)
	
	styles.AddElementStyle(tags.Tag("WebApp")).
		WithBackground("#1168bd").
		WithColor("#ffffff").
		WithShape(shapes.WebBrowser)
	
	styles.AddElementStyle(tags.Tag("MobileApp")).
		WithBackground("#1168bd").
		WithColor("#ffffff").
		WithShape(shapes.MobileDevicePortrait)
	
	styles.AddElementStyle(tags.Tag("API")).
		WithBackground("#1168bd").
		WithColor("#ffffff").
		WithShape(shapes.Component)
	
	// Add deployment node styling
	styles.AddElementStyle(tags.DeploymentNode).
		WithBackground("#999999").
		WithColor("#ffffff")
	
	// Add external system styling
	styles.AddElementStyle(tags.Tag("External")).
		WithBackground("#999999").
		WithColor("#ffffff").
		WithStrokeWidth(2)
	
	// Advanced element styling
	styles.AddElementStyle(tags.Person).
		WithBorder(2).
		WithBorderStyle(gostructurizr.Solid).
		WithShadow(true).
		WithFontFamily(gostructurizr.SansSerif).
		WithFontStyle("normal").
		WithFontSize(24)
	
	// Advanced relationship styling
	syncStyle := styles.AddAdvancedRelationshipStyle(tags.Synchronous)
	syncStyle.WithDashed().
		WithColor("#ff0000").
		WithWidth(2).
		WithFontSize(12).
		WithFontColor("#ff0000").
		WithArrow()
	
	// Generate the DSL
	var b strings.Builder
	r := renderer.NewDSLRenderer(&b)
	if err := r.Render(workspace); err != nil {
		panic(fmt.Sprintf("renderer didn't succeed: %s", err))
	}
	fmt.Println(b.String())
}
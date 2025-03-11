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
		WithName("Deployment Example").
		WithDesc("Example showing deployment nodes and infrastructure")
	
	model := workspace.Model()
	
	// Define the enterprise context
	model.SetEnterprise("ACME Financial")
	
	// Define people
	customer := model.AddPerson("Customer", "A customer of ACME Financial")
	support := model.AddPerson("Customer Support", "Customer support staff")
	
	// Define software systems
	onlineBankingSystem := model.AddSoftwareSystem("Online Banking System", "Allows customers to view accounts and make transactions")
	mainframeBankingSystem := model.AddSoftwareSystem("Mainframe Banking System", "Stores core banking information")
	
	// Define relationships between people and systems
	customer.Uses(onlineBankingSystem, "Uses")
	support.Uses(onlineBankingSystem, "Uses to support customers")
	onlineBankingSystem.Uses(mainframeBankingSystem, "Gets account information from")
	
	// Define containers for the online banking system
	webApp := onlineBankingSystem.AddContainer("Web Application", "Provides online banking functionality to customers", "Java and Spring MVC")
	apiApp := onlineBankingSystem.AddContainer("API Application", "Provides API for mobile and web applications", "Java and Spring Boot")
	database := onlineBankingSystem.AddContainer("Database", "Stores user data, sessions, etc.", "Oracle")
	
	// Define container relationships
	webApp.Uses(apiApp, "Makes API calls to").WithTechnology("HTTPS")
	apiApp.Uses(database, "Reads from and writes to").WithTechnology("JDBC")
	apiApp.Uses(mainframeBankingSystem, "Makes API calls to").WithTechnology("WebSphere MQ")
	
	// Define deployment nodes
	// Production environment
	awsCloud := model.AddDeploymentNode("AWS", "Amazon Web Services", "Cloud Infrastructure", gostructurizr.ProductionEnvironment)
	
	// Web tier with auto-scaling group
	webTier := awsCloud.AddDeploymentNode("Web Tier", "Web application tier", "Amazon EC2 Auto Scaling Group", gostructurizr.ProductionEnvironment)
	webServer1 := webTier.AddDeploymentNode("Web Server 1", "Web server instance", "Amazon EC2", gostructurizr.ProductionEnvironment)
	webServer2 := webTier.AddDeploymentNode("Web Server 2", "Web server instance", "Amazon EC2", gostructurizr.ProductionEnvironment)
	
	// App tier with auto-scaling group
	appTier := awsCloud.AddDeploymentNode("App Tier", "Application tier", "Amazon EC2 Auto Scaling Group", gostructurizr.ProductionEnvironment)
	appServer1 := appTier.AddDeploymentNode("App Server 1", "Application server instance", "Amazon EC2", gostructurizr.ProductionEnvironment)
	appServer2 := appTier.AddDeploymentNode("App Server 2", "Application server instance", "Amazon EC2", gostructurizr.ProductionEnvironment)
	
	// Database tier
	dbTier := awsCloud.AddDeploymentNode("Database Tier", "Database tier", "Amazon RDS", gostructurizr.ProductionEnvironment)
	primaryDB := dbTier.AddDeploymentNode("Primary DB", "Primary database", "Oracle RDS", gostructurizr.ProductionEnvironment)
	standbyDB := dbTier.AddDeploymentNode("Standby DB", "Standby database", "Oracle RDS", gostructurizr.ProductionEnvironment)
	
	// Infrastructure nodes
	loadBalancer := awsCloud.AddInfrastructureNode("Load Balancer", "Elastic Load Balancer", "AWS ELB")
	appLoadBalancer := appTier.AddInfrastructureNode("Internal Load Balancer", "Internal load balancer", "AWS ELB")
	
	// Add container instances
	webAppInstance1 := webServer1.AddContainerInstance(webApp)
	webAppInstance2 := webServer2.AddContainerInstance(webApp)
	apiAppInstance1 := appServer1.AddContainerInstance(apiApp)
	apiAppInstance2 := appServer2.AddContainerInstance(apiApp)
	databaseInstance := primaryDB.AddContainerInstance(database)
	
	// Add health checks
	webAppInstance1.AddHealthCheck("Web Status", "https://web1.example.com/health")
	webAppInstance2.AddHealthCheck("Web Status", "https://web2.example.com/health")
	apiAppInstance1.AddHealthCheck("API Status", "https://api1.example.com/health")
	apiAppInstance2.AddHealthCheck("API Status", "https://api2.example.com/health")
	databaseInstance.AddHealthCheck("DB Status", "https://db.example.com/health")
	
	// External infrastructure
	onPremises := model.AddDeploymentNode("On-Premises", "On-premises infrastructure", "Data Center", gostructurizr.ProductionEnvironment)
	onPremises.WithLocation(gostructurizr.InternalLocation)
	
	mainframe := onPremises.AddDeploymentNode("Mainframe", "Mainframe system", "IBM z/OS", gostructurizr.ProductionEnvironment)
	
	// Define relationships between deployment nodes
	loadBalancer.Uses(webServer1, "Routes requests to").WithTechnology("HTTPS")
	loadBalancer.Uses(webServer2, "Routes requests to").WithTechnology("HTTPS")
	webAppInstance1.Uses(appLoadBalancer, "Makes API calls to").WithTechnology("HTTPS")
	webAppInstance2.Uses(appLoadBalancer, "Makes API calls to").WithTechnology("HTTPS")
	appLoadBalancer.Uses(appServer1, "Routes requests to").WithTechnology("HTTPS")
	appLoadBalancer.Uses(appServer2, "Routes requests to").WithTechnology("HTTPS")
	apiAppInstance1.Uses(databaseInstance, "Reads from and writes to").WithTechnology("JDBC")
	apiAppInstance2.Uses(databaseInstance, "Reads from and writes to").WithTechnology("JDBC")
	primaryDB.Uses(standbyDB, "Replicates data to").WithTechnology("Oracle Data Guard")
	apiAppInstance1.Uses(mainframe, "Makes API calls to").WithTechnology("WebSphere MQ")
	apiAppInstance2.Uses(mainframe, "Makes API calls to").WithTechnology("WebSphere MQ")
	
	// Create views
	views := workspace.Views()
	
	// System Context view
	contextView := views.CreateSystemContextView(onlineBankingSystem)
	contextView.WithKey("SystemContext")
	contextView.WithDescription("System Context diagram for the Online Banking System")
	contextView.AddAllElements()
	contextView.AddAllPeople()
	contextView.WithAutoLayout()
	
	// Container view
	containerView := views.CreateContainerView(onlineBankingSystem)
	containerView.WithKey("Containers")
	containerView.WithDescription("Container diagram for the Online Banking System")
	containerView.AddAllContainers()
	containerView.AddAllPeople()
	containerView.AddSoftwareSystem(mainframeBankingSystem)
	containerView.WithAutoLayout()
	
	// Deployment view
	deploymentView := views.CreateDeploymentView(onlineBankingSystem, gostructurizr.ProductionEnvironment)
	deploymentView.WithKey("ProductionDeployment")
	deploymentView.WithDescription("Production deployment diagram for the Online Banking System")
	deploymentView.AddDeploymentNode(awsCloud)
	deploymentView.AddDeploymentNode(onPremises)
	deploymentView.AddAllRelationships()
	deploymentView.WithAutoLayout()
	
	// Add styling
	styles := views.Configuration().Styles()
	styles.AddElementStyle(tags.Person).WithBackground("#08427b").WithColor("#ffffff").WithShape(shapes.Person)
	styles.AddElementStyle(tags.SoftwareSystem).WithBackground("#1168bd").WithColor("#ffffff")
	styles.AddElementStyle(tags.Container).WithBackground("#438dd5").WithColor("#ffffff")
	styles.AddElementStyle(tags.DeploymentNode).WithBackground("#999999").WithColor("#ffffff")
	styles.AddElementStyle(tags.InfrastructureNode).WithBackground("#85bbf0").WithColor("#000000")
	
	// Generate the DSL
	var b strings.Builder
	r := renderer.NewDSLRenderer(&b)
	if err := r.Render(workspace); err != nil {
		panic(fmt.Sprintf("renderer didn't succeed: %s", err))
	}
	fmt.Println(b.String())
}
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
		WithName("Deployment Environments Example").
		WithDesc("An example of deployment environments with infrastructure and container instances")

	// Create the model
	model := workspace.Model()

	// Create users and systems
	customer := model.AddPerson("Customer", "A customer of the online store")

	webStore := model.AddSoftwareSystem("Web Store", "Online retail system")
	webStore.Tags().Add("WebStore")

	// Define containers
	webApplication := webStore.AddContainer("Web Application", "The main web application", "Java and Spring Boot")
	database := webStore.AddContainer("Database", "Customer and order information", "MySQL")
	cache := webStore.AddContainer("Cache", "Caches product information", "Redis")

	// Define relationships
	customer.Uses(webApplication, "Visits website using")
	webApplication.Uses(database, "Reads from and writes to")
	webApplication.Uses(cache, "Reads from and writes to")

	// Create views
	views := workspace.Views()

	// Add a container view
	containerView := views.CreateContainerView(webStore).
		WithKey("Containers").
		WithDescription("Container view for Web Store")
	containerView.WithAutoLayout()
	containerView.AddAllContainers()
	containerView.AddAllPeople()

	// Add deployment views for multiple environments
	// Development environment
	devView := views.CreateDevView(webStore).
		WithKey("DevelopmentDeployment").
		WithDescription("Development deployment")
	devView.WithAutoLayout()

	// Create deployment nodes
	devServer := model.AddDeploymentNode("Developer Laptop", "Developer Laptop", "Windows 10", strukt.DevelopmentEnvironment)

	dockerEngine := devServer.AddChildNode("Docker Engine", "Docker Engine", "Docker CE")

	webServerContainer := dockerEngine.AddChildNode("Web Server", "Web Server", "Docker Container: Tomcat")
	webAppInstance := webServerContainer.AddContainerInstance(webApplication)
	webAppInstance.AddHealthCheck("Web App Health", "http://localhost:8080/actuator/health").
		WithInterval(30).WithTimeout(2000)

	dbContainer := dockerEngine.AddChildNode("Database Server", "Database Server", "Docker Container: MySQL")
	dbContainer.AddContainerInstance(database)

	cacheContainer := dockerEngine.AddChildNode("Cache Server", "Cache Server", "Docker Container: Redis")
	cacheContainer.AddContainerInstance(cache)

	devView.Add(devServer)

	// Production environment
	prodView := views.CreateProdView(webStore).
		WithKey("ProductionDeployment").
		WithDescription("Production deployment")
	prodView.WithAutoLayout()

	// Create AWS infrastructure
	amazonWebServices := model.AddDeploymentNode("Amazon Web Services", "Cloud Provider", "AWS", strukt.ProductionEnvironment)

	amazonRegion := amazonWebServices.AddChildNode("US-East-1", "Region", "AWS us-east-1")

	// Create multiple zones for redundancy
	zoneA := amazonRegion.AddChildNode("us-east-1a", "Availability Zone", "AWS us-east-1a")
	zoneB := amazonRegion.AddChildNode("us-east-1b", "Availability Zone", "AWS us-east-1b")

	// Load balancer
	elb := amazonRegion.AddInfrastructureNode("Elastic Load Balancer", "ELB", "AWS")

	// Web tier in zone A
	webTierA := zoneA.AddChildNode("Web Tier", "Auto Scaling Group", "AWS EC2")
	webServerA := webTierA.AddChildNode("Web Server", "Amazon EC2", "Amazon Linux")
	webAppInstanceA := webServerA.AddContainerInstance(webApplication)
	webAppInstanceA.AddHealthCheck("Health", "http://web-app/actuator/health").
		WithInterval(60).WithTimeout(5000)

	// Web tier in zone B
	webTierB := zoneB.AddChildNode("Web Tier", "Auto Scaling Group", "AWS EC2")
	webServerB := webTierB.AddChildNode("Web Server", "Amazon EC2", "Amazon Linux")
	webAppInstanceB := webServerB.AddContainerInstance(webApplication)
	webAppInstanceB.AddHealthCheck("Health", "http://web-app/actuator/health").
		WithInterval(60).WithTimeout(5000)

	// Database tier
	rdsInstance := amazonRegion.AddInfrastructureNode("RDS MySQL", "Amazon RDS", "AWS RDS")
	// Note: We can't directly add container instances to infrastructure nodes in this implementation
	// In a real implementation, we'd need to add this functionality

	// Cache - ElastiCache
	elastiCache := amazonRegion.AddInfrastructureNode("ElastiCache", "Amazon ElastiCache", "AWS ElastiCache Redis")
	// Infrastructure nodes don't support container instances in this implementation

	// Add relationships between infrastructure
	webServerA.Uses(rdsInstance, "Connects to")
	webServerA.Uses(elastiCache, "Reads from and writes to")
	webServerB.Uses(rdsInstance, "Connects to")
	webServerB.Uses(elastiCache, "Reads from and writes to")
	elb.Uses(webServerA, "Routes requests to")
	elb.Uses(webServerB, "Routes requests to")

	prodView.Add(amazonWebServices)

	// Add styles
	styles := views.Configuration().Styles()

	// Element styles
	styles.AddElementStyle(tags.Person).
		WithShape(shapes.Person).
		WithBackground("#08427B").
		WithColor("#ffffff")

	styles.AddElementStyle(tags.Container).
		WithBackground("#438DD5").
		WithColor("#ffffff")

	styles.AddElementStyle(tags.DeploymentNode).
		WithBackground("#ffffff").
		WithColor("#000000")

	styles.AddElementStyle(tags.InfrastructureNode).
		WithBackground("#C5E6FF").
		WithColor("#000000").
		WithShape(shapes.Ellipse)

	styles.AddElementStyle(tags.ContainerInstance).
		WithBackground("#438DD5").
		WithColor("#ffffff").
		WithBorder(2)

	// Write the DSL to stdout
	renderer := renderer.NewDSLRenderer(os.Stdout)
	err := renderer.Render(workspace)
	if err != nil {
		fmt.Printf("Error rendering: %v\n", err)
		os.Exit(1)
	}
}
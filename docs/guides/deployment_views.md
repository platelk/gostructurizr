# Working with Deployment Views

Deployment views in the C4 model show how software systems and containers map to infrastructure. This guide explains how to create deployment views with GoStructurizr.

## Why Deployment Views?

Deployment views help answer critical questions about your system's runtime environment:

- Where are your containers deployed?
- What infrastructure elements support your system?
- How are containers replicated across environments?
- What are the technology choices for deployment?

## Basic Deployment View Concepts

In GoStructurizr, deployment modeling involves several key elements:

- **Deployment Environment**: A context representing different deployment zones (Dev, Test, Staging, Production)
- **Deployment Node**: Infrastructure elements (e.g., cloud regions, servers, containers)
- **Container Instance**: A runtime instance of a container from your container diagram
- **Infrastructure Node**: Supporting infrastructure services (e.g., load balancers, queues)

## Creating a Basic Deployment View

Here's how to create a simple deployment view:

```go
// Get your model with previously defined software system and containers
workspace := gostructurizr.Workspace().WithName("Deployment Example")
model := workspace.Model()

// Define your system and containers first
system := model.AddSoftwareSystem("Banking System", "Core banking platform")
webApp := system.AddContainer("Web Application", "Provides UI for users", "Go, Gin")
api := system.AddContainer("API Service", "Backend API", "Go, gRPC")
database := system.AddContainer("Database", "Stores user data", "PostgreSQL")

// Define a deployment environment
productionEnv := model.AddDeploymentEnvironment("Production")

// Define top-level infrastructure
awsCloud := productionEnv.AddDeploymentNode("AWS", "Amazon Web Services", "Cloud provider")
usEastRegion := awsCloud.AddDeploymentNode("us-east-1", "US East Region", "AWS Region")

// Define compute resources
ec2WebTier := usEastRegion.AddDeploymentNode("EC2 Web Tier", "Web Server Instances", "t3.medium EC2")
webAppInstance := ec2WebTier.AddContainerInstance(webApp, "Web Application Instance", "Containerized web UI")

ec2ApiTier := usEastRegion.AddDeploymentNode("EC2 API Tier", "API Server Instances", "t3.large EC2")
apiInstance := ec2ApiTier.AddContainerInstance(api, "API Service Instance", "Containerized API")

rdsDbTier := usEastRegion.AddDeploymentNode("RDS", "Managed Database", "db.m5.large RDS")
dbInstance := rdsDbTier.AddContainerInstance(database, "Production Database", "PostgreSQL 13")

// Define relationships
webAppInstance.Uses(apiInstance, "Makes API calls to", "JSON/HTTPS")
apiInstance.Uses(dbInstance, "Reads from and writes to", "SQL/TCP")

// Create a deployment view
views := workspace.Views()
deploymentView := views.CreateDeploymentView(system, "Production").
    WithKey("ProductionDeployment").
    WithDescription("Deployment diagram for the Banking System in production.")

// Add all deployment nodes
deploymentView.AddAllDeploymentNodes()

// Apply auto-layout
deploymentView.WithAutoLayout()
```

## Working with Multiple Environments

One of the strengths of GoStructurizr is modeling different deployment environments:

```go
// Define environments
devEnv := model.AddDeploymentEnvironment("Development")
stagingEnv := model.AddDeploymentEnvironment("Staging")
prodEnv := model.AddDeploymentEnvironment("Production")

// Development environment
devServer := devEnv.AddDeploymentNode("Developer Workstation", "Local development", "MacBook Pro")
devServer.AddContainerInstance(webApp)
devServer.AddContainerInstance(api)
devDbServer := devEnv.AddDeploymentNode("Dev Database", "Development database", "Docker")
devDbServer.AddContainerInstance(database)

// Staging environment (simplified production-like)
stagingAws := stagingEnv.AddDeploymentNode("AWS", "Staging AWS Account", "Cloud")
stagingRegion := stagingAws.AddDeploymentNode("us-east-1", "US East Region", "AWS Region")
stagingEcs := stagingRegion.AddDeploymentNode("ECS Cluster", "Elastic Container Service", "t3.medium")
stagingEcs.AddContainerInstance(webApp)
stagingEcs.AddContainerInstance(api)
stagingRds := stagingRegion.AddDeploymentNode("RDS", "Staging Database", "db.t3.medium")
stagingRds.AddContainerInstance(database)

// Create deployment views for each environment
devDeploymentView := views.CreateDeploymentView(system, "Development").
    WithKey("DevelopmentDeployment").
    WithDescription("Development environment deployment.")
devDeploymentView.AddAllDeploymentNodes()

stagingDeploymentView := views.CreateDeploymentView(system, "Staging").
    WithKey("StagingDeployment").
    WithDescription("Staging environment deployment.")
stagingDeploymentView.AddAllDeploymentNodes()

prodDeploymentView := views.CreateDeploymentView(system, "Production").
    WithKey("ProductionDeployment").
    WithDescription("Production environment deployment.")
prodDeploymentView.AddAllDeploymentNodes()
```

## Adding Infrastructure Nodes

Infrastructure nodes represent supporting infrastructure that doesn't directly host your containers:

```go
// Add supporting infrastructure in production
elb := usEastRegion.AddInfrastructureNode("ELB", "Elastic Load Balancer", "Application Load Balancer")
cdn := awsCloud.AddInfrastructureNode("CloudFront", "Content Delivery Network", "AWS CloudFront")
waf := awsCloud.AddInfrastructureNode("WAF", "Web Application Firewall", "AWS WAF")

// Define relationships
waf.Uses(cdn, "Filters requests to")
cdn.Uses(elb, "Routes requests to")
elb.Uses(webAppInstance, "Routes requests to")
```

## Adding Health Checks

Health checks can be added to deployment nodes and infrastructure nodes:

```go
// Add health checks to nodes
webAppInstance.AddHealthCheck("HTTP Check", "/health").
    WithInterval(30).
    WithTimeout(5)

apiInstance.AddHealthCheck("gRPC Health", "/grpc.health.v1.Health/Check").
    WithInterval(30).
    WithTimeout(10)

dbInstance.AddHealthCheck("Database Connection", "SELECT 1").
    WithInterval(60).
    WithTimeout(15)

elb.AddHealthCheck("ELB Status Check", "/status").
    WithInterval(60)
```

## Styling Deployment Elements

Apply specific styles to your deployment nodes and infrastructure nodes:

```go
styles := views.Configuration().Styles()

// Style deployment nodes by environment
styles.AddElementStyle("Development").
    WithBackground("#cfcfcf").
    WithColor("#000000")

styles.AddElementStyle("Staging").
    WithBackground("#d7c365").
    WithColor("#000000")

styles.AddElementStyle("Production").
    WithBackground("#b22222").
    WithColor("#ffffff")

// Style infrastructure nodes
styles.AddElementStyle("Infrastructure Node").
    WithBackground("#aaaaaa").
    WithColor("#000000")

// Style container instances
styles.AddElementStyle("Container Instance").
    WithBackground("#438dd5").
    WithColor("#ffffff")
```

## Deployment Node Properties

You can add custom properties to deployment nodes:

```go
// Add properties to describe infrastructure details
ec2WebTier.WithProperties(map[string]string{
    "Instance Type": "t3.medium",
    "Auto Scaling": "2-6 instances",
    "Operating System": "Amazon Linux 2",
    "Memory": "4 GB",
    "Disk": "100 GB EBS",
})

rdsDbTier.WithProperties(map[string]string{
    "Instance Type": "db.m5.large",
    "Multi-AZ": "true",
    "Storage": "100 GB",
    "Encryption": "true",
    "Backup Retention": "7 days",
})
```

## Nested Deployment Nodes

Deployment nodes can be nested to represent hierarchical infrastructure:

```go
// Create a more complex production environment with nested nodes
awsProductionEnv := prodEnv.AddDeploymentNode("AWS", "Amazon Web Services", "Cloud provider")

// VPC
vpc := awsProductionEnv.AddDeploymentNode("VPC", "Production VPC", "10.0.0.0/16")

// Availability Zones
az1 := vpc.AddDeploymentNode("us-east-1a", "Availability Zone 1", "AZ")
az2 := vpc.AddDeploymentNode("us-east-1b", "Availability Zone 2", "AZ")

// Public subnets
publicSubnet1 := az1.AddDeploymentNode("Public Subnet", "Public subnet", "10.0.1.0/24")
publicSubnet2 := az2.AddDeploymentNode("Public Subnet", "Public subnet", "10.0.2.0/24")

// Private subnets
privateSubnet1 := az1.AddDeploymentNode("Private Subnet", "Private subnet", "10.0.3.0/24")
privateSubnet2 := az2.AddDeploymentNode("Private Subnet", "Private subnet", "10.0.4.0/24")

// Database subnets
dbSubnet1 := az1.AddDeploymentNode("Database Subnet", "Database subnet", "10.0.5.0/24")
dbSubnet2 := az2.AddDeploymentNode("Database Subnet", "Database subnet", "10.0.6.0/24")

// Add applications to subnets
webServer1 := publicSubnet1.AddDeploymentNode("Web Server 1", "Web server instance", "t3.medium EC2")
webServer1.AddContainerInstance(webApp)

webServer2 := publicSubnet2.AddDeploymentNode("Web Server 2", "Web server instance", "t3.medium EC2")
webServer2.AddContainerInstance(webApp)

apiServer1 := privateSubnet1.AddDeploymentNode("API Server 1", "API server instance", "t3.large EC2")
apiServer1.AddContainerInstance(api)

apiServer2 := privateSubnet2.AddDeploymentNode("API Server 2", "API server instance", "t3.large EC2")
apiServer2.AddContainerInstance(api)

dbCluster := dbSubnet1.AddDeploymentNode("RDS Cluster", "Database cluster", "Aurora PostgreSQL")
dbPrimary := dbCluster.AddDeploymentNode("Primary DB", "Primary database instance", "db.r5.large")
dbPrimary.AddContainerInstance(database)

dbReplica := dbSubnet2.AddDeploymentNode("RDS Replica", "Read replica", "db.r5.large")
dbReplica.AddContainerInstance(database, "Read Replica")
```

## Complete Example

For a complete deployment view example, see the [deployment example in the examples directory](../examples/deployment_example.md).

## Resources

- [Structurizr Deployment View Documentation](https://structurizr.com/help/documentation/deployment-diagram)
- [C4 Model Deployment Diagrams](https://c4model.com/#DeploymentDiagram)
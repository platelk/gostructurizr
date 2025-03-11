# Microservices Architecture Example

This example demonstrates how to model a microservices architecture using GoStructurizr. It includes container views, component breakdowns, deployment modeling, and dynamic views showing service interactions.

## Complete Code Example

```go
package main

import (
    "fmt"
    "strings"
    "github.com/platelk/gostructurizr"
    "github.com/platelk/gostructurizr/renderer"
    "github.com/platelk/gostructurizr/shapes"
    "github.com/platelk/gostructurizr/tags"
)

func main() {
    // Create a workspace for our e-commerce microservices architecture
    workspace := gostructurizr.Workspace().
        WithName("E-Commerce Microservices").
        WithDesc("Microservices architecture for an e-commerce platform")

    model := workspace.Model()

    // Define users and external systems
    customer := model.AddPerson("Customer", "A customer of the e-commerce platform")
    adminUser := model.AddPerson("Admin User", "Admin staff managing the platform")
    paymentProvider := model.AddSoftwareSystem("Payment Gateway", "External payment processing service")
    emailProvider := model.AddSoftwareSystem("Email Service", "External email delivery service")
    warehouseSystem := model.AddSoftwareSystem("Warehouse System", "External warehouse management system")

    // Define the e-commerce system
    ecommerce := model.AddSoftwareSystem("E-Commerce Platform", "Online shopping platform")
    customer.Uses(ecommerce, "Shops online using")
    adminUser.Uses(ecommerce, "Manages products, orders and inventory using")
    ecommerce.Uses(paymentProvider, "Processes payments using")
    ecommerce.Uses(emailProvider, "Sends emails using")
    ecommerce.Uses(warehouseSystem, "Sends fulfillment requests to")

    // Define API Gateway
    apiGateway := ecommerce.AddContainer("API Gateway", "Gateway for all API requests", "Go, Echo")
    customer.Uses(apiGateway, "Makes API calls to", "HTTPS")
    adminUser.Uses(apiGateway, "Makes API calls to", "HTTPS")

    // Define Web and Mobile applications
    webApp := ecommerce.AddContainer("Web Application", "E-commerce web frontend", "React, TypeScript")
    webApp.Uses(apiGateway, "Makes API calls to", "HTTPS")
    customer.Uses(webApp, "Uses", "HTTPS")
    adminUser.Uses(webApp, "Uses", "HTTPS")

    mobileApp := ecommerce.AddContainer("Mobile Application", "E-commerce mobile app", "React Native")
    mobileApp.Uses(apiGateway, "Makes API calls to", "HTTPS")
    customer.Uses(mobileApp, "Uses")

    // Define frontend BFF (Backend for Frontend)
    customerBFF := ecommerce.AddContainer("Customer BFF", "Backend for customer-facing frontends", "Go, Gin")
    adminBFF := ecommerce.AddContainer("Admin BFF", "Backend for admin frontends", "Go, Gin")
    apiGateway.Uses(customerBFF, "Routes customer requests to", "gRPC")
    apiGateway.Uses(adminBFF, "Routes admin requests to", "gRPC")

    // Define microservices
    productService := ecommerce.AddContainer("Product Service", "Manages product catalog", "Go")
    customerService := ecommerce.AddContainer("Customer Service", "Manages customer information", "Go")
    orderService := ecommerce.AddContainer("Order Service", "Manages orders", "Go")
    cartService := ecommerce.AddContainer("Cart Service", "Manages shopping carts", "Go")
    inventoryService := ecommerce.AddContainer("Inventory Service", "Manages product inventory", "Go")
    paymentService := ecommerce.AddContainer("Payment Service", "Handles payment processing", "Go")
    notificationService := ecommerce.AddContainer("Notification Service", "Sends notifications", "Go")
    searchService := ecommerce.AddContainer("Search Service", "Handles product search", "Go, Elasticsearch")
    
    // BFF to microservice connections
    customerBFF.Uses(productService, "Gets product information from", "gRPC")
    customerBFF.Uses(customerService, "Gets customer information from", "gRPC")
    customerBFF.Uses(orderService, "Gets and creates orders via", "gRPC")
    customerBFF.Uses(cartService, "Manages cart via", "gRPC")
    customerBFF.Uses(searchService, "Searches products via", "gRPC")
    
    adminBFF.Uses(productService, "Manages products via", "gRPC")
    adminBFF.Uses(orderService, "Manages orders via", "gRPC")
    adminBFF.Uses(inventoryService, "Manages inventory via", "gRPC")
    adminBFF.Uses(customerService, "Views customer data via", "gRPC")

    // Service-to-service communication
    orderService.Uses(inventoryService, "Checks product availability via", "gRPC")
    orderService.Uses(cartService, "Gets cart information via", "gRPC")
    orderService.Uses(customerService, "Gets customer information via", "gRPC")
    orderService.Uses(paymentService, "Processes payments via", "gRPC")
    orderService.Uses(notificationService, "Requests order notifications via", "Kafka")
    paymentService.Uses(notificationService, "Requests payment notifications via", "Kafka")
    paymentService.Uses(paymentProvider, "Processes payments using", "HTTPS")
    notificationService.Uses(emailProvider, "Sends emails using", "HTTPS")
    inventoryService.Uses(warehouseSystem, "Updates inventory in", "HTTPS")
    searchService.Uses(productService, "Indexes product data from", "Kafka")

    // Define databases
    productDB := ecommerce.AddContainer("Product Database", "Stores product information", "PostgreSQL")
    customerDB := ecommerce.AddContainer("Customer Database", "Stores customer information", "PostgreSQL")
    orderDB := ecommerce.AddContainer("Order Database", "Stores order information", "PostgreSQL")
    cartDB := ecommerce.AddContainer("Cart Database", "Stores shopping cart data", "Redis")
    inventoryDB := ecommerce.AddContainer("Inventory Database", "Stores inventory levels", "PostgreSQL")
    searchDB := ecommerce.AddContainer("Search Database", "Stores search indexes", "Elasticsearch")
    
    // Connect services to databases
    productService.Uses(productDB, "Reads from and writes to", "SQL/TCP")
    customerService.Uses(customerDB, "Reads from and writes to", "SQL/TCP")
    orderService.Uses(orderDB, "Reads from and writes to", "SQL/TCP")
    cartService.Uses(cartDB, "Reads from and writes to", "Redis Protocol")
    inventoryService.Uses(inventoryDB, "Reads from and writes to", "SQL/TCP")
    searchService.Uses(searchDB, "Reads from and writes to", "REST/HTTP")

    // Add message broker
    messageBroker := ecommerce.AddContainer("Message Broker", "Handles async communication between services", "Kafka")
    orderService.Uses(messageBroker, "Publishes order events to", "Kafka Protocol")
    productService.Uses(messageBroker, "Publishes product events to", "Kafka Protocol")
    inventoryService.Uses(messageBroker, "Publishes inventory events to", "Kafka Protocol")
    paymentService.Uses(messageBroker, "Publishes payment events to", "Kafka Protocol")
    notificationService.Uses(messageBroker, "Subscribes to events from", "Kafka Protocol")
    searchService.Uses(messageBroker, "Subscribes to product events from", "Kafka Protocol")

    // Add service components for the Order Service
    orderController := orderService.AddComponent("Order Controller", "Handles order API requests", "Go, gRPC")
    orderManager := orderService.AddComponent("Order Manager", "Contains order business logic", "Go")
    orderRepository := orderService.AddComponent("Order Repository", "Data access for orders", "Go, SQL")
    paymentClient := orderService.AddComponent("Payment Client", "Client for the Payment Service", "Go, gRPC")
    inventoryClient := orderService.AddComponent("Inventory Client", "Client for the Inventory Service", "Go, gRPC")
    orderEventPublisher := orderService.AddComponent("Order Event Publisher", "Publishes order events", "Go, Kafka")

    // Connect order service components
    customerBFF.Uses(orderController, "Creates and fetches orders via", "gRPC")
    adminBFF.Uses(orderController, "Manages orders via", "gRPC")
    orderController.Uses(orderManager, "Uses")
    orderManager.Uses(orderRepository, "Uses")
    orderRepository.Uses(orderDB, "Reads from and writes to", "SQL/TCP")
    orderManager.Uses(paymentClient, "Uses")
    orderManager.Uses(inventoryClient, "Uses")
    orderManager.Uses(orderEventPublisher, "Uses")
    paymentClient.Uses(paymentService, "Makes API calls to", "gRPC")
    inventoryClient.Uses(inventoryService, "Makes API calls to", "gRPC")
    orderEventPublisher.Uses(messageBroker, "Publishes events to", "Kafka Protocol")

    // Create views
    views := workspace.Views()

    // System Context view
    contextView := views.CreateSystemContextView(ecommerce).
        WithKey("SystemContext").
        WithDescription("System Context diagram for the E-Commerce Platform")
    contextView.AddAllElements()
    contextView.WithAutoLayout()

    // Container view
    containerView := views.CreateContainerView(ecommerce).
        WithKey("Containers").
        WithDescription("Container diagram for the E-Commerce Platform")
    containerView.AddAllContainers()
    containerView.AddPerson(customer)
    containerView.AddPerson(adminUser)
    containerView.AddSoftwareSystem(paymentProvider)
    containerView.AddSoftwareSystem(emailProvider)
    containerView.AddSoftwareSystem(warehouseSystem)
    containerView.WithAutoLayout()

    // Order Service component view
    componentView := views.CreateComponentView(orderService).
        WithKey("OrderServiceComponents").
        WithDescription("Component diagram for the Order Service")
    componentView.AddAllComponents()
    componentView.Add(customerBFF)
    componentView.Add(adminBFF)
    componentView.Add(paymentService)
    componentView.Add(inventoryService)
    componentView.Add(orderDB)
    componentView.Add(messageBroker)
    componentView.WithAutoLayout()

    // Create dynamic view for order processing
    dynamicView := views.CreateDynamicView(ecommerce).
        WithKey("OrderProcessing").
        WithDescription("Shows how orders are processed")

    // Define order processing scenario
    dynamicView.Add(customer, "Places order", webApp)
    dynamicView.Add(webApp, "Submits order", apiGateway, "HTTPS/JSON")
    dynamicView.Add(apiGateway, "Forwards request", customerBFF, "gRPC")
    dynamicView.Add(customerBFF, "Creates order", orderService, "gRPC")
    dynamicView.Add(orderService, "Checks inventory", inventoryService, "gRPC")
    dynamicView.Add(inventoryService, "Reserves inventory", inventoryDB, "SQL/TCP")
    dynamicView.Add(inventoryDB, "Returns reservation status", inventoryService)
    dynamicView.Add(inventoryService, "Returns availability", orderService)
    dynamicView.Add(orderService, "Gets cart items", cartService, "gRPC")
    dynamicView.Add(cartService, "Retrieves cart", cartDB, "Redis Protocol")
    dynamicView.Add(cartDB, "Returns cart items", cartService)
    dynamicView.Add(cartService, "Returns cart details", orderService)
    dynamicView.Add(orderService, "Processes payment", paymentService, "gRPC")
    dynamicView.Add(paymentService, "Authorizes payment", paymentProvider, "HTTPS")
    dynamicView.Add(paymentProvider, "Returns payment status", paymentService)
    dynamicView.Add(paymentService, "Returns payment confirmation", orderService)
    dynamicView.Add(orderService, "Stores order", orderDB, "SQL/TCP")
    dynamicView.Add(orderService, "Publishes order created event", messageBroker, "Kafka")
    dynamicView.Add(messageBroker, "Delivers order event", notificationService, "Kafka")
    dynamicView.Add(notificationService, "Sends order confirmation", emailProvider, "HTTPS")
    dynamicView.Add(emailProvider, "Delivers email", customer, "Email")
    dynamicView.Add(orderService, "Returns order confirmation", customerBFF)
    dynamicView.Add(customerBFF, "Returns order details", apiGateway)
    dynamicView.Add(apiGateway, "Returns order confirmation", webApp)
    dynamicView.Add(webApp, "Shows order confirmation", customer)

    // Create production deployment view
    productionEnv := model.AddDeploymentEnvironment("Production")
    
    // Cloud provider
    aws := productionEnv.AddDeploymentNode("AWS", "Amazon Web Services", "Cloud Provider")
    
    // VPC
    vpc := aws.AddDeploymentNode("VPC", "Production VPC", "10.0.0.0/16")
    
    // Load balancer
    loadBalancer := vpc.AddInfrastructureNode("Load Balancer", "Application Load Balancer", "AWS ELB")
    
    // Web hosting
    cloudfrontCDN := aws.AddInfrastructureNode("CloudFront", "Content Delivery Network", "AWS CloudFront")
    s3Bucket := aws.AddDeploymentNode("S3", "Static Website Hosting", "AWS S3")
    s3Bucket.AddContainerInstance(webApp)
    
    // Kubernetes cluster
    eksCluster := vpc.AddDeploymentNode("EKS", "Kubernetes Cluster", "Amazon EKS")
    
    // API namespace
    apiNamespace := eksCluster.AddDeploymentNode("API Namespace", "Kubernetes namespace for API services", "Namespace")
    apiGatewayPod := apiNamespace.AddDeploymentNode("API Gateway Pod", "API Gateway Replica Set", "Pod")
    apiGatewayPod.AddContainerInstance(apiGateway)
    
    customerBFFPod := apiNamespace.AddDeploymentNode("Customer BFF Pod", "Customer BFF Replica Set", "Pod")
    customerBFFPod.AddContainerInstance(customerBFF)
    
    adminBFFPod := apiNamespace.AddDeploymentNode("Admin BFF Pod", "Admin BFF Replica Set", "Pod")
    adminBFFPod.AddContainerInstance(adminBFF)
    
    // Services namespace
    servicesNamespace := eksCluster.AddDeploymentNode("Services Namespace", "Kubernetes namespace for microservices", "Namespace")
    
    productPod := servicesNamespace.AddDeploymentNode("Product Pod", "Product Service Replica Set", "Pod")
    productPod.AddContainerInstance(productService)
    
    customerPod := servicesNamespace.AddDeploymentNode("Customer Pod", "Customer Service Replica Set", "Pod")
    customerPod.AddContainerInstance(customerService)
    
    orderPod := servicesNamespace.AddDeploymentNode("Order Pod", "Order Service Replica Set", "Pod")
    orderPod.AddContainerInstance(orderService)
    
    cartPod := servicesNamespace.AddDeploymentNode("Cart Pod", "Cart Service Replica Set", "Pod")
    cartPod.AddContainerInstance(cartService)
    
    inventoryPod := servicesNamespace.AddDeploymentNode("Inventory Pod", "Inventory Service Replica Set", "Pod")
    inventoryPod.AddContainerInstance(inventoryService)
    
    paymentPod := servicesNamespace.AddDeploymentNode("Payment Pod", "Payment Service Replica Set", "Pod")
    paymentPod.AddContainerInstance(paymentService)
    
    notificationPod := servicesNamespace.AddDeploymentNode("Notification Pod", "Notification Service Replica Set", "Pod")
    notificationPod.AddContainerInstance(notificationService)
    
    searchPod := servicesNamespace.AddDeploymentNode("Search Pod", "Search Service Replica Set", "Pod")
    searchPod.AddContainerInstance(searchService)
    
    // Databases
    dbNamespace := eksCluster.AddDeploymentNode("Database Namespace", "Kubernetes namespace for databases", "Namespace")
    
    kafkaBrokerNode := dbNamespace.AddDeploymentNode("Kafka Broker", "Kafka message broker", "StatefulSet")
    kafkaBrokerNode.AddContainerInstance(messageBroker)
    
    // RDS and ElastiCache
    rdsSubnet := vpc.AddDeploymentNode("RDS Subnet", "Database subnet", "Subnet")
    
    rdsProductDB := rdsSubnet.AddDeploymentNode("Product DB", "Product Database Instance", "RDS PostgreSQL")
    rdsProductDB.AddContainerInstance(productDB)
    
    rdsCustomerDB := rdsSubnet.AddDeploymentNode("Customer DB", "Customer Database Instance", "RDS PostgreSQL")
    rdsCustomerDB.AddContainerInstance(customerDB)
    
    rdsOrderDB := rdsSubnet.AddDeploymentNode("Order DB", "Order Database Instance", "RDS PostgreSQL")
    rdsOrderDB.AddContainerInstance(orderDB)
    
    rdsInventoryDB := rdsSubnet.AddDeploymentNode("Inventory DB", "Inventory Database Instance", "RDS PostgreSQL")
    rdsInventoryDB.AddContainerInstance(inventoryDB)
    
    elasticacheNode := vpc.AddDeploymentNode("ElastiCache", "Redis Cluster", "Amazon ElastiCache")
    elasticacheNode.AddContainerInstance(cartDB)
    
    elasticsearchNode := vpc.AddDeploymentNode("Elasticsearch", "Search Database", "Amazon Elasticsearch Service")
    elasticsearchNode.AddContainerInstance(searchDB)
    
    // Define connections
    loadBalancer.Uses(cloudfrontCDN, "Routes requests to")
    cloudfrontCDN.Uses(s3Bucket, "Serves content from")
    loadBalancer.Uses(apiGatewayPod, "Routes API requests to")
    
    // Create deployment view
    deploymentView := views.CreateDeploymentView(ecommerce, "Production").
        WithKey("ProductionDeployment").
        WithDescription("Deployment diagram for the E-Commerce Platform in AWS")
    deploymentView.AddAllDeploymentNodes()
    deploymentView.WithAutoLayout()

    // Apply styles
    styles := views.Configuration().Styles()

    // Style people
    styles.AddElementStyle(tags.Person).
        WithBackground("#08427b").
        WithColor("#ffffff").
        WithShape(shapes.Person)

    // Style external systems
    styles.AddElementStyle(tags.SoftwareSystem).
        WithBackground("#1168bd").
        WithColor("#ffffff")
    
    // External systems get different color
    styles.AddElementStyle("External").
        WithBackground("#999999").
        WithColor("#ffffff")
    paymentProvider.AddTags("External")
    emailProvider.AddTags("External")
    warehouseSystem.AddTags("External")

    // Style containers
    styles.AddElementStyle(tags.Container).
        WithBackground("#438dd5").
        WithColor("#ffffff")
    
    // Style components
    styles.AddElementStyle(tags.Component).
        WithBackground("#85bbf0").
        WithColor("#000000")
    
    // Style database containers
    styles.AddElementStyle("Database").
        WithShape(shapes.Cylinder).
        WithBackground("#438dd5").
        WithColor("#ffffff")
    productDB.AddTags("Database")
    customerDB.AddTags("Database")
    orderDB.AddTags("Database")
    cartDB.AddTags("Database")
    inventoryDB.AddTags("Database")
    searchDB.AddTags("Database")

    // Style message broker
    styles.AddElementStyle("Message Broker").
        WithBackground("#85bbf0").
        WithColor("#000000")
    messageBroker.AddTags("Message Broker")

    // Style relationships
    styles.AddRelationshipStyle(tags.Relationship).
        WithThickness(2).
        WithColor("#707070").
        WithDashed(false)
    
    styles.AddRelationshipStyle("gRPC").
        WithColor("#5a9c40")
    
    styles.AddRelationshipStyle("Kafka").
        WithColor("#d04a43").
        WithDashed(true)

    // Generate DSL
    var b strings.Builder
    r := renderer.NewDSLRenderer(&b)
    if err := r.Render(workspace); err != nil {
        panic(fmt.Sprintf("renderer didn't succeed: %s", err))
    }
    fmt.Println(b.String())
}
```

## What This Example Demonstrates

This comprehensive example demonstrates:

1. **Complete Microservices Architecture:**
   - Separation of concerns with distinct microservices
   - Backend-for-Frontend (BFF) pattern
   - API Gateway pattern
   - Service-to-service communication

2. **Multiple View Types:**
   - System Context view
   - Container view showing all microservices
   - Component view drilling into a specific service
   - Dynamic view showing an order processing flow
   - Deployment view showing the production infrastructure

3. **Architectural Patterns:**
   - Event-driven architecture using Kafka
   - REST and gRPC communication
   - Database-per-service pattern
   - Containerized deployment with Kubernetes

4. **Cloud-Native Deployment:**
   - Kubernetes-based deployment
   - Managed cloud services (RDS, ElastiCache, etc.)
   - CDN and load balancing
   - Production-ready infrastructure

## Key Design Decisions

### Service Boundaries

Services are divided based on domain boundaries following Domain-Driven Design principles:
- **Product Service**: Handles product catalog management
- **Customer Service**: Manages customer information
- **Order Service**: Processes and manages orders
- **Cart Service**: Manages shopping carts
- **Inventory Service**: Tracks product availability
- **Payment Service**: Processes payments
- **Notification Service**: Sends emails and notifications
- **Search Service**: Provides product search functionality

### Communication Patterns

Multiple communication styles are demonstrated:
- **Synchronous**: Using gRPC for service-to-service calls where immediate response is needed
- **Asynchronous**: Using Kafka for event-based communication where eventual consistency is acceptable
- **REST**: For external system integration

### Data Management

Each service has its own database, following the database-per-service pattern:
- PostgreSQL for structured data (products, orders, customers)
- Redis for ephemeral data (shopping carts)
- Elasticsearch for search functionality

### Deployment Strategy

The deployment demonstrates a real-world Kubernetes-based architecture:
- Services deployed as Kubernetes pods
- Namespace-based logical separation
- AWS managed services for databases
- Load balancing and CDN for frontend

## Extending This Example

You can extend this example by:

1. Adding authentication and authorization services
2. Including observability infrastructure (logging, monitoring, tracing)
3. Adding a recommendation service with machine learning capabilities
4. Implementing a CQRS pattern with read/write segregation
5. Adding a data analytics pipeline
6. Adding staging and development deployment environments
7. Creating filtered views for specific stakeholders

## Resources

- [Structurizr DSL Documentation](https://structurizr.com/dsl)
- [C4 Model for Visualizing Architecture](https://c4model.com)
- [Microservices Patterns](https://microservices.io/patterns/)
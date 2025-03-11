# Dynamic Views in GoStructurizr

While static views (context, container, component) show the structure of your architecture, dynamic views illustrate how elements interact to fulfill specific scenarios. This guide explains how to create and use dynamic views in GoStructurizr.

## Why Dynamic Views?

Dynamic views serve several key purposes:

- Show how elements interact at runtime
- Visualize specific user journeys or use cases
- Illustrate key processes and data flows
- Highlight synchronous vs. asynchronous communication
- Document complex sequences and dependencies

## Basic Concepts

A dynamic view in GoStructurizr consists of the following elements:

- **Scope**: The software system or container being described
- **Key**: A unique identifier for the view
- **Elements**: The participants in the interaction
- **Relationships**: The interactions between elements (with sequence)
- **Description**: What scenario the view represents

## Creating a Basic Dynamic View

Here's how to create a simple dynamic view for a user login scenario:

```go
// First define your model with people, systems and containers
workspace := gostructurizr.Workspace().WithName("Dynamic View Example")
model := workspace.Model()

// Define the elements
user := model.AddPerson("User", "A user of the system")
webApp := model.AddSoftwareSystem("Banking System", "Banking application")
webUI := webApp.AddContainer("Web Application", "Provides UI", "React")
api := webApp.AddContainer("API", "Backend API", "Go")
authService := webApp.AddContainer("Auth Service", "Handles authentication", "Go")
database := webApp.AddContainer("Database", "Stores user data", "PostgreSQL")

// Create a dynamic view
views := workspace.Views()
dynamicView := views.CreateDynamicView(webApp).
    WithKey("UserLogin").
    WithDescription("Shows the login process for a user")

// Add interactions in sequence
dynamicView.Add(user, "Enters credentials", webUI)
dynamicView.Add(webUI, "Sends login request", api)
dynamicView.Add(api, "Validates credentials", authService)
dynamicView.Add(authService, "Retrieves user data", database)
dynamicView.Add(database, "Returns user profile", authService)
dynamicView.Add(authService, "Returns authentication token", api)
dynamicView.Add(api, "Returns authentication token", webUI)
dynamicView.Add(webUI, "Displays dashboard", user)

// Apply auto-layout
dynamicView.WithAutoLayout()
```

## Adding Context to Interactions

Add more context to each interaction:

```go
// Add detailed descriptions and technology information
dynamicView.Add(user, "Enters username and password in login form", webUI)
dynamicView.Add(webUI, "POSTs to /api/login with credentials", api, "HTTPS")
dynamicView.Add(api, "Calls validateCredentials() with username/password", authService, "gRPC")
dynamicView.Add(authService, "Executes SELECT query for user credentials", database, "SQL over TCP")
dynamicView.Add(database, "Returns user record with hashed password", authService)
dynamicView.Add(authService, "Generates and signs JWT token", api)
dynamicView.Add(api, "Returns 200 OK with JWT token in response body", webUI)
dynamicView.Add(webUI, "Stores token in local storage and navigates to dashboard", user)
```

## Parallel Interactions

Show interactions that happen in parallel:

```go
// User registration with parallel processing
dynamicView := views.CreateDynamicView(webApp).
    WithKey("UserRegistration").
    WithDescription("Shows the user registration process")

dynamicView.Add(user, "Fills registration form", webUI)
dynamicView.Add(webUI, "Submits registration data", api)
dynamicView.Add(api, "Creates user account", authService)
dynamicView.Add(authService, "Stores user data", database)

// Parallel processes start here
dynamicView.StartParallelSequence()

// First parallel process - email notification
dynamicView.Add(authService, "Sends registration confirmation", emailService)
dynamicView.Add(emailService, "Sends welcome email", user)

// Switch to second parallel process
dynamicView.NextParallelSequence()

// Second parallel process - analytics
dynamicView.Add(authService, "Logs new user registration", analyticsService)
dynamicView.Add(analyticsService, "Updates user metrics", metricsDatabase)

// End parallel sequences
dynamicView.EndParallelSequence()

// Continue with normal flow
dynamicView.Add(api, "Returns success response", webUI)
dynamicView.Add(webUI, "Shows success message", user)
```

## Container-Scoped Dynamic Views

You can scope a dynamic view to a specific container to show component interactions:

```go
// First define your components within a container
paymentProcessor := api.AddComponent("Payment Processor", "Processes payments", "Go")
fraudDetection := api.AddComponent("Fraud Detection", "Detects suspicious transactions", "Go")
paymentGateway := api.AddComponent("Payment Gateway Connector", "Connects to payment provider", "Go")
notificationService := api.AddComponent("Notification Service", "Sends notifications", "Go")

// Create a container-scoped dynamic view
dynamicView := views.CreateDynamicViewForContainer(api).
    WithKey("PaymentProcessing").
    WithDescription("Shows how payments are processed")

// Define interaction sequence
dynamicView.Add(webUI, "Submits payment", paymentProcessor)
dynamicView.Add(paymentProcessor, "Validates payment data", fraudDetection)
dynamicView.Add(fraudDetection, "Returns risk assessment", paymentProcessor)
dynamicView.Add(paymentProcessor, "Processes payment", paymentGateway)
dynamicView.Add(paymentGateway, "Sends payment to provider", externalPaymentSystem)
dynamicView.Add(externalPaymentSystem, "Confirms payment", paymentGateway)
dynamicView.Add(paymentGateway, "Returns payment confirmation", paymentProcessor)
dynamicView.Add(paymentProcessor, "Requests payment notification", notificationService)
dynamicView.Add(notificationService, "Sends payment receipt", user)
dynamicView.Add(paymentProcessor, "Returns payment result", webUI)
```

## Styling Dynamic Views

Style dynamic view elements and interactions:

```go
// Get the styles configuration
styles := views.Configuration().Styles()

// Style specific elements in dynamic views
styles.AddElementStyle("User").
    WithBackground("#08427b").
    WithColor("#ffffff").
    WithShape(shapes.Person)

// Style asynchronous communications
styles.AddRelationshipStyle("Asynchronous").
    WithDashed(true).
    WithColor("#ff0000")

// Add specific styling for a relationship in the view
paymentNotificationRelationship := paymentProcessor.Uses(notificationService, "Requests payment notification")
paymentNotificationRelationship.AddTags("Asynchronous")
```

## Best Practices for Dynamic Views

1. **Focus on one scenario**: Each dynamic view should illustrate a single, specific scenario
2. **Keep it simple**: Limit to 8-10 interactions when possible
3. **Use meaningful descriptions**: Clearly describe what happens in each interaction
4. **Show technology details**: Include protocols and data formats where relevant
5. **Group related views**: Create multiple dynamic views for related but distinct scenarios
6. **Differentiate sync/async**: Use styling to distinguish between synchronous and asynchronous communications
7. **Include error paths**: Consider creating separate views for important error scenarios
8. **Use consistent naming**: Ensure interaction descriptions are consistent in style and terminology

## Advanced Example: Order Processing

Here's a more complex example showing an order processing flow:

```go
// Define the elements
customer := model.AddPerson("Customer", "A customer of the e-commerce system")
webStore := model.AddSoftwareSystem("E-commerce System", "Online store")
webApp := webStore.AddContainer("Web Application", "Web store UI", "React")
mobileApp := webStore.AddContainer("Mobile App", "Mobile store UI", "Flutter")
apiGateway := webStore.AddContainer("API Gateway", "API entry point", "Go")
orderService := webStore.AddContainer("Order Service", "Manages orders", "Go")
inventoryService := webStore.AddContainer("Inventory Service", "Manages inventory", "Go")
paymentService := webStore.AddContainer("Payment Service", "Processes payments", "Go")
notificationService := webStore.AddContainer("Notification Service", "Sends notifications", "Go")
orderDB := webStore.AddContainer("Order Database", "Stores orders", "MongoDB")
inventoryDB := webStore.AddContainer("Inventory Database", "Stores inventory", "PostgreSQL")
paymentProvider := model.AddSoftwareSystem("Payment Provider", "External payment processor")

// Create a dynamic view for order placement
dynamicView := views.CreateDynamicView(webStore).
    WithKey("OrderPlacement").
    WithDescription("Shows how a customer places an order")

// Customer journey from different entry points
dynamicView.Add(customer, "Places order", webApp)

// Order processing sequence
dynamicView.Add(webApp, "Submits order", apiGateway, "HTTPS/JSON")
dynamicView.Add(apiGateway, "Creates order", orderService, "gRPC")
dynamicView.Add(orderService, "Validates inventory", inventoryService, "gRPC")
dynamicView.Add(inventoryService, "Checks stock levels", inventoryDB, "SQL")
dynamicView.Add(inventoryDB, "Returns available items", inventoryService)
dynamicView.Add(inventoryService, "Reserves inventory", inventoryDB, "SQL")
dynamicView.Add(inventoryService, "Returns inventory status", orderService)

// Payment processing
dynamicView.Add(orderService, "Requests payment", paymentService, "gRPC")
dynamicView.Add(paymentService, "Processes payment", paymentProvider, "HTTPS/JSON")
dynamicView.Add(paymentProvider, "Confirms payment", paymentService)
dynamicView.Add(paymentService, "Returns payment confirmation", orderService)
dynamicView.Add(orderService, "Stores order", orderDB, "MongoDB Driver")

// Parallel notifications
dynamicView.StartParallelSequence()

// Email notification
dynamicView.Add(orderService, "Requests order confirmation", notificationService, "Kafka")
dynamicView.Add(notificationService, "Sends order confirmation", customer, "Email")

// Inventory updates
dynamicView.NextParallelSequence()
dynamicView.Add(orderService, "Confirms inventory deduction", inventoryService, "gRPC")
dynamicView.Add(inventoryService, "Updates inventory levels", inventoryDB, "SQL")

// End parallel
dynamicView.EndParallelSequence()

// Return to customer
dynamicView.Add(orderService, "Returns order confirmation", apiGateway)
dynamicView.Add(apiGateway, "Returns order confirmation", webApp)
dynamicView.Add(webApp, "Displays order confirmation", customer)
```

## Working with Multiple Dynamic Views

Create multiple dynamic views to show different scenarios or variations:

```go
// Create several dynamic views for different scenarios
loginView := views.CreateDynamicView(webStore).
    WithKey("UserLogin").
    WithDescription("Shows the login process")
// ... add interactions

registrationView := views.CreateDynamicView(webStore).
    WithKey("UserRegistration").
    WithDescription("Shows the registration process")
// ... add interactions  

orderPlacementView := views.CreateDynamicView(webStore).
    WithKey("OrderPlacement").
    WithDescription("Shows the order placement process")
// ... add interactions

orderCancellationView := views.CreateDynamicView(webStore).
    WithKey("OrderCancellation").
    WithDescription("Shows the order cancellation process")
// ... add interactions
```

## Resources

- [Structurizr Dynamic Views Documentation](https://structurizr.com/help/dynamic-diagram)
- [C4 model Dynamic Diagrams](https://c4model.com/#DynamicDiagram)
- [UML Sequence Diagrams](https://www.uml-diagrams.org/sequence-diagrams.html) (for inspiration)
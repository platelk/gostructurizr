# Getting Started with GoStructurizr

This guide will walk you through the basic steps to create software architecture diagrams using GoStructurizr.

## Prerequisites

- Go 1.16 or newer
- Basic understanding of the [C4 model](../concepts/c4_model.md)

## Installation

Start by installing the GoStructurizr package:

```bash
go get github.com/platelk/gostructurizr
```

## Basic Structure of a GoStructurizr Program

Every GoStructurizr program follows this general structure:

1. Create a workspace
2. Define the model (people, systems, containers, components)
3. Create views of the model
4. Apply styles to elements
5. Generate the Structurizr DSL output

## Step 1: Create a Workspace

The workspace is the top-level container for your architecture model:

```go
workspace := gostructurizr.Workspace().
    WithName("Banking System").
    WithDesc("Architecture model for our banking system.")

// Access the model and views
model := workspace.Model()
views := workspace.Views()
```

## Step 2: Define Your Model

The model contains all the elements of your architecture - people, systems, containers, and components:

```go
// Add a person (user)
customer := model.AddPerson("Customer", "A customer of the bank.")

// Add a software system
bankingSystem := model.AddSoftwareSystem("Banking System", "Handles customer accounts and transactions.")

// Create a relationship between them
customer.Uses(bankingSystem, "Manages accounts and makes transactions")

// Add containers within the system
webApp := bankingSystem.AddContainer("Web Application", "Provides banking functionality to customers.", "Go, Gin")
mobileApp := bankingSystem.AddContainer("Mobile App", "Provides banking functionality on mobile devices.", "Flutter")
api := bankingSystem.AddContainer("API", "Handles account transactions.", "Go, gRPC")
database := bankingSystem.AddContainer("Database", "Stores customer information and transactions.", "PostgreSQL")

// Define relationships between containers
customer.Uses(webApp, "Uses")
customer.Uses(mobileApp, "Uses")
webApp.Uses(api, "Makes API calls to")
mobileApp.Uses(api, "Makes API calls to")
api.Uses(database, "Reads from and writes to")

// Add components within a container
accountsController := api.AddComponent("Accounts Controller", "Handles account-related API requests.", "Go")
transactionsController := api.AddComponent("Transactions Controller", "Handles transaction-related API requests.", "Go")
accountsService := api.AddComponent("Accounts Service", "Business logic for accounts.", "Go")
transactionsService := api.AddComponent("Transactions Service", "Business logic for transactions.", "Go")

// Define relationships between components
webApp.Uses(accountsController, "Makes API calls to", "JSON/HTTPS")
webApp.Uses(transactionsController, "Makes API calls to", "JSON/HTTPS")
accountsController.Uses(accountsService, "Uses")
transactionsController.Uses(transactionsService, "Uses")
accountsService.Uses(database, "Reads from and writes to", "SQL/TCP")
transactionsService.Uses(database, "Reads from and writes to", "SQL/TCP")
```

## Step 3: Create Views

Views provide different perspectives of your model:

```go
// System Context view
contextView := views.CreateSystemContextView(bankingSystem).
    WithKey("SystemContext").
    WithDescription("System Context diagram for the banking system.")
contextView.AddAllElements()

// Container view
containerView := views.CreateContainerView(bankingSystem).
    WithKey("Containers").
    WithDescription("Container diagram for the banking system.")
containerView.AddAllContainers()
containerView.AddPerson(customer)

// Component view
componentView := views.CreateComponentView(api).
    WithKey("ApiComponents").
    WithDescription("Component diagram for the API.")
componentView.AddAllComponents()
componentView.Add(webApp)
componentView.Add(mobileApp)
componentView.Add(database)

// Set up automatic layout
contextView.WithAutoLayout()
containerView.WithAutoLayout()
componentView.WithAutoLayout()
```

## Step 4: Apply Styles

Apply styles to make your diagrams more visually appealing and follow the C4 model conventions:

```go
styles := views.Configuration().Styles()

// Style people
styles.AddElementStyle(tags.Person).
    WithShape(shapes.Person).
    WithBackground("#08427b").
    WithColor("#ffffff")

// Style software systems
styles.AddElementStyle(tags.SoftwareSystem).
    WithBackground("#1168bd").
    WithColor("#ffffff")

// Style containers by technology
styles.AddElementStyle(tags.Container).
    WithBackground("#438dd5").
    WithColor("#ffffff")

styles.AddElementStyle("Database").
    WithShape(shapes.Cylinder).
    WithBackground("#438dd5").
    WithColor("#ffffff")

// Style components
styles.AddElementStyle(tags.Component).
    WithBackground("#85bbf0").
    WithColor("#000000")

// Style relationships
styles.AddRelationshipStyle(tags.Relationship).
    WithThickness(2).
    WithColor("#707070").
    WithDashed(false)
```

## Step 5: Generate DSL Output

Finally, render your model to Structurizr DSL:

```go
var b strings.Builder
r := renderer.NewDSLRenderer(&b)
if err := r.Render(workspace); err != nil {
    panic(fmt.Sprintf("renderer didn't succeed: %s", err))
}
fmt.Println(b.String())
```

## Complete Example

Here's a complete example showing a simple banking system architecture:

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
    // Create workspace
    workspace := gostructurizr.Workspace().
        WithName("Banking System").
        WithDesc("Architecture model for our banking system.")
    model := workspace.Model()
    views := workspace.Views()
    
    // Define model elements
    customer := model.AddPerson("Customer", "A customer of the bank.")
    bankingSystem := model.AddSoftwareSystem("Banking System", "Handles customer accounts and transactions.")
    customer.Uses(bankingSystem, "Manages accounts and makes transactions")
    
    // Define containers
    webApp := bankingSystem.AddContainer("Web Application", "Provides banking functionality to customers.", "Go, Gin")
    api := bankingSystem.AddContainer("API", "Handles account transactions.", "Go, gRPC")
    database := bankingSystem.AddContainer("Database", "Stores customer information and transactions.", "PostgreSQL")
    customer.Uses(webApp, "Uses")
    webApp.Uses(api, "Makes API calls to")
    api.Uses(database, "Reads from and writes to")
    
    // Create views
    contextView := views.CreateSystemContextView(bankingSystem).
        WithKey("SystemContext").
        WithDescription("System Context diagram for the banking system.")
    contextView.AddAllElements()
    
    containerView := views.CreateContainerView(bankingSystem).
        WithKey("Containers").
        WithDescription("Container diagram for the banking system.")
    containerView.AddAllContainers()
    containerView.AddPerson(customer)
    
    // Apply auto-layout
    contextView.WithAutoLayout()
    containerView.WithAutoLayout()
    
    // Apply styles
    styles := views.Configuration().Styles()
    styles.AddElementStyle(tags.Person).
        WithShape(shapes.Person).
        WithBackground("#08427b").
        WithColor("#ffffff")
    styles.AddElementStyle(tags.SoftwareSystem).
        WithBackground("#1168bd").
        WithColor("#ffffff")
    styles.AddElementStyle(tags.Container).
        WithBackground("#438dd5").
        WithColor("#ffffff")
    styles.AddElementStyle("Database").
        WithShape(shapes.Cylinder).
        WithBackground("#438dd5").
        WithColor("#ffffff")
    
    // Generate DSL
    var b strings.Builder
    r := renderer.NewDSLRenderer(&b)
    if err := r.Render(workspace); err != nil {
        panic(fmt.Sprintf("renderer didn't succeed: %s", err))
    }
    fmt.Println(b.String())
}
```

## Next Steps

Now that you understand the basics, explore:

- [Advanced styling techniques](./advanced_styling.md)
- [Creating deployment views](./deployment_views.md)
- [Working with dynamic views](./dynamic_views.md)
- [Using filtered views](./filtered_views.md)

For more examples, check out the [examples directory](../examples/).
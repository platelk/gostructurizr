# The C4 Model

The C4 model, created by Simon Brown, is a way of describing and communicating software architecture through a hierarchy of diagrams. It was designed to be simple and intuitive, allowing developers and non-technical stakeholders alike to understand system architecture.

## Core Concepts

The C4 model consists of four levels of abstraction, each providing a different perspective on the system:

### Level 1: System Context Diagram

The System Context diagram represents the highest level of abstraction. It shows:

- Your software system as a single box in the center
- The people (end users, actors, roles, personas) who use it
- Other software systems that yours interacts with

**Purpose**: To establish how your system fits into the world around it. This diagram answers the question, "What are we building and who uses it?"

```go
// Creating a System Context Diagram
user := model.AddPerson("User", "A user of my software system.")
system := model.AddSoftwareSystem("My System", "Core banking system.")
externalAPI := model.AddSoftwareSystem("Payment Gateway", "External payment processing.")

user.Uses(system, "Manages accounts")
system.Uses(externalAPI, "Processes payments")

contextView := views.CreateSystemContextView(system).
    WithKey("SystemContext").
    WithDescription("System Context diagram for the banking system.")
contextView.AddAllElements()
```

### Level 2: Container Diagram

The Container diagram zooms into your software system, showing the high-level technical building blocks ("containers"):

- Applications (web applications, mobile apps)
- Data stores (databases, file systems)
- Microservices or smaller deployable units
- Infrastructure components

**Purpose**: To show the major technology decisions and how responsibilities are distributed across containers. This answers, "What are the technology building blocks and how do they interact?"

```go
// Creating a Container Diagram
webApp := system.AddContainer("Web Application", "Provides banking functionality to customers.", "Go, Gin")
api := system.AddContainer("API", "Handles account transactions.", "Go, gRPC")
database := system.AddContainer("Database", "Stores customer information and transactions.", "PostgreSQL")

webApp.Uses(api, "Makes API calls to")
api.Uses(database, "Reads from and writes to")
user.Uses(webApp, "Uses")

containerView := views.CreateContainerView(system).
    WithKey("Containers").
    WithDescription("Container diagram for the banking system.")
containerView.AddAllContainers()
containerView.AddPerson(user)
```

### Level 3: Component Diagram

The Component diagram zooms into an individual container to show its internal components:

- Classes, modules, or logical groupings of code
- Their responsibilities and interactions

**Purpose**: To decompose containers into components, showing how a container is constructed. This answers, "What are the major structural building blocks and their interactions?"

```go
// Creating a Component Diagram
accountController := api.AddComponent("Account Controller", "Handles account API requests.", "Go")
transactionController := api.AddComponent("Transaction Controller", "Handles transaction API requests.", "Go")
accountService := api.AddComponent("Account Service", "Business logic for accounts.", "Go")
transactionService := api.AddComponent("Transaction Service", "Business logic for transactions.", "Go")

webApp.Uses(accountController, "Makes API calls to", "JSON/HTTPS")
webApp.Uses(transactionController, "Makes API calls to", "JSON/HTTPS")
accountController.Uses(accountService, "Uses")
transactionController.Uses(transactionService, "Uses")
accountService.Uses(database, "Reads from and writes to", "SQL/TCP")
transactionService.Uses(database, "Reads from and writes to", "SQL/TCP")

componentView := views.CreateComponentView(api).
    WithKey("ApiComponents").
    WithDescription("Component diagram for the API.")
componentView.AddAllComponents()
componentView.Add(webApp)
componentView.Add(database)
```

### Level 4: Code Diagram

The Code diagram is optional and shows how a component is implemented:

- Classes, interfaces, packages
- Their relationships and dependencies

**Purpose**: To show how components are implemented as code. This is typically represented using UML class diagrams or similar.

## Visual Language

The C4 model uses a consistent set of shapes to represent different elements:

- **Person**: A person who interacts with the system (typically shown as a stick figure)
- **Software System**: The highest level of abstraction for the software (box)
- **Container**: An application or data store within your system (box)
- **Component**: A logical grouping of code within a container (box)

GoStructurizr provides styling functionality to ensure your diagrams follow these conventions:

```go
// Setting up standard C4 model styling
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

styles.AddElementStyle(tags.Component).
    WithBackground("#85bbf0").
    WithColor("#000000")
```

## Resources

- [C4 Model Official Website](https://c4model.com/)
- [Structurizr DSL Documentation](https://structurizr.com/dsl)
- Simon Brown's [The C4 Model for Visualising Software Architecture](https://www.infoq.com/articles/C4-architecture-model/)
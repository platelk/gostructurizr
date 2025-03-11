# Comparing GoStructurizr with Official Structurizr

This document outlines the similarities and differences between GoStructurizr and the official Structurizr tooling. It will help you understand the current capabilities of GoStructurizr and what features might be added in the future.

## Overview

GoStructurizr is a Go implementation of the [Structurizr DSL](https://structurizr.com/dsl), which allows you to create software architecture models based on the [C4 model](https://c4model.com) programmatically in Go. The official Structurizr provides tools in Java, .NET, and other languages, along with a web-based service for visualizing and sharing diagrams.

## Core Concepts Supported in Both

Both GoStructurizr and official Structurizr support these core concepts:

1. **C4 Model Elements**:
   - Person
   - Software System
   - Container
   - Component
   - Relationships
   - Tags

2. **C4 View Types**:
   - System Context View
   - Container View
   - Component View
   - Deployment View
   - Dynamic View
   - Filtered View

3. **Styling**:
   - Element styling (colors, shapes, borders)
   - Relationship styling (line styles, colors, thickness)
   - Tags for applying consistent styling

4. **DSL Generation**:
   - Both produce Structurizr DSL output

## Features Available in GoStructurizr

GoStructurizr currently supports:

1. **Core Model Creation**:
   - Building hierarchical models with software systems, containers, and components
   - Defining relationships between elements
   - Creating different view types
   - Applying styles to elements and relationships

2. **Deployment Modeling**:
   - Multiple deployment environments (Dev, Test, Staging, Production)
   - Deployment nodes with hierarchical structure
   - Container instances
   - Infrastructure nodes

3. **Dynamic Views**:
   - Sequential interactions
   - Parallel sequences
   - Nested scopes

4. **Advanced Styling**:
   - Custom tags
   - Shape customization
   - Detailed styling properties

5. **Filtered Views**:
   - Tag-based filtering
   - Include/exclude patterns

## Features Unique to Official Structurizr

The official Structurizr includes several features not currently implemented in GoStructurizr:

1. **Documentation Integration**:
   - Architecture Decision Records (ADRs)
   - Embedded documentation with Markdown/AsciiDoc
   - Documentation templates (Arc42, etc.)

2. **Workspace Management**:
   - Versioning and history tracking
   - Workspace sharing and collaboration
   - Access control and permissions

3. **Visualization**:
   - Interactive diagram rendering
   - Image export (PNG, SVG)
   - Diagram editor

4. **Integrations**:
   - PlantUML export
   - Mermaid export
   - WebSequenceDiagrams export

5. **Additional View Types**:
   - System Landscape view
   - Timelines and roadmaps
   - Custom themes

6. **Enterprise Features**:
   - Workspace locking
   - Review functionality
   - Team-based access controls

## Implementation Differences

Besides feature differences, there are some implementation differences to be aware of:

1. **API Design**:
   - GoStructurizr uses a fluent interface with method chaining
   - For example: `system.AddContainer("Web App", "...").WithTags("Web").WithProperties(...)`

2. **Naming Conventions**:
   - GoStructurizr follows Go naming conventions (e.g., `AddContainer`, `WithTags`)
   - Constants like borders and routing strategies use a prefix style (e.g., `BorderSolid`, `RoutingDirect`)

3. **Configuration**:
   - GoStructurizr currently focuses on programmatic configuration
   - The official tools support configuration via properties files and environment variables

## Interoperability

GoStructurizr generates standard Structurizr DSL that can be:
1. Loaded into the Structurizr web service
2. Visualized using the Structurizr CLI
3. Converted to other formats using Structurizr tools

This means you can use GoStructurizr to generate your models and then visualize them with the official Structurizr tooling.

## Example: Same Model in GoStructurizr vs. Official Java Library

**GoStructurizr**:
```go
workspace := gostructurizr.Workspace().WithName("Banking System")
model := workspace.Model()

// Define people and systems
customer := model.AddPerson("Customer", "A bank customer")
internetBanking := model.AddSoftwareSystem("Internet Banking System", "Allows customers to view accounts")

// Create relationships
customer.Uses(internetBanking, "Views account information using")

// Create views
views := workspace.Views()
contextView := views.CreateSystemContextView(internetBanking).
    WithKey("SystemContext").
    WithDescription("System Context diagram")
contextView.AddAllElements()
```

**Official Structurizr Java**:
```java
Workspace workspace = new Workspace("Banking System", "");
Model model = workspace.getModel();

// Define people and systems
Person customer = model.addPerson("Customer", "A bank customer");
SoftwareSystem internetBanking = model.addSoftwareSystem("Internet Banking System", "Allows customers to view accounts");

// Create relationships
customer.uses(internetBanking, "Views account information using");

// Create views
ViewSet views = workspace.getViews();
SystemContextView contextView = views.createSystemContextView(internetBanking, "SystemContext", "System Context diagram");
contextView.addAllElements();
```

## Future Direction

GoStructurizr aims to continue expanding its feature set to match more of the official Structurizr capabilities, with potential focus on:

1. Adding comprehensive documentation capabilities
2. Supporting export to different diagram formats
3. Implementing additional view types
4. Adding support for custom themes and branding
5. Creating a client for the Structurizr API

## Resources

- [Official Structurizr](https://structurizr.com/)
- [Structurizr DSL Grammar](https://github.com/structurizr/dsl/blob/master/docs/language-reference.md)
- [C4 Model](https://c4model.com/)
- [Structurizr Java API](https://github.com/structurizr/java)
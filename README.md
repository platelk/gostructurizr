# GoStructurizr

GoStructurizr is a Go implementation of the [Structurizr DSL](https://structurizr.com/dsl) for creating software architecture diagrams based on the [C4 model](https://c4model.com/).

## Overview

The C4 model provides a way to visualize a software system's architecture at different levels of abstraction:
- **Context**: Shows the system and its relationships with users and other systems
- **Containers**: Reveals the high-level technology choices and how responsibilities are distributed
- **Components**: Zooms in on individual containers to show their components
- **Code**: Details how components are implemented (optional in Structurizr)

GoStructurizr allows you to define these models programmatically in Go and generate Structurizr DSL, which can be visualized using the Structurizr toolchain.

## Installation

```bash
go get github.com/platelk/gostructurizr
```

## Quick Start

Here's a minimal example that creates a system context diagram:

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
    // Create a workspace to contain the model
    workspace := gostructurizr.Workspace().
        WithName("Getting Started").
        WithDesc("This is a model of my software system.")
    
    model := workspace.Model()

    // Define people and systems
    user := model.AddPerson("User", "A user of my software system.")
    softwareSystem := model.AddSoftwareSystem("Software System", "My software system.")
    user.Uses(softwareSystem, "Uses")

    // Create a system context view
    views := workspace.Views()
    contextView := views.CreateSystemContextView(softwareSystem).
        WithKey("SystemContext").
        WithDescription("An example of a System Context diagram.")
    contextView.AddAllElements()
    contextView.AddAllPeople()
    contextView.WithAutoLayout()

    // Add styling
    styles := views.Configuration().Styles()
    styles.AddElementStyle(tags.SoftwareSystem).
        WithBackground("#1168bd").
        WithColor("#ffffff")
    styles.AddElementStyle(tags.Person).
        WithBackground("#08427b").
        WithColor("#ffffff").
        WithShape(shapes.Person)

    // Generate DSL
    var b strings.Builder
    r := renderer.NewDSLRenderer(&b)
    if err := r.Render(workspace); err != nil {
        panic(fmt.Sprintf("renderer didn't succeed: %s", err))
    }
    fmt.Println(b.String())
}
```

## Documentation

For detailed documentation, see the [docs](./docs) directory:

- [Concepts](./docs/concepts): C4 model and core Structurizr concepts
- [Guides](./docs/guides): How-to guides for common tasks
- [Examples](./docs/examples): Comprehensive examples for various use cases

## Features

- ✅ Full C4 model support
- ✅ System context, container, component, and deployment views
- ✅ Styled elements and relationships
- ✅ Dynamic views for interactions
- ✅ Deployment environments
- ✅ Filtered views
- ✅ Custom tags and styling

## License

This project is licensed under the [MIT License](LICENSE).
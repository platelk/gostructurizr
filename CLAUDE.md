# gostructurizr Development Guide

## Commands
- Build: `go build`
- Run tests: `go test ./...`
- Run single test: `go test -run TestName ./package/path`
- Format code: `gofmt -w .`
- Lint: `go vet ./...`

## Code Style Guidelines
- **Imports**: Standard library first, then external packages, alphabetized within groups
- **Formatting**: Follow standard Go format with `gofmt`
- **Types**: Use descriptive structs with pointer fields, implement builder pattern with method chaining
- **Naming**: 
  - PascalCase for exported functions/types
  - camelCase for unexported functions
  - Use descriptive, domain-specific names
- **Error Handling**: Return errors with context using `fmt.Errorf` with `%w` for wrapping
- **Documentation**: Document exported functions, types, and packages with proper godoc comments
- **Testing**: Use `require` package for assertions in tests

## Project Structure
This is a Go implementation of Structurizr DSL for creating software architecture diagrams using the C4 model. See examples in `/example` directory for usage patterns.

## Structurizr DSL Implementation
- **Constants**: String literals for DSL generation are defined as constants in:
  - `/dsl/dsl.go`: Core DSL keywords, border styles, routing styles, documentation properties
  - `/tags/tags.go`: Element type tags (Person, System, Container, Database, etc.)
  - `/routing/routing.go`: Direct and Curved routing constants
- **Builders**: Use a fluent interface pattern with method chaining (With* prefix methods)
- **Renderers**: Components that transform model objects into DSL text representation
- **Views**: Different visualization perspectives (SystemContext, Container, Component, etc.)

## Key Concepts
- **C4 Model**: Hierarchical approach to software architecture (Context, Containers, Components, Code)
- **Element Styles**: Visual styling properties for architecture elements
- **Relationship Styles**: Visual styling for connections between elements
- **Filtered Views**: Views that show/hide elements based on tags
- **Deployment Environments**: Different runtime contexts (Development, Test, Production)
- **Tag-based Styling**: Using tags to apply consistent styling across similar elements
- **Parent-Child Relationships**: Hierarchical structure for deployment and infrastructure nodes

## Examples
- `/example/getting_started/getting_started.go`: Basic example of C4 model
- `/example/advanced_styling/main.go`: Using advanced styling features
- `/example/deployment_environments/main.go`: Deployment across environments
- `/example/filtered_views/main.go`: Tag-based filtered views
- `/example/microservices/microservices.go`: Microservices architecture patterns
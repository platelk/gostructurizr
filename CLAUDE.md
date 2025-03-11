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
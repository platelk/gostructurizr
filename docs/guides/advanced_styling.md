# Advanced Styling in GoStructurizr

Effective visual representation is crucial for communicating architecture. This guide explores advanced styling options in GoStructurizr to create clear, informative, and visually appealing diagrams.

## Why Styling Matters

Proper styling serves several purposes:
- Makes diagrams more readable and professional
- Helps differentiate between element types
- Creates visual hierarchies to guide viewers
- Highlights important architectural characteristics
- Creates consistent visual language across diagrams

## Element Styles

Element styles control the visual appearance of model elements like persons, systems, containers, and components.

### Basic Element Styling

Start with basic styling based on element types:

```go
// Access the styles
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

// Style containers
styles.AddElementStyle(tags.Container).
    WithBackground("#438dd5").
    WithColor("#ffffff")

// Style components
styles.AddElementStyle(tags.Component).
    WithBackground("#85bbf0").
    WithColor("#000000")
```

### Styling Based on Custom Tags

Add custom tags to elements and apply specific styling:

```go
// Add custom tags to elements
webApp := system.AddContainer("Web Application", "Frontend UI", "React")
webApp.AddTags("Web Browser")

api := system.AddContainer("API Gateway", "API Gateway", "Go")
api.AddTags("API")

database := system.AddContainer("Database", "Stores user data", "PostgreSQL")
database.AddTags("Database")

// Apply styles based on custom tags
styles.AddElementStyle("Web Browser").
    WithShape(shapes.WebBrowser).
    WithBackground("#438dd5").
    WithColor("#ffffff")

styles.AddElementStyle("API").
    WithShape(shapes.Hexagon).
    WithBackground("#929000").
    WithColor("#ffffff")

styles.AddElementStyle("Database").
    WithShape(shapes.Cylinder).
    WithBackground("#438dd5").
    WithColor("#ffffff")
```

### Advanced Element Style Properties

GoStructurizr supports numerous styling properties:

```go
styles.AddElementStyle("Important").
    WithBackground("#ff0000").          // Background color
    WithColor("#ffffff").               // Text color
    WithFontSize(24).                   // Font size
    WithBorder(gostructurizr.BorderSolid). // Border style
    WithBorderColor("#000000").         // Border color
    WithBorderWidth(4).                 // Border width
    WithOpacity(90).                    // Opacity percentage
    WithMetadata(true).                 // Show metadata
    WithDescription(true).              // Show description
    WithWidth(400).                     // Element width
    WithHeight(300).                    // Element height
    WithStroke("#000000").              // Stroke color
    WithFontFamily("Arial").            // Font family
    WithIcon("path/to/icon.png").       // Icon URL/path
    WithShape(shapes.RoundedBox)        // Element shape
```

### Available Shapes

GoStructurizr provides a variety of shapes via the `shapes` package:

```go
// Person shapes
shapes.Person              // Person/stick figure

// Container shapes
shapes.Box                 // Standard rectangle
shapes.RoundedBox          // Rectangle with rounded corners
shapes.Circle              // Circle
shapes.Ellipse             // Ellipse
shapes.Hexagon             // Hexagon
shapes.Cylinder            // Database cylinder
shapes.Folder              // Folder
shapes.WebBrowser          // Web browser window
shapes.MobileDevicePortrait  // Mobile device portrait mode
shapes.MobileDeviceLandscape // Mobile device landscape mode
shapes.Component           // Component notation
shapes.Pipe                // Pipe shape
shapes.Robot               // Robot shape for AI/ML
```

## Relationship Styles

Relationship styles control the appearance of connections between elements.

### Basic Relationship Styling

Style relationships based on relationship type or tags:

```go
// Style all relationships
styles.AddRelationshipStyle(tags.Relationship).
    WithThickness(2).
    WithColor("#707070").
    WithDashed(false)

// Style specific relationship types
styles.AddRelationshipStyle("HTTPS").
    WithThickness(2).
    WithColor("#5a9c40").
    WithDashed(false)

styles.AddRelationshipStyle("Asynchronous").
    WithThickness(2).
    WithColor("#d04a43").
    WithDashed(true)
```

### Advanced Relationship Properties

Relationships support various styling options:

```go
styles.AddRelationshipStyle("Critical").
    WithThickness(4).                       // Line thickness
    WithColor("#ff0000").                   // Line color
    WithDashed(true).                       // Dashed line style
    WithRouting(gostructurizr.RoutingDirect). // Routing style
    WithPosition(60).                       // Position of label
    WithOpacity(90).                        // Line opacity
    WithFontSize(24).                       // Label font size
    WithWidth(400).                         // Label width
    WithFollowEdges(true)                   // Whether lines follow edges
```

### Relationship Routing

Control how relationship lines are drawn:

```go
// Direct straight line
styles.AddRelationshipStyle("Direct").
    WithRouting(gostructurizr.RoutingDirect)

// Orthogonal routing (right angles)
styles.AddRelationshipStyle("Orthogonal").
    WithRouting(gostructurizr.RoutingOrthogonal)

// Curved lines
styles.AddRelationshipStyle("Curved").
    WithRouting(gostructurizr.RoutingCurved)
```

## Custom Element Groups

Group related elements together visually:

```go
// Create a group
frontendGroup := model.AddGroup(softwareSystem, "Frontend", "User interface layer")
frontendGroup.Add(webApp)
frontendGroup.Add(mobileApp)

// Style the group
styles.AddElementStyle("Group").
    WithColor("#222222").
    WithBackground("#dddddd").
    WithBorderColor("#222222")
```

## Tag-Based Styling for Element Classification

Use tags to create a classification system:

```go
// Add classification tags
authService.AddTags("Security Critical")
paymentService.AddTags("Business Critical")
loggingService.AddTags("Support Service")

// Style elements by classification
styles.AddElementStyle("Security Critical").
    WithBorderColor("#d04a43").
    WithBorderWidth(4)

styles.AddElementStyle("Business Critical").
    WithBorderColor("#5a9c40").
    WithBorderWidth(4)

styles.AddElementStyle("Support Service").
    WithBorderColor("#808080").
    WithOpacity(70)
```

## Hierarchical Coloring System

Create a color scheme based on architectural layers:

```go
// Define styles by layer
styles.AddElementStyle("UI Layer").
    WithBackground("#1168bd")

styles.AddElementStyle("Application Layer").
    WithBackground("#438dd5")

styles.AddElementStyle("Domain Layer").
    WithBackground("#85bbf0")

styles.AddElementStyle("Infrastructure Layer").
    WithBackground("#b8d2f0")

// Apply tags to elements
webApp.AddTags("UI Layer")
apiGateway.AddTags("Application Layer")
authService.AddTags("Domain Layer")
database.AddTags("Infrastructure Layer")
```

## Using Icons

Add icons to elements for quick visual identification:

```go
// Using built-in icons (if supported)
styles.AddElementStyle("Database").
    WithShape(shapes.Cylinder)

// Using custom icons by URL
styles.AddElementStyle("Kubernetes").
    WithIcon("https://kubernetes.io/images/favicon.png")

// Using multiple icons
styles.AddElementStyle("CloudService").
    WithIcon("https://example.com/cloud.png")
```

## Element Metadata and Properties

Control display of additional information:

```go
// Add properties to elements
database.WithProperties(map[string]string{
    "Database Type": "PostgreSQL",
    "Version": "13",
    "Replicated": "Yes",
    "Backup Schedule": "Daily",
})

// Configure property visibility in styles
styles.AddElementStyle("Database").
    WithMetadata(true)  // Show metadata
```

## View-Specific Style Overrides

Apply specific styles only to certain views:

```go
// Override styles for a specific view
contextView.AddStyleForElement(softwareSystem, 
    gostructurizr.ElementStyle().
        WithBackground("#ff0000").
        WithColor("#ffffff"))

// Override relationship styles for a view
containerView.AddStyleForRelationship(webApp.Uses(api, "Makes API calls to"),
    gostructurizr.RelationshipStyle().
        WithThickness(4).
        WithColor("#ff0000"))
```

## Creating a Custom Styling Theme

Define a consistent theme across all diagrams:

```go
// Create a function to apply a consistent theme
func applyCompanyTheme(views *gostructurizr.Views) {
    styles := views.Configuration().Styles()
    
    // Base styles
    styles.AddElementStyle(tags.Person).
        WithShape(shapes.Person).
        WithBackground("#0078D7").
        WithColor("#ffffff")
    
    styles.AddElementStyle(tags.SoftwareSystem).
        WithBackground("#0078D7").
        WithColor("#ffffff")
    
    styles.AddElementStyle(tags.Container).
        WithBackground("#00BCF2").
        WithColor("#ffffff")
    
    styles.AddElementStyle(tags.Component).
        WithBackground("#50E6FF").
        WithColor("#ffffff")
        
    // Custom technology styles
    styles.AddElementStyle("Microsoft Azure").
        WithBackground("#0072C6")
    
    styles.AddElementStyle("Go").
        WithBackground("#00ADD8")
    
    styles.AddElementStyle("React").
        WithBackground("#61DAFB").
        WithColor("#000000")
        
    // Relationship styles
    styles.AddRelationshipStyle(tags.Relationship).
        WithColor("#707070").
        WithThickness(2)
}

// Apply the theme
applyCompanyTheme(workspace.Views())
```

## Best Practices for Effective Styling

1. **Be consistent**: Use the same color scheme and styling across all diagrams
2. **Use color meaningfully**: Assign colors to represent architectural concepts
3. **Limit your palette**: Don't use too many colors (5-7 is typically enough)
4. **Consider accessibility**: Ensure sufficient contrast for readability
5. **Don't over-style**: Keep diagrams clean and focused
6. **Group related elements**: Use styling to show logical groupings
7. **Highlight what matters**: Use distinctive styling for critical components
8. **Document your styles**: Create a legend or style guide

## Complete Styling Example

See the [advanced styling example](../examples/advanced_styling_example.md) for a complete implementation.

## Resources

- [Structurizr Styling Documentation](https://structurizr.com/help/styling)
- [Color Brewer](https://colorbrewer2.org) - For selecting accessible color palettes
- [C4 Model Standard Colors](https://c4model.com/#Notation)
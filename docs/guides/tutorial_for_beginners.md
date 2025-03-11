# GoStructurizr Tutorial for Beginners

This tutorial walks you through creating your first architecture diagram with GoStructurizr, designed specifically for beginners with little to no experience with the C4 model or Structurizr.

## Prerequisites

- Basic knowledge of Go programming
- Go installed on your system (version 1.16 or newer)
- A code editor of your choice

## What We'll Build

We'll create a simple architecture model for a blogging platform with:
- Users who write and read blog posts
- A web application frontend
- A backend API service
- A database for storing content

We'll visualize this system using multiple views:
1. System context diagram
2. Container diagram
3. Component diagram (for one container)

## Step 1: Set Up Your Project

First, create a new directory for your project and initialize a Go module:

```bash
mkdir blog-architecture
cd blog-architecture
go mod init github.com/yourusername/blog-architecture
```

Then, install the GoStructurizr package:

```bash
go get github.com/platelk/gostructurizr
```

## Step 2: Create Your First Model

Create a file named `main.go` with the following content:

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
	// Create a workspace to contain our model
	workspace := gostructurizr.Workspace().
		WithName("Blog Platform").
		WithDesc("Architecture model for a simple blogging platform")

	model := workspace.Model()

	// Define the people (users of our system)
	author := model.AddPerson("Author", "A person who writes blog posts")
	reader := model.AddPerson("Reader", "A person who reads blog posts")

	// Define the software system
	blogSystem := model.AddSoftwareSystem("Blog Platform", "Allows authors to publish content and readers to read it")
	
	// Define relationships
	author.Uses(blogSystem, "Creates and edits posts using")
	reader.Uses(blogSystem, "Reads content using")

	// Generate the DSL code
	views := workspace.Views()
	contextView := views.CreateSystemContextView(blogSystem).
		WithKey("SystemContext").
		WithDescription("System Context diagram for the Blog Platform")
	contextView.AddAllElements()
	
	// Apply some basic styling
	styles := views.Configuration().Styles()
	styles.AddElementStyle(tags.Person).
		WithBackground("#08427b").
		WithColor("#ffffff").
		WithShape(shapes.Person)
	styles.AddElementStyle(tags.SoftwareSystem).
		WithBackground("#1168bd").
		WithColor("#ffffff")

	// Set up auto-layout
	contextView.WithAutoLayout()

	// Generate DSL
	var b strings.Builder
	r := renderer.NewDSLRenderer(&b)
	if err := r.Render(workspace); err != nil {
		panic(fmt.Sprintf("renderer didn't succeed: %s", err))
	}

	// Output the DSL
	fmt.Println(b.String())
}
```

## Step 3: Run Your Code

Run your code to generate the Structurizr DSL:

```bash
go run main.go > blog-architecture.dsl
```

This will generate a DSL file that describes our simple architecture.

## Step 4: Expand the Model with Containers

Now, let's enhance our model by adding containers within our software system. Update the `main.go` file:

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
	// Create a workspace to contain our model
	workspace := gostructurizr.Workspace().
		WithName("Blog Platform").
		WithDesc("Architecture model for a simple blogging platform")

	model := workspace.Model()

	// Define the people (users of our system)
	author := model.AddPerson("Author", "A person who writes blog posts")
	reader := model.AddPerson("Reader", "A person who reads blog posts")

	// Define the software system
	blogSystem := model.AddSoftwareSystem("Blog Platform", "Allows authors to publish content and readers to read it")
	
	// Define relationships at system level
	author.Uses(blogSystem, "Creates and edits posts using")
	reader.Uses(blogSystem, "Reads content using")

	// Add containers to our system
	webApp := blogSystem.AddContainer(
		"Web Application", 
		"Provides the web interface for authors and readers", 
		"React",
	)
	
	apiService := blogSystem.AddContainer(
		"API Service", 
		"Provides blog functionality via a JSON API", 
		"Go",
	)
	
	database := blogSystem.AddContainer(
		"Database", 
		"Stores blog posts, comments, and user data", 
		"PostgreSQL",
	)
	
	// Define relationships between people and containers
	author.Uses(webApp, "Creates and edits posts using", "HTTPS")
	reader.Uses(webApp, "Reads blog posts using", "HTTPS")
	
	// Define relationships between containers
	webApp.Uses(apiService, "Makes API calls to", "JSON/HTTPS")
	apiService.Uses(database, "Reads from and writes to", "SQL/TCP")

	// Create views
	views := workspace.Views()
	
	// System Context view
	contextView := views.CreateSystemContextView(blogSystem).
		WithKey("SystemContext").
		WithDescription("System Context diagram for the Blog Platform")
	contextView.AddAllElements()
	
	// Container view
	containerView := views.CreateContainerView(blogSystem).
		WithKey("Containers").
		WithDescription("Container diagram for the Blog Platform")
	containerView.AddAllContainers()
	containerView.AddPerson(author)
	containerView.AddPerson(reader)
	
	// Apply styling
	styles := views.Configuration().Styles()
	styles.AddElementStyle(tags.Person).
		WithBackground("#08427b").
		WithColor("#ffffff").
		WithShape(shapes.Person)
	
	styles.AddElementStyle(tags.SoftwareSystem).
		WithBackground("#1168bd").
		WithColor("#ffffff")
	
	styles.AddElementStyle(tags.Container).
		WithBackground("#438dd5").
		WithColor("#ffffff")
	
	// Add specific styling for the database
	database.AddTags("Database")
	styles.AddElementStyle("Database").
		WithShape(shapes.Cylinder)

	// Set up auto-layout for all views
	contextView.WithAutoLayout()
	containerView.WithAutoLayout()

	// Generate DSL
	var b strings.Builder
	r := renderer.NewDSLRenderer(&b)
	if err := r.Render(workspace); err != nil {
		panic(fmt.Sprintf("renderer didn't succeed: %s", err))
	}

	// Output the DSL
	fmt.Println(b.String())
}
```

## Step 5: Add Components to a Container

Now let's add components to the API Service to show its internal architecture:

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
	// Create a workspace to contain our model
	workspace := gostructurizr.Workspace().
		WithName("Blog Platform").
		WithDesc("Architecture model for a simple blogging platform")

	model := workspace.Model()

	// Define the people (users of our system)
	author := model.AddPerson("Author", "A person who writes blog posts")
	reader := model.AddPerson("Reader", "A person who reads blog posts")

	// Define the software system
	blogSystem := model.AddSoftwareSystem("Blog Platform", "Allows authors to publish content and readers to read it")
	
	// Define relationships at system level
	author.Uses(blogSystem, "Creates and edits posts using")
	reader.Uses(blogSystem, "Reads content using")

	// Add containers to our system
	webApp := blogSystem.AddContainer(
		"Web Application", 
		"Provides the web interface for authors and readers", 
		"React",
	)
	
	apiService := blogSystem.AddContainer(
		"API Service", 
		"Provides blog functionality via a JSON API", 
		"Go",
	)
	
	database := blogSystem.AddContainer(
		"Database", 
		"Stores blog posts, comments, and user data", 
		"PostgreSQL",
	)
	
	// Define relationships between people and containers
	author.Uses(webApp, "Creates and edits posts using", "HTTPS")
	reader.Uses(webApp, "Reads blog posts using", "HTTPS")
	
	// Define relationships between containers
	webApp.Uses(apiService, "Makes API calls to", "JSON/HTTPS")
	apiService.Uses(database, "Reads from and writes to", "SQL/TCP")

	// Add components to the API Service
	postController := apiService.AddComponent(
		"Post Controller",
		"Handles requests related to blog posts",
		"Go Handler",
	)
	
	userController := apiService.AddComponent(
		"User Controller",
		"Handles user authentication and profile management",
		"Go Handler",
	)
	
	postService := apiService.AddComponent(
		"Post Service",
		"Contains the business logic for creating and retrieving posts",
		"Go Service",
	)
	
	userService := apiService.AddComponent(
		"User Service",
		"Contains the business logic for user management",
		"Go Service",
	)
	
	postRepository := apiService.AddComponent(
		"Post Repository",
		"Provides access to post data in the database",
		"Go Repository",
	)
	
	userRepository := apiService.AddComponent(
		"User Repository",
		"Provides access to user data in the database",
		"Go Repository",
	)
	
	// Define component relationships
	webApp.Uses(postController, "Makes API calls to", "JSON/HTTPS")
	webApp.Uses(userController, "Makes API calls to", "JSON/HTTPS")
	
	postController.Uses(postService, "Uses")
	userController.Uses(userService, "Uses")
	
	postService.Uses(postRepository, "Uses")
	userService.Uses(userRepository, "Uses")
	
	postRepository.Uses(database, "Reads from and writes to", "SQL/TCP")
	userRepository.Uses(database, "Reads from and writes to", "SQL/TCP")

	// Create views
	views := workspace.Views()
	
	// System Context view
	contextView := views.CreateSystemContextView(blogSystem).
		WithKey("SystemContext").
		WithDescription("System Context diagram for the Blog Platform")
	contextView.AddAllElements()
	
	// Container view
	containerView := views.CreateContainerView(blogSystem).
		WithKey("Containers").
		WithDescription("Container diagram for the Blog Platform")
	containerView.AddAllContainers()
	containerView.AddPerson(author)
	containerView.AddPerson(reader)
	
	// Component view
	componentView := views.CreateComponentView(apiService).
		WithKey("ApiComponents").
		WithDescription("Component diagram for the API Service")
	componentView.AddAllComponents()
	componentView.Add(webApp)
	componentView.Add(database)
	
	// Apply styling
	styles := views.Configuration().Styles()
	styles.AddElementStyle(tags.Person).
		WithBackground("#08427b").
		WithColor("#ffffff").
		WithShape(shapes.Person)
	
	styles.AddElementStyle(tags.SoftwareSystem).
		WithBackground("#1168bd").
		WithColor("#ffffff")
	
	styles.AddElementStyle(tags.Container).
		WithBackground("#438dd5").
		WithColor("#ffffff")
	
	styles.AddElementStyle(tags.Component).
		WithBackground("#85bbf0").
		WithColor("#000000")
	
	// Add specific styling for the database
	database.AddTags("Database")
	styles.AddElementStyle("Database").
		WithShape(shapes.Cylinder)
	
	// Add tags to components by their responsibility
	postController.AddTags("Controller")
	userController.AddTags("Controller")
	postService.AddTags("Service")
	userService.AddTags("Service")
	postRepository.AddTags("Repository")
	userRepository.AddTags("Repository")
	
	// Style components by their responsibility
	styles.AddElementStyle("Controller").
		WithBackground("#85bbf0")
	
	styles.AddElementStyle("Service").
		WithBackground("#a4c2f4")
	
	styles.AddElementStyle("Repository").
		WithBackground("#b6d7a8")

	// Set up auto-layout for all views
	contextView.WithAutoLayout()
	containerView.WithAutoLayout()
	componentView.WithAutoLayout()

	// Generate DSL
	var b strings.Builder
	r := renderer.NewDSLRenderer(&b)
	if err := r.Render(workspace); err != nil {
		panic(fmt.Sprintf("renderer didn't succeed: %s", err))
	}

	// Output the DSL
	fmt.Println(b.String())
}
```

## Step 6: Viewing Your Diagrams

The DSL output from your code can be visualized using the Structurizr tooling. You have several options:

1. **Structurizr Lite** (easiest option):
   - Download [Structurizr Lite](https://structurizr.com/download)
   - Follow the startup instructions
   - Copy your DSL output to the workspace.dsl file
   - View the diagrams in your browser

2. **Online web service**:
   - Sign up for the [Structurizr web service](https://structurizr.com)
   - Create a new workspace
   - Upload your DSL file

3. **Structurizr CLI**:
   - Download the [Structurizr CLI](https://github.com/structurizr/cli/releases)
   - Run: `structurizr-cli export -workspace blog-architecture.dsl -format png`
   - This will export PNG files for each diagram

## Understanding the Code

Let's break down the key parts of our code:

1. **Workspace Creation**
   ```go
   workspace := gostructurizr.Workspace().
       WithName("Blog Platform").
       WithDesc("Architecture model for a simple blogging platform")
   ```
   Creates a top-level container for our architecture model

2. **Adding People**
   ```go
   author := model.AddPerson("Author", "A person who writes blog posts")
   ```
   Adds a person who interacts with our system

3. **Adding Systems**
   ```go
   blogSystem := model.AddSoftwareSystem("Blog Platform", "Allows authors to publish content and readers to read it")
   ```
   Adds the main software system being described

4. **Adding Containers**
   ```go
   webApp := blogSystem.AddContainer(
       "Web Application", 
       "Provides the web interface for authors and readers", 
       "React",
   )
   ```
   Adds a container (application, service, database) inside the system

5. **Adding Components**
   ```go
   postController := apiService.AddComponent(
       "Post Controller",
       "Handles requests related to blog posts",
       "Go Handler",
   )
   ```
   Adds a component (code module) inside a container

6. **Defining Relationships**
   ```go
   author.Uses(webApp, "Creates and edits posts using", "HTTPS")
   ```
   Describes how elements interact with each other

7. **Creating Views**
   ```go
   contextView := views.CreateSystemContextView(blogSystem)
   ```
   Creates a specific view of our model

8. **Styling Elements**
   ```go
   styles.AddElementStyle(tags.Person).
       WithBackground("#08427b").
       WithColor("#ffffff").
       WithShape(shapes.Person)
   ```
   Applies visual styling to make diagrams more readable

## Next Steps

After completing this tutorial, you can explore:

1. **Dynamic Views**: Model how elements interact in specific scenarios
2. **Deployment Views**: Model how your software is deployed to infrastructure
3. **Filtered Views**: Create views that show only specific aspects of your architecture
4. **Custom Styling**: Create more elaborate styling schemes for your diagrams

Refer to the [GoStructurizr documentation](../README.md) for more advanced features and examples.

## Common Pitfalls and Solutions

- **Missing relationships**: Make sure you define relationships between all elements that interact
- **Ambiguous naming**: Use clear, distinct names for different architecture elements
- **Too much detail**: Start with high-level views and gradually add detail where needed
- **Inconsistent styling**: Use a consistent color and shape scheme across all diagrams

## Conclusion

Congratulations! You've created your first architecture model with GoStructurizr. This foundation gives you the skills to start modeling your own software architectures using the C4 model approach.

Remember that good architecture diagrams focus on communication—they help stakeholders understand your system at the appropriate level of detail for their needs.
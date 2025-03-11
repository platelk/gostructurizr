package gostructurizr

// ModelNode represents the top-level software architecture model within a workspace.
// It serves as the container for all architectural elements (systems, people, etc.)
// and their relationships. A model in the C4 approach represents the entire world
// of software systems that are relevant to the architecture being described.
//
// The model maintains collections of people, software systems, relationships,
// enterprise boundaries, and deployment environments. The elements defined in the
// model are visualized through different views (context, container, component).
//
// For more information on the C4 model concept: https://c4model.com/
type ModelNode struct {
	properties      Properties                       // Custom properties for this model
	softwareGroups  []GroupNode[*SoftwareSystemNode] // Logical groupings of software systems
	softwareSystems []*SoftwareSystemNode            // All software systems in the model
	persons         []*PersonNode                    // All people/actors in the model
	uses            []*RelationShipNode              // All relationships between elements
	enterprise      *EnterpriseNode                  // Optional enterprise boundary definition
	deploymentNodes []*DeploymentNodeNode            // All deployment nodes for infrastructure
}

// Model creates a new empty model to represent the software architecture.
// The model acts as the root container for all architecture elements and
// serves as the foundation for creating views of the architecture.
//
// Returns:
//   - A new ModelNode instance with empty collections
//
// Example:
//
//	// Create a new workspace with a model
//	workspace := gostructurizr.Workspace()
//	model := workspace.Model()
//
//	// Add elements to the model
//	user := model.AddPerson("User", "A user of the system")
//	system := model.AddSoftwareSystem("System", "My software system")
func Model() *ModelNode {
	return &ModelNode{}
}

// AddPerson creates and adds a person to the model.
// In the C4 model, a person represents a human user who interacts with 
// your software systems. People are the users or actors that derive value
// from using your software systems (e.g., customers, employees, administrators).
//
// Parameters:
//   - name: The name of the person (e.g., "Administrator", "Customer", "Support Staff")
//   - desc: A description of the person's role and responsibilities
//
// Returns:
//   - A new PersonNode that can be styled and connected via relationships
//
// Example:
//
//	customer := model.AddPerson("Customer", "A user who purchases products")
//	admin := model.AddPerson("Administrator", "A staff member who manages the system")
func (m *ModelNode) AddPerson(name, desc string) *PersonNode {
	p := Person(name, desc)
	m.persons = append(m.persons, p)
	p.model = m
	return p
}

// Persons returns all persons defined in this model.
// This is useful for iterating through all people/actors in the system,
// for example when adding all people to a view.
//
// Returns:
//   - A slice containing all PersonNode instances in the model
//
// Example:
//
//	// Add all people to a view
//	for _, person := range model.Persons() {
//	    contextView.Add(person)
//	}
func (m *ModelNode) Persons() []*PersonNode {
	return m.persons
}

// AddSoftwareSystem creates and adds a software system to the model.
// In the C4 model, a software system is the highest level of abstraction and
// represents a product or application. It's something that delivers value to its users.
//
// Parameters:
//   - name: The name of the software system (e.g., "Online Banking System")
//   - desc: A description of what the software system does and its purpose
//
// Returns:
//   - A new SoftwareSystemNode that can contain containers and be styled and connected
//
// Example:
//
//	bankingSystem := model.AddSoftwareSystem("Banking System", "Handles customer accounts and transactions")
//	crm := model.AddSoftwareSystem("CRM System", "Manages customer relationships")
func (m *ModelNode) AddSoftwareSystem(name, desc string) *SoftwareSystemNode {
	s := SoftwareSystem(name, desc)
	m.softwareSystems = append(m.softwareSystems, s)
	s.model = m
	return s
}

// SoftwareSystems returns all software systems defined in this model.
// This is useful for iterating through all systems in the architecture,
// for example when adding all systems to a view or finding a particular system.
//
// Returns:
//   - A slice containing all SoftwareSystemNode instances in the model
//
// Example:
//
//	// Add all software systems to a view
//	for _, system := range model.SoftwareSystems() {
//	    landscapeView.Add(system)
//	}
func (m *ModelNode) SoftwareSystems() []*SoftwareSystemNode {
	return m.softwareSystems
}

// RelationShip returns all relationships defined in this model.
// Relationships represent the interactions and dependencies between
// elements in the model (people, systems, containers, components).
//
// Returns:
//   - A slice containing all RelationShipNode instances in the model
func (m *ModelNode) RelationShip() []*RelationShipNode {
	return m.uses
}

// addRelationShip creates a new relationship between two elements.
// This is an internal method used to establish a "uses" relationship
// between any two named elements in the model.
//
// Parameters:
//   - from: The source element that initiates the relationship
//   - to: The target element that receives the relationship
//   - desc: A description of how the source uses the target
//
// Returns:
//   - A new RelationShipNode representing the relationship
func (m *ModelNode) addRelationShip(from, to Namer, desc string) *RelationShipNode {
	r := Uses(from, to, desc)
	m.uses = append(m.uses, r)
	return r
}

// SetEnterprise sets the enterprise boundary for this model.
// In the C4 model, an enterprise boundary helps identify which software
// systems are internal to the organization (inside the boundary) versus
// which are external (outside the boundary).
//
// Parameters:
//   - name: The name of the enterprise (typically the organization name)
//
// Returns:
//   - A new EnterpriseNode representing the enterprise boundary
//
// Example:
//
//	enterprise := model.SetEnterprise("ACME Corporation")
//	internalSystem := model.AddSoftwareSystem("Inventory System", "...")
//	enterprise.Add(internalSystem) // Mark as internal to the enterprise
func (m *ModelNode) SetEnterprise(name string) *EnterpriseNode {
	m.enterprise = Enterprise(name)
	m.enterprise.model = m
	return m.enterprise
}

// Enterprise returns the enterprise node for this model.
// This can be nil if no enterprise boundary has been defined.
//
// Returns:
//   - The EnterpriseNode for this model, or nil if not set
//
// Example:
//
//	if enterprise := model.Enterprise(); enterprise != nil {
//	    // Add a new system to the enterprise
//	    system := model.AddSoftwareSystem("HR System", "...")
//	    enterprise.Add(system)
//	}
func (m *ModelNode) Enterprise() *EnterpriseNode {
	return m.enterprise
}

// AddDeploymentEnvironment adds a deployment environment to the model.
// A deployment environment represents a distinct context in which software systems
// are deployed (e.g., Development, Test, Staging, Production).
//
// Parameters:
//   - name: The name of the environment (e.g., "Production", "Development")
//
// Returns:
//   - A DeploymentEnvironment constant that can be used with deployment nodes
//
// Example:
//
//	production := model.AddDeploymentEnvironment("Production")
//	awsCloud := production.AddDeploymentNode("AWS", "Amazon Web Services", "Cloud provider")
func (m *ModelNode) AddDeploymentEnvironment(name string) DeploymentEnvironment {
	// Convert the name to a standard environment if possible
	switch name {
	case "Development":
		return DevelopmentEnvironment
	case "Test":
		return TestEnvironment
	case "Staging":
		return StagingEnvironment
	case "Production":
		return ProductionEnvironment
	default:
		// For custom environments, use the name directly
		return DeploymentEnvironment(name)
	}
}

// AddDeploymentNode adds a deployment node to this model with specified environment.
// Deployment nodes represent the infrastructure elements where software containers
// are deployed, such as physical servers, virtual machines, containers, or cloud services.
//
// Parameters:
//   - name: The name of the deployment node (e.g., "AWS Region", "Kubernetes Cluster")
//   - desc: A description of the deployment node's purpose
//   - technology: The technology used (e.g., "Amazon EC2", "Docker Container")
//   - environment: The deployment environment (e.g., Production, Development)
//
// Returns:
//   - A new DeploymentNodeNode representing the infrastructure
//
// Example:
//
//	prodEnv := model.AddDeploymentEnvironment("Production")
//	awsCloud := model.AddDeploymentNode("AWS", "Amazon Web Services", "Cloud", prodEnv)
func (m *ModelNode) AddDeploymentNode(name, desc, technology string, environment DeploymentEnvironment) *DeploymentNodeNode {
	node := DeploymentNode(name, desc, technology, environment)
	m.deploymentNodes = append(m.deploymentNodes, node)
	node.model = m
	return node
}

// AddDeploymentNodeForEnv adds a deployment node with specified environment.
// This is an alias for AddDeploymentNode for backward compatibility.
//
// Parameters:
//   - name: The name of the deployment node
//   - desc: A description of the deployment node's purpose
//   - technology: The technology used
//   - environment: The deployment environment
//
// Returns:
//   - A new DeploymentNodeNode representing the infrastructure
func (m *ModelNode) AddDeploymentNodeForEnv(name, desc, technology string, environment DeploymentEnvironment) *DeploymentNodeNode {
	return m.AddDeploymentNode(name, desc, technology, environment)
}

// AddDevNode adds a deployment node with Development environment.
// This is a convenience method for adding a deployment node in the Development environment.
//
// Parameters:
//   - name: The name of the deployment node
//   - desc: A description of the deployment node's purpose
//   - technology: The technology used
//
// Returns:
//   - A new DeploymentNodeNode in the Development environment
//
// Example:
//
//	devWorkstation := model.AddDevNode("Developer Workstation", "Local development", "MacBook Pro")
func (m *ModelNode) AddDevNode(name, desc, technology string) *DeploymentNodeNode {
	return m.AddDeploymentNode(name, desc, technology, DevelopmentEnvironment)
}

// AddTestNode adds a deployment node with Test environment.
// This is a convenience method for adding a deployment node in the Test environment.
//
// Parameters:
//   - name: The name of the deployment node
//   - desc: A description of the deployment node's purpose
//   - technology: The technology used
//
// Returns:
//   - A new DeploymentNodeNode in the Test environment
func (m *ModelNode) AddTestNode(name, desc, technology string) *DeploymentNodeNode {
	return m.AddDeploymentNode(name, desc, technology, TestEnvironment)
}

// AddStagingNode adds a deployment node with Staging environment.
// This is a convenience method for adding a deployment node in the Staging environment.
//
// Parameters:
//   - name: The name of the deployment node
//   - desc: A description of the deployment node's purpose
//   - technology: The technology used
//
// Returns:
//   - A new DeploymentNodeNode in the Staging environment
func (m *ModelNode) AddStagingNode(name, desc, technology string) *DeploymentNodeNode {
	return m.AddDeploymentNode(name, desc, technology, StagingEnvironment)
}

// AddProdNode adds a deployment node with Production environment.
// This is a convenience method for adding a deployment node in the Production environment.
//
// Parameters:
//   - name: The name of the deployment node
//   - desc: A description of the deployment node's purpose
//   - technology: The technology used
//
// Returns:
//   - A new DeploymentNodeNode in the Production environment
//
// Example:
//
//	awsCloud := model.AddProdNode("AWS", "Amazon Web Services", "Cloud provider")
func (m *ModelNode) AddProdNode(name, desc, technology string) *DeploymentNodeNode {
	return m.AddDeploymentNode(name, desc, technology, ProductionEnvironment)
}

// DeploymentNodes returns all deployment nodes in this model.
// This is useful for working with all infrastructure nodes regardless of environment.
//
// Returns:
//   - A slice containing all DeploymentNodeNode instances in the model
func (m *ModelNode) DeploymentNodes() []*DeploymentNodeNode {
	return m.deploymentNodes
}

// FindDeploymentNodesForEnvironment returns all deployment nodes for a specific environment.
// This allows filtering deployment nodes by environment (e.g., only get Production nodes).
//
// Parameters:
//   - environment: The deployment environment to filter by
//
// Returns:
//   - A slice containing DeploymentNodeNode instances for the specified environment
//
// Example:
//
//	prodNodes := model.FindDeploymentNodesForEnvironment(ProductionEnvironment)
//	for _, node := range prodNodes {
//	    deploymentView.Add(node)
//	}
func (m *ModelNode) FindDeploymentNodesForEnvironment(environment DeploymentEnvironment) []*DeploymentNodeNode {
	var result []*DeploymentNodeNode
	for _, node := range m.deploymentNodes {
		if node.environment == environment {
			result = append(result, node)
		}
	}
	return result
}

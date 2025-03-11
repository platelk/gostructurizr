package gostructurizr

type ModelNode struct {
	properties      Properties
	softwareGroups  []GroupNode[*SoftwareSystemNode]
	softwareSystems []*SoftwareSystemNode
	persons         []*PersonNode
	uses            []*RelationShipNode
	enterprise      *EnterpriseNode
	deploymentNodes []*DeploymentNodeNode
}

func Model() *ModelNode {
	return &ModelNode{}
}

func (m *ModelNode) AddPerson(name, desc string) *PersonNode {
	p := Person(name, desc)
	m.persons = append(m.persons, p)
	p.model = m
	return p
}

func (m *ModelNode) Persons() []*PersonNode {
	return m.persons
}

func (m *ModelNode) AddSoftwareSystem(name, desc string) *SoftwareSystemNode {
	s := SoftwareSystem(name, desc)
	m.softwareSystems = append(m.softwareSystems, s)
	s.model = m
	return s
}

func (m *ModelNode) SoftwareSystems() []*SoftwareSystemNode {
	return m.softwareSystems
}

func (m *ModelNode) RelationShip() []*RelationShipNode {
	return m.uses
}

func (m *ModelNode) addRelationShip(from, to Namer, desc string) *RelationShipNode {
	r := Uses(from, to, desc)
	m.uses = append(m.uses, r)
	return r
}

// SetEnterprise sets the enterprise for this model
func (m *ModelNode) SetEnterprise(name string) *EnterpriseNode {
	m.enterprise = Enterprise(name)
	m.enterprise.model = m
	return m.enterprise
}

// Enterprise returns the enterprise node for this model
func (m *ModelNode) Enterprise() *EnterpriseNode {
	return m.enterprise
}

// AddDeploymentNode adds a deployment node to this model with specified environment
func (m *ModelNode) AddDeploymentNode(name, desc, technology string, environment DeploymentEnvironment) *DeploymentNodeNode {
	node := DeploymentNode(name, desc, technology, environment)
	m.deploymentNodes = append(m.deploymentNodes, node)
	node.model = m
	return node
}

// AddDeploymentNodeForEnv adds a deployment node with specified environment
func (m *ModelNode) AddDeploymentNodeForEnv(name, desc, technology string, environment DeploymentEnvironment) *DeploymentNodeNode {
	return m.AddDeploymentNode(name, desc, technology, environment)
}

// AddDevNode adds a deployment node with Development environment
func (m *ModelNode) AddDevNode(name, desc, technology string) *DeploymentNodeNode {
	return m.AddDeploymentNode(name, desc, technology, DevelopmentEnvironment)
}

// AddTestNode adds a deployment node with Test environment
func (m *ModelNode) AddTestNode(name, desc, technology string) *DeploymentNodeNode {
	return m.AddDeploymentNode(name, desc, technology, TestEnvironment)
}

// AddStagingNode adds a deployment node with Staging environment
func (m *ModelNode) AddStagingNode(name, desc, technology string) *DeploymentNodeNode {
	return m.AddDeploymentNode(name, desc, technology, StagingEnvironment)
}

// AddProdNode adds a deployment node with Production environment
func (m *ModelNode) AddProdNode(name, desc, technology string) *DeploymentNodeNode {
	return m.AddDeploymentNode(name, desc, technology, ProductionEnvironment)
}

// DeploymentNodes returns all deployment nodes in this model
func (m *ModelNode) DeploymentNodes() []*DeploymentNodeNode {
	return m.deploymentNodes
}

// FindDeploymentNodesForEnvironment returns all deployment nodes for a specific environment
func (m *ModelNode) FindDeploymentNodesForEnvironment(environment DeploymentEnvironment) []*DeploymentNodeNode {
	var result []*DeploymentNodeNode
	for _, node := range m.deploymentNodes {
		if node.environment == environment {
			result = append(result, node)
		}
	}
	return result
}

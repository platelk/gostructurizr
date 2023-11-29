package main

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/renderer"
	"github.com/platelk/gostructurizr/routing"
	"github.com/platelk/gostructurizr/shapes"
	"github.com/platelk/gostructurizr/tags"
	"strings"
)

const (
	microserviceTag = "Microservices"
	messageBusTag   = "Message Bus"
	datastoreTag    = "Database"
)

func main() {
	workspace := gostructurizr.Workspace().WithName("Microservices example").WithDesc("An example of a microservices architecture, which includes asynchronous and parallel behaviour.")
	model := workspace.Model()

	mySoftwareSystem := model.AddSoftwareSystem("Customer Information System", "Stores information ")
	customer := model.AddPerson("Customer", "A customer")
	customerApplication := mySoftwareSystem.AddContainer("Customer Application").WithDesc("Allows customers to manage their profile.").WithTechnology("Angular")

	customerService := mySoftwareSystem.AddContainer("Customer Service").WithDesc("The point of access for customer information.").WithTechnology("Java and Spring Boot")
	customerService.WithTag(microserviceTag)
	customerDatabase := mySoftwareSystem.AddContainer("Customer Database").WithDesc("Stores customer information.").WithTechnology("Oracle 12c")
	customerDatabase.WithTag(datastoreTag)

	reportingService := mySoftwareSystem.AddContainer("Reporting Service").WithDesc("Creates normalised data for reporting purposes.").WithTechnology("Ruby")
	reportingService.WithTag(microserviceTag)
	reportingDatabase := mySoftwareSystem.AddContainer("Reporting Database").WithDesc("Stores a normalised version of all business data for ad hoc reporting purposes.").WithTechnology("MySQL")
	reportingDatabase.WithTag(datastoreTag)

	auditService := mySoftwareSystem.AddContainer("Audit Service").WithDesc("Provides organisation-wide auditing facilities.").WithTechnology("C# .NET")
	auditService.WithTag(microserviceTag)
	auditStore := mySoftwareSystem.AddContainer("Audit Database").WithDesc("Stores information about events that have happened.").WithTechnology("Event Store")
	auditStore.WithTag(datastoreTag)

	messageBus := mySoftwareSystem.AddContainer("Message Bus").WithDesc("Transport for business events.").WithTechnology("RabbitMQ")
	messageBus.WithTag(messageBusTag)

	customer.Uses(customerApplication, "Uses")

	customerApplication.Uses(customerService, "Updates customer information using").WithTechnology("JSON/HTTPS").WithInteractionStyle(gostructurizr.Synchronous)
	customerService.Uses(messageBus, "Sends customer update events to").WithInteractionStyle(gostructurizr.Asynchronous)
	customerService.Uses(customerDatabase, "Stores data in").WithTechnology("JDBC").WithInteractionStyle(gostructurizr.Synchronous)
	customerService.Uses(customerApplication, "Sends events to").WithTechnology("WebSocket").WithInteractionStyle(gostructurizr.Asynchronous)
	messageBus.Uses(reportingService, "Sends customer update events to").WithInteractionStyle(gostructurizr.Asynchronous)
	messageBus.Uses(auditService, "Sends customer update events to").WithInteractionStyle(gostructurizr.Asynchronous)
	reportingService.Uses(reportingDatabase, "Stores data in").WithInteractionStyle(gostructurizr.Synchronous)
	auditService.Uses(auditStore, "Stores events in").WithInteractionStyle(gostructurizr.Synchronous)

	views := workspace.Views()
	containerView := views.CreateContainerView(mySoftwareSystem)
	containerView.AddAllElements()
	containerView.WithAutoLayout().WithKey("Containers_All")
	containerView.WithInclude(gostructurizr.On(mySoftwareSystem).WithAfferent(true).WithEfferent(true))
	dynamicView := views.CreateDynamicView(mySoftwareSystem).
		WithKey("CustomerUpdateEvent").
		WithDescription("This diagram shows what happens when a customer updates their details.")
	dynamicView.Add(customer, customerApplication)
	dynamicView.Add(customerApplication, customerService)

	dynamicView.Add(customerService, customerDatabase)
	dynamicView.Add(customerService, messageBus)

	dynamicView.StartParallelSequence()
	dynamicView.Add(messageBus, reportingService)
	dynamicView.Add(reportingService, reportingDatabase)
	dynamicView.EndParallelSequence()

	dynamicView.StartParallelSequence()
	dynamicView.Add(messageBus, auditService)
	dynamicView.Add(auditService, auditStore)
	dynamicView.EndParallelSequence()

	dynamicView.StartParallelSequence()
	dynamicView.Add(customerService, customerApplication, "Confirms update to")
	dynamicView.EndParallelSequence()

	styles := views.Configuration().Styles()
	styles.AddElementStyle(tags.Element).WithColor("#000000")
	styles.AddElementStyle(tags.Person).WithBackground("#ffbf00").WithShape(shapes.Person)
	styles.AddElementStyle(tags.Container).WithBackground("#facc2E")
	styles.AddElementStyle(messageBusTag).WithWidth(1600).WithShape(shapes.Pipe)
	styles.AddElementStyle(microserviceTag).WithShape(shapes.Hexagon)
	styles.AddElementStyle(datastoreTag).WithBackground("#f5da81").WithShape(shapes.Cylinder)
	styles.AddRelationshipStyle(tags.RelationShip).WithRouting(routing.Orthogonal)

	styles.AddRelationshipStyle(tags.Asynchronous).WithDash(true)
	styles.AddRelationshipStyle(tags.Synchronous).WithDash(false)

	var b strings.Builder
	r := renderer.NewDSLRenderer(&b)
	if err := r.Render(workspace); err != nil {
		panic(fmt.Sprintf("renderer didn't succeed: %s", err))
	}
	fmt.Println(b.String())
}

/*
   public static void main(String[] args) throws Exception {
       Workspace workspace = new Workspace("Microservices example", "An example of a microservices architecture, which includes asynchronous and parallel behaviour.");
       Model model = workspace.getModel();

       SoftwareSystem mySoftwareSystem = model.AddSoftwareSystem("Customer Information System", "Stores information ");
       Person customer = model.AddPerson("Customer", "A customer");
       Container customerApplication = mySoftwareSystem.AddContainer("Customer Application", "Allows customers to manage their profile.", "Angular");

       Container customerService = mySoftwareSystem.AddContainer("Customer Service", "The point of access for customer information.", "Java and Spring Boot");
       customerService.AddTags(MICROSERVICE_TAG);
       Container customerDatabase = mySoftwareSystem.AddContainer("Customer Database", "Stores customer information.", "Oracle 12c");
       customerDatabase.AddTags(DATASTORE_TAG);

       Container reportingService = mySoftwareSystem.AddContainer("Reporting Service", "Creates normalised data for reporting purposes.", "Ruby");
       reportingService.AddTags(MICROSERVICE_TAG);
       Container reportingDatabase = mySoftwareSystem.AddContainer("Reporting Database", "Stores a normalised version of all business data for ad hoc reporting purposes.", "MySQL");
       reportingDatabase.AddTags(DATASTORE_TAG);

       Container auditService = mySoftwareSystem.AddContainer("Audit Service", "Provides organisation-wide auditing facilities.", "C# .NET");
       auditService.AddTags(MICROSERVICE_TAG);
       Container auditStore = mySoftwareSystem.AddContainer("Audit Store", "Stores information about events that have happened.", "Event Store");
       auditStore.AddTags(DATASTORE_TAG);

       Container messageBus = mySoftwareSystem.AddContainer("Message Bus", "Transport for business events.", "RabbitMQ");
       messageBus.AddTags(MESSAGE_BUS_TAG);

       customer.uses(customerApplication, "Uses");
       customerApplication.uses(customerService, "Updates customer information using", "JSON/HTTPS", InteractionStyle.Synchronous);
       customerService.uses(messageBus, "Sends customer update events to", "", InteractionStyle.Asynchronous);
       customerService.uses(customerDatabase, "Stores data in", "JDBC", InteractionStyle.Synchronous);
       customerService.uses(customerApplication, "Sends events to", "WebSocket", InteractionStyle.Asynchronous);
       messageBus.uses(reportingService, "Sends customer update events to", "", InteractionStyle.Asynchronous);
       messageBus.uses(auditService, "Sends customer update events to", "", InteractionStyle.Asynchronous);
       reportingService.uses(reportingDatabase, "Stores data in", "", InteractionStyle.Synchronous);
       auditService.uses(auditStore, "Stores events in", "", InteractionStyle.Synchronous);

       ViewSet views = workspace.getViews();

       ContainerView containerView = views.createContainerView(mySoftwareSystem, "Containers", null);
       containerView.AddAllElements();

       DynamicView dynamicView = views.createDynamicView(mySoftwareSystem, "CustomerUpdateEvent", "This diagram shows what happens when a customer updates their details.");
       dynamicView.Add(customer, customerApplication);
       dynamicView.Add(customerApplication, customerService);

       dynamicView.Add(customerService, customerDatabase);
       dynamicView.Add(customerService, messageBus);

       dynamicView.startParallelSequence();
       dynamicView.Add(messageBus, reportingService);
       dynamicView.Add(reportingService, reportingDatabase);
       dynamicView.endParallelSequence();

       dynamicView.startParallelSequence();
       dynamicView.Add(messageBus, auditService);
       dynamicView.Add(auditService, auditStore);
       dynamicView.endParallelSequence();

       dynamicView.startParallelSequence();
       dynamicView.Add(customerService, "Confirms update to", customerApplication);
       dynamicView.endParallelSequence();

       Styles styles = views.getConfiguration().getStyles();
       styles.AddElementStyle(Tags.ELEMENT).color("#000000");
       styles.AddElementStyle(Tags.PERSON).background("#ffbf00").shape(Shape.Person);
       styles.AddElementStyle(Tags.CONTAINER).background("#facc2E");
       styles.AddElementStyle(MESSAGE_BUS_TAG).width(1600).shape(Shape.Pipe);
       styles.AddElementStyle(MICROSERVICE_TAG).shape(Shape.Hexagon);
       styles.AddElementStyle(DATASTORE_TAG).background("#f5da81").shape(Shape.Cylinder);
       styles.AddRelationshipStyle(Tags.RELATIONSHIP).routing(Routing.Orthogonal);

       styles.AddRelationshipStyle(Tags.ASYNCHRONOUS).dashed(true);
       styles.AddRelationshipStyle(Tags.SYNCHRONOUS).dashed(false);

       StructurizrClient client = new StructurizrClient("key", "secret");
       client.putWorkspace(4241, workspace);
   }
*/

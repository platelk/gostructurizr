package main

import "github.com/platelk/gostructurizr"

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

}

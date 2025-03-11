workspace "Deployment Environments Example" "An example of deployment environments with infrastructure and container instances" {
    model {
        customer = person "Customer" "A customer of the online store"
        webStore = softwareSystem "Web Store" "Online retail system" {
            tags "WebStore"
            webApplication = container "Web Application" "The main web application" "Java and Spring Boot"
            database = container "Database" "Customer and order information" "MySQL"
            cache = container "Cache" "Caches product information" "Redis"
        }

        customer -> webApplication "Visits website using"
        webApplication -> database "Reads from and writes to"
        webApplication -> cache "Reads from and writes to"
        webServer -> rdsMySql "Connects to"
        webServer -> elastiCache "Reads from and writes to"
        webServer -> rdsMySql "Connects to"
        webServer -> elastiCache "Reads from and writes to"
        elasticLoadBalancer -> webServer "Routes requests to"
        elasticLoadBalancer -> webServer "Routes requests to"
    }
    views {
        container webStore "Containers" "Container view for Web Store" {
            include *
            autoLayout
        }
        deploymentView {
            softwareSystem webStore
            environment "Development"
            key "DevelopmentDeployment"
            description "Development deployment"
            autoLayout
            include developerLaptop
        }
        deploymentView {
            softwareSystem webStore
            environment "Production"
            key "ProductionDeployment"
            description "Production deployment"
            autoLayout
            include amazonWebServices
        }
        styles {
            element "Person" {
                shape person
                background #08427B
                color #ffffff
            }
            element "Container" {
                background #438DD5
                color #ffffff
            }
            element "Deployment Node" {
                background #ffffff
                color #000000
            }
            element "Infrastructure Node" {
                shape Ellipse
                background #C5E6FF
                color #000000
            }
            element "Container Instance" {
                background #438DD5
                color #ffffff
                border 2
            }
        }
    }
}

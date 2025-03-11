workspace "Advanced Styling Example" "This is an example of advanced element and relationship styling in Structurizr" {
    model {
        user = person "User" "A user of the system" "External"
        webApplication = softwareSystem "Web Application" "The main web application" {
            tags "WebApp"
        }
        database = softwareSystem "Database" "The primary database" {
            tags "Database"
        }
        cache = softwareSystem "Cache" "Redis cache" {
            tags "Cache"
        }
        messaging = softwareSystem "Messaging" "Kafka messaging platform" {
            tags "Queue"
        }
        api = softwareSystem "API" "External REST API" {
            tags "API"
        }

        user -> webApplication "Uses"
        webApplication -> database "Reads from and writes to"
        webApplication -> cache "Reads from and writes to"
        webApplication -> messaging "Publishes events to"
        webApplication -> api "Makes API calls to"
    }
    views {
        systemContext webApplication "SystemContext" "System Context diagram"{
            include *
            autoLayout
        }
        styles {
            element "Person" {
                shape person
                background #08427B
                color #ffffff
                fontSize 24
                border 20
                fontFamily "Arial"
                shadow true
            }
            element "Software system" {
                shape RoundedBox
                background #1168BD
                color #ffffff
            }
            element "Database" {
                shape cylinder
                background #1168BD
                color #ffffff
            }
            element "WebApp" {
                background #62A420
                borderStyle "Dashed"
                strokeWidth 2
            }
            element "API" {
                shape hexagon
                background #85BBF0
                border 4
            }
            element "Cache" {
                background #D4A017
                rotation 15
            }
            element "Queue" {
                shape pipe
                background #E62D2D
            }
            element "External" {
                background #999999
                fontStyle "italic"
            }
            element "relationship" {
                background #707070
            }
            element "Sync" {
                color "#289CE1"
                thickness 2
                style "Solid"
                fontSize 12
                fontColor "#289CE1"
                routing "Direct"
            }
            element "Async" {
                color "#E62D2D"
                thickness 2
                style "Dashed"
                fontSize 12
                fontColor "#E62D2D"
                routing "Curved"
            }
            element "Cache" {
                color "#D4A017"
                thickness 2
                style "Dotted"
                fontColor "#D4A017"
                destinationTerminator "Arrow"
            }
            element "Database" {
                color "#1168BD"
                thickness 2
                routing "Orthogonal"
            }
        }
    }
}

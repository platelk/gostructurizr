workspace "Filtered Views Example" "This is an example of filtered views in Structurizr" {
    model {
        customerA = person "Customer A" "A premium customer" "Customer,Premium"
        customerB = person "Customer B" "A regular customer" "Customer,Regular"
        administrator = person "Administrator" "System administrator" "Staff,Admin"
        supportStaff = person "Support Staff" "Customer support" "Staff,Support"
        webApplication = softwareSystem "Web Application" "The main web application" {
            tags "Internal, WebApp"
        }
        customerDatabase = softwareSystem "Customer Database" "Stores customer information" {
            tags "Internal, Database, Critical"
        }
        reportingSystem = softwareSystem "Reporting System" "Generates business reports" {
            tags "Internal, Reporting"
        }
        adminPortal = softwareSystem "Admin Portal" "Admin management interface" {
            tags "Internal, AdminTool"
        }
        paymentProvider = softwareSystem "Payment Provider" "Processes payments" {
            tags "External, Payment"
        }
        emailSystem = softwareSystem "Email System" "Sends emails to customers" {
            tags "External, Communication"
        }
        monitoringSystem = softwareSystem "Monitoring System" "Monitors system health" {
            tags "External, Monitoring"
        }

        customerA -> webApplication "Uses"
        customerB -> webApplication "Uses"
        administrator -> adminPortal "Manages system using"
        supportStaff -> adminPortal "Views customer info using"
        webApplication -> customerDatabase "Reads from and writes to"
        webApplication -> paymentProvider "Makes payments using"
        webApplication -> emailSystem "Sends emails using"
        adminPortal -> customerDatabase "Reads from"
        adminPortal -> reportingSystem "Generates reports using"
        reportingSystem -> customerDatabase "Reads from"
        monitoringSystem -> webApplication "Monitors"
        monitoringSystem -> customerDatabase "Monitors"
    }
    views {
        systemContext webApplication "SystemContext" "The system context diagram"{
            include *
            autoLayout
        }
        filteredView {
            baseView "SystemContext"
            title "Shows customer interaction"
            key "CustomerView"
            Include Tag "Customer"
            Include Tag "WebApp"
            Include Tag "Payment"
            Exclude Tag "Admin"
            Exclude Tag "Monitoring"
            Exclude Tag "Reporting"
            autoLayout
        }
        filteredView {
            baseView "SystemContext"
            title "Shows admin capabilities"
            key "AdminView"
            Include Tag "Admin"
            Include Tag "Staff"
            Include Tag "AdminTool"
            Include Tag "Reporting"
            Exclude Tag "Customer"
            Exclude Tag "Payment"
            autoLayout
        }
        filteredView {
            baseView "SystemContext"
            title "Shows critical systems only"
            key "CriticalView"
            Include Tag "Critical"
            Include Tag "WebApp"
            autoLayout
        }
        filteredView {
            baseView "SystemContext"
            title "Shows only external integrations"
            key "ExternalView"
            Include Tag "External"
            Include Tag "WebApp"
            autoLayout
        }
        styles {
            element "Person" {
                shape person
                background #08427B
                color #ffffff
            }
            element "Software system" {
                shape RoundedBox
                background #1168BD
                color #ffffff
            }
            element "Customer" {
                background #3498DB
            }
            element "Premium" {
                background #2E86C1
                fontStyle "bold"
            }
            element "Regular" {
                background #5DADE2
            }
            element "Staff" {
                background #16A085
            }
            element "Database" {
                shape cylinder
                background #9B59B6
            }
            element "AdminTool" {
                background #2C3E50
            }
            element "Critical" {
                background #E74C3C
                border 4
                borderStyle "Solid"
            }
            element "External" {
                background #95A5A6
                fontStyle "italic"
            }
        }
    }
}

package main

import ("fmt"
        "github.com/ncw/swift"
)

func main() {

    fmt.Println("hello there")

    // Create a connection
    c := swift.Connection {
        UserName: "username",
        ApiKey: "password",
        AuthUrl: "auth-url",
        Domain: "domain", // Name of the domain 
        AuthVersion: 3,
        TenantId: "tenantId",  // Id of the tenant/project         
    }
	
    // Authenticate
    err := c.Authenticate()
    if err != nil {
        panic(err)
    }

    // List all containers
    containers, err := c.ContainerNamesAll(nil)
    if err != nil {
        panic(err)
    }
    fmt.Println("Containers: ")
    fmt.Println(containers)

    // Create new container
    err1 := c.ContainerCreate("testContainer", nil)
    if err1 != nil {
      panic(err1)
    }
    fmt.Println("Container created")

    // List objects
    object, headers, err := c.Object(CONTAINER, OBJECT)
    if err != nil {
      panic(err)
    }
    fmt.Println("Object Info: ")
    fmt.Println(object.Name)
    fmt.Println(headers.ObjectMetadata())

    // Account info
    info, headers, err := c.Account()
    if err != nil {
        panic(err)
    }
    fmt.Println("Account Info: ")
    fmt.Println(info)
    fmt.Println(headers)   

}

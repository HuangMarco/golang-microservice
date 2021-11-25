package main

import (
	"fmt"

// how to organize the package, repository etc: https://go.dev/doc/code

	"github.com/golang-microservice/goblog/accountservice/service"
)
var appName = "accountservice"

func main(){
	
  fmt.Printf("Starting %v\n",appName)

	service.StartWebServer("8080")
}

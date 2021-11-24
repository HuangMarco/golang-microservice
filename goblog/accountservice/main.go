package main

import (
	"fmt"

	"github.com/golang-microservice/goblog/accountservice/service"
)
var appName = "accountservice"

func main(){
	
  fmt.Printf("Starting %v\n",appName)

	service.StartWebServer("8080")
}

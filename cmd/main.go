package main

import (
	"log"

	"github.com/akshay0074700747/book_store/configurations"
	dependencyinjection "github.com/akshay0074700747/book_store/dependency_injection"
)

func main() {

	//loading the configurations from the env
	configs, err := configurations.LoadConfigurationss()
	if err != nil {
		log.Fatalf("fatal error %s , exiting...", err.Error())
	}

	//this application in built in clean architecture
	//this function is initialising all the dependencies
	controller := dependencyinjection.InjectDependencies(configs)

	//starting up the server in the specified port
	controller.Start(configs.Port)
}

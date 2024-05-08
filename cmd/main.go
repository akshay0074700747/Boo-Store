package main

import (
	"log"

	"github.com/akshay0074700747/book_store/configurations"
	dependencyinjection "github.com/akshay0074700747/book_store/dependency_injection"
)

func main() {

	configs, err := configurations.LoadConfigurationss()
	if err != nil {
		log.Fatalf("fatal error %s happend, exiting...", err.Error())
	}

	controller := dependencyinjection.InjectDependencies(configs)
	controller.Start(configs.Port)
}

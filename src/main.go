// API REST EXAMPLE
//
// This is an example. This API serves article objects.
//
//     Schemes: http
//     Host: localhost:8000
//     Version: 0.1.0
//     basePath: /
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"golang-backend-service/src/infrastructure"
	"golang-backend-service/src/logging"
	"golang-backend-service/src/repository"
	"golang-backend-service/src/usecases"
	"log"
	"os"
)

func main() {

	//Get configuration
	certFile := os.Getenv("CERT_FILE") // example: CERT_FILE=./keys/mycert.pem
	keyFile := os.Getenv("KEY_FILE")   // example: KEY_FILE=./keys/myKey.key
	addrhttp := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	addrhttps := fmt.Sprintf(":%s", os.Getenv("HTTPS_PORT"))

	//Initialize the logger
	//Log Levels "ERROR", "WARNING", "INFO", "DEBUG"
	logLevel := os.Getenv("LOG_LEVEL")
	err := logging.Init(logLevel)
	if err != nil {
		log.Fatal("Error on initializing logger: ", err)
	}

	logging.Debug.Println("initialize configuration successful")

	//Create repository
	ar := repository.NewInMemoryArticlesRepository()
	logging.Debug.Println("Repository layer created")

	//Create use case
	uc := usecases.NewArticleUseCase(ar)
	logging.Debug.Println("UseCase layer created")

	//Create infrastructure, router and handler
	httpRouter := infrastructure.NewHttpRouter(addrhttp, addrhttps, certFile, keyFile)
	infrastructure.NewArticleHandler(&httpRouter.Router, uc)
	logging.Debug.Println("http infrastructure created")
	logging.Debug.Println("Swagger on https://localhost:8443/swaggerui/")

	//Start http(s) listener
	errs := httpRouter.HandleRequests()

	// This will run forever until channel receives error
	select {
	case err := <-errs:
		logging.Error.Printf("Could not start serving service due to (error: %s)", err)
	}

}

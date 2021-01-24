package main

import (
	"log"

	"github.com/codershangfeng/vote-service/app/internal/api/restapi"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewVoteServiceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080
	api.GetHealthHandler = operations.GetHealthHandlerFunc(
		func(ghp operations.GetHealthParams) middleware.Responder {
			return operations.NewGetHealthOK()
		})

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

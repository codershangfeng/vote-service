package main

import (
	"log"

	"github.com/codershangfeng/vote-service/app/internal/api/restapi"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
	"github.com/codershangfeng/vote-service/app/internal/handler"
	"github.com/go-openapi/loads"
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

	// TODO: Move to server module later
	api.ProbeGetHealthHandler = probe.GetHealthHandlerFunc(
		handler.GetHealthHandler,
	)
	api.VoteGetVoteByIDHandler = vote.GetVoteByIDHandlerFunc(
		handler.GetVoteByIDHandler,
	)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

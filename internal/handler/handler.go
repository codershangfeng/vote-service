package handler

import (
	"github.com/codershangfeng/vote-service/app/internal/api/models"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
	"github.com/go-openapi/runtime/middleware"
)

// GetHealthHandler defines retrieving health status of GET request agaist probe
func GetHealthHandler(ghp probe.GetHealthParams) middleware.Responder {
	return probe.NewGetHealthOK()
}

// GetVoteByIDHandler defines retrieving vote item by ID of GET request against vote
func GetVoteByIDHandler(gvbip vote.GetVoteByIDParams) middleware.Responder {
	return vote.NewGetVoteByIDOK().WithPayload(&models.Vote{ID: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"})
}

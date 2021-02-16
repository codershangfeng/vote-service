package handler

import (
	"fmt"
	"math/rand"

	"github.com/codershangfeng/vote-service/app/internal/api/models"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
	"github.com/go-openapi/runtime/middleware"
)

var db map[int64]models.Vote

func init() {
	db = map[int64]models.Vote{
		1: {ID: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"},
		2: {ID: 2, Options: []string{"Noodle", "Dumpling"}, Topic: "Which food do you prefer?"},
		3: {ID: 3, Options: []string{"Basketball", "Billiards"}, Topic: "Which sports do you prefer?"},
		4: {ID: 4, Options: []string{"Beethoven", "Mozart"}, Topic: "Which artist do you prefer?"},
	}
}

// GetHealthHandler defines retrieving health status of GET request agaist probe
func GetHealthHandler(ghp probe.GetHealthParams) middleware.Responder {
	return probe.NewGetHealthOK()
}

// GetVoteByIDHandler defines retrieving vote item by ID of GET request against vote
func GetVoteByIDHandler(gvbip vote.GetVoteByIDParams) middleware.Responder {
	v, ok := db[gvbip.VoteID]

	// Bug!
	r := rand.Intn(2)

	if r == 0 {
		fmt.Errorf("Bug triggered when r = {%d}!", r)
		return vote.NewGetVoteByIDBadRequest()
	}

	if !ok {
		return vote.NewGetVoteByIDNotFound()
	}
	return vote.NewGetVoteByIDOK().WithPayload(&v)
}

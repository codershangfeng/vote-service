package handler

import (
	"github.com/codershangfeng/vote-service/app/internal/api/models"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/votes"
	"github.com/codershangfeng/vote-service/app/internal/persistence"
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

var repository persistence.Repository

func InitRepository(repo persistence.Repository) {
	repository = repo
}

// GetHealthHandler defines retrieving health status of GET request agaist probe
func GetHealthHandler(ghp probe.GetHealthParams) middleware.Responder {
	return probe.NewGetHealthOK()
}

// GetVoteByIDHandler defines retrieving vote item by ID of GET request against vote
func GetVoteByIDHandler(gvbip vote.GetVoteByIDParams) middleware.Responder {
	entity := repository.GetVoteEntity(gvbip.VoteID)
	if entity == nil {
		return vote.NewGetVoteByIDNotFound()
	}
	v := models.Vote{ID: entity.ID, Options: entity.Options, Topic: entity.Topic}
	return vote.NewGetVoteByIDOK().WithPayload(&v)
}

func GetVotes(gvp votes.GetVotesParams) middleware.Responder {
	entities := repository.GetVoteEntities()

	// vs := make(models.Votes, len(db))

	// for k, v := range db {
	// 	value := v
	// 	vs[k-1] = &value
	// }
	vs := make(models.Votes, 0, len(entities))

	for _, e := range entities {
		vs = append(vs, &models.Vote{ID: e.ID, Options: e.Options, Topic: e.Topic})
	}

	return votes.NewGetVotesOK().WithPayload(vs)
}

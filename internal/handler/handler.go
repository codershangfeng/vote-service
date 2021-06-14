package handler

import (
	"github.com/codershangfeng/vote-service/app/internal/api/models"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/votes"
	"github.com/codershangfeng/vote-service/app/internal/persistence"
	"github.com/go-openapi/runtime/middleware"
)

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

	vs := make(models.Votes, 0, len(entities))

	for _, e := range entities {
		vs = append(vs, &models.Vote{ID: e.ID, Options: e.Options, Topic: e.Topic})
	}

	return votes.NewGetVotesOK().WithPayload(vs)
}

func SaveVote(svp votes.SaveVoteParams) middleware.Responder {
	v := svp.Vote
	repository.SaveVoteEntity(persistence.VoteEntity{ID: v.ID, Options: v.Options, Topic: v.Topic})
	return votes.NewSaveVoteCreated()
}

func DeleteVote(dvbip vote.DeleteVoteByIDParams) middleware.Responder {
	if err := repository.DeleteVoteEntity(dvbip.VoteID); err != nil {
		return vote.NewDeleteVoteByIDNotFound()
	}

	return vote.NewDeleteVoteByIDOK()
}

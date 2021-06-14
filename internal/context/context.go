package context

import (
	"github.com/codershangfeng/vote-service/app/internal/api/restapi"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/votes"
	"github.com/codershangfeng/vote-service/app/internal/handler"
	"github.com/codershangfeng/vote-service/app/internal/persistence"
	"github.com/go-openapi/loads"
)

// AppContext defines context for a vote service server.
type AppContext struct {
	Port int
}

// NewServer creates a new server based on its context
func (ctx *AppContext) NewServer(api *operations.VoteServiceAPI) *restapi.Server {
	server := restapi.NewServer(api)

	server.Port = ctx.Port

	return server
}

// NewAPIHandler returns the api handler of server
func NewAPIHandler(repo persistence.Repository) (*operations.VoteServiceAPI, error) {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")

	if err != nil {
		return nil, err
	}

	api := operations.NewVoteServiceAPI(swaggerSpec)

	if repo != nil {
		handler.InitRepository(repo)
	} else {
		handler.InitRepository(persistence.NewRepository())
	}

	api.ProbeGetHealthHandler = probe.GetHealthHandlerFunc(
		handler.GetHealthHandler,
	)

	api.VoteGetVoteByIDHandler = vote.GetVoteByIDHandlerFunc(
		handler.GetVoteByIDHandler,
	)

	api.VotesGetVotesHandler = votes.GetVotesHandlerFunc(
		handler.GetVotes,
	)

	api.VotesSaveVoteHandler = votes.SaveVoteHandlerFunc(
		handler.SaveVote,
	)

	return api, nil
}

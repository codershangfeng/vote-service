package context

import (
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
	"github.com/go-openapi/loads"
)

// App is an entry of Swagger Appliation with configuration(context)
type App struct {
	SwaggerSpec            *loads.Document
	ProbeGetHealthHandler  probe.GetHealthHandler
	VoteGetVoteByIDHandler vote.GetVoteByIDHandler
}

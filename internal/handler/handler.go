package handler

import (
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/go-openapi/runtime/middleware"
)

func BuildGetHealthHandlerFunc() probe.GetHealthHandlerFunc {
	return getHealthHandler
}

func getHealthHandler(ghp probe.GetHealthParams) middleware.Responder {
	return probe.NewGetHealthOK()
}

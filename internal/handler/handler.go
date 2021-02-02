package handler

import (
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/go-openapi/runtime/middleware"
)

// GetHealthHandler defines handling flow of GET request against health endpoint
func GetHealthHandler(ghp probe.GetHealthParams) middleware.Responder {
	return probe.NewGetHealthOK()
}

package handler

import (
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func BuildGetHealthHandlerFunc() operations.GetHealthHandlerFunc {
	return getHealthHandler
}

func getHealthHandler(ghp operations.GetHealthParams) middleware.Responder {
	return operations.NewGetHealthOK()
}

// +build integration

package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/codershangfeng/vote-service/app/internal/api/restapi"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
	"github.com/codershangfeng/vote-service/app/internal/handler"
	"github.com/go-openapi/loads"
)

func getAPI() (*operations.VoteServiceAPI, error) {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	api := operations.NewVoteServiceAPI(swaggerSpec)
	return api, nil
}

func getAPIHandler() (http.Handler, error) {
	api, err := getAPI()
	if err != nil {
		return nil, err
	}

	// TODO: Reuse server's ConfigureAPI() later
	api.ProbeGetHealthHandler = probe.GetHealthHandlerFunc(
		handler.GetHealthHandler,
	)
	api.VoteGetVoteByIDHandler = vote.GetVoteByIDHandlerFunc(
		handler.GetVoteByIDHandler,
	)
	// h := setupGlobalMiddleware(api)
	h := setupGlobalMiddleware(api.Serve(setupMiddlewares))
	err = api.Validate()
	if err != nil {
		return nil, err
	}
	return h, nil
}

var ts *httptest.Server

func TestMain(m *testing.M) {
	h, err := getAPIHandler()
	if err != nil {
		log.Fatal("Get api handler failed due to ", err)
		os.Exit(1)
	}
	ts = httptest.NewServer(h)
	defer ts.Close()

	os.Exit(m.Run())
}

func TestGetHealthAPI(t *testing.T) {
	res, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Errorf("Failed to send request to get health endpoint: %s", err)
	}

	got := res.StatusCode
	expect := 200
	if got != expect {
		t.Errorf("Expect get health endpoint return %d, but got %d", expect, got)
	}
}

func TestGetVoteByIDAPI(t *testing.T) {
	res, err := http.Get(ts.URL + "/vote/1")

	if err != nil {
		t.Errorf("Failed to send request to get vote by id endpoint: %s", err)
	}

	got := res.StatusCode
	expect := 200

	if got != expect {
		t.Errorf("Expect get vote by id return %d, but got %d", expect, got)
	}
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

// Another implementation of setupGlobalMiddleware ->
//
// func setupGlobalMiddleware(handler http.Handler) http.Handler {
//     return uiMiddleware(handler)
// }

// func uiMiddleware(handler http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Shortcut helpers for swagger-ui
//         if r.URL.Path == "/swagger-ui" || r.URL.Path == "/api/help" {
//             http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
//             return
//         }
//         // Serving ./swagger-ui/
//         if strings.Index(r.URL.Path, "/swagger-ui/") == 0 {
//             http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("swagger-ui"))).ServeHTTP(w, r)
//             return
//         }
//         handler.ServeHTTP(w, r)
//     })
// }

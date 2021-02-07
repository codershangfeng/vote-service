// +build integration

package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations"
	"github.com/codershangfeng/vote-service/app/internal/context"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	api, err := context.NewAPIHandler()

	if err != nil {
		log.Fatal("Error when create api handler: ", err)
		os.Exit(1)
	}

	ts = httptest.NewServer(configureTestAPI(api))
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

func configureTestAPI(api *operations.VoteServiceAPI) http.Handler {
	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
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


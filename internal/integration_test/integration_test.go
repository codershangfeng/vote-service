package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations"
	"github.com/codershangfeng/vote-service/app/internal/context"
	"github.com/codershangfeng/vote-service/app/internal/persistence"
	"github.com/stretchr/testify/assert"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	mockRepo := persistence.NewRepository()
	mockRepo.SaveVoteEntity(persistence.VoteEntity{
		Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?",
	})
	mockRepo.SaveVoteEntity(persistence.VoteEntity{
		Options: []string{"Noodle", "Dumpling"}, Topic: "Which food do you prefer?",
	})

	api, err := context.NewAPIHandler(mockRepo)

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

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "{\"options\":[\"Innocence\",\"Firework\"],\"topic\":\"Which song do you prefer?\",\"vid\":1}\n", string(body))
}

func TestGetVotesAPI(t *testing.T) {
	res, err := http.Get(ts.URL + "/votes")

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "[{\"options\":[\"Innocence\",\"Firework\"],\"topic\":\"Which song do you prefer?\",\"vid\":1},{\"options\":[\"Noodle\",\"Dumpling\"],\"topic\":\"Which food do you prefer?\",\"vid\":2}]\n", string(body))
}

func TestSaveVoteAPI(t *testing.T) {
	res, err := http.Post(ts.URL+"/votes", "application/json", strings.NewReader("{\"options\":[\"Innocence\",\"Firework\"],\"topic\":\"Which song do you prefer?\"}"))

	assert.Nil(t, err)
	assert.Equal(t, 201, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "{\"options\":[\"Innocence\",\"Firework\"],\"topic\":\"Which song do you prefer?\",\"vid\":3}\n", string(body))
}

func TestUpdateVoteAPI(t *testing.T) {
	client := &http.Client{Timeout: 3 * time.Second}
	req, err := http.NewRequest(http.MethodPut, ts.URL+"/vote/1", strings.NewReader("{\"options\":[\"Innocence\",\"Firework\"],\"topic\":\"Which song do you prefer?\"}"))
	req.Header.Set("Content-Type", "application/json")
	assert.Nil(t, err)

	res, err := client.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
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

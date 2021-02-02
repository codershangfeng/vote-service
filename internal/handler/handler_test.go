// +build unit

package handler

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

	"github.com/codershangfeng/vote-service/app/internal/api/models"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
)

func TestGetHealthHandler(t *testing.T) {
	got := GetHealthHandler(probe.NewGetHealthParams())

	if got.(*probe.GetHealthOK) == probe.NewGetHealthOK() {
		t.Errorf("Expected NewGetHealthOK response, but got: {%T}", got)
	}
}

func TestGetVoteByIDHandler(t *testing.T) {
	params := vote.NewGetVoteByIDParams()
	params.VoteID = 1
	got := GetVoteByIDHandler(params)
	expect := vote.NewGetVoteByIDOK().WithPayload(&models.Vote{ID: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"})

	if !reflect.DeepEqual(got.(*vote.GetVoteByIDOK), expect) {
		t.Errorf("Expected get vote by ID: \n{%s}, but got: \n{%s}\n", marshal(expect), marshal(got))
	}
}

func marshal(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("Failed to marshal object {%T} to json\n", v)
	}
	return string(bytes)
}

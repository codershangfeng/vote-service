// +build unit

package handler

import (
	"encoding/json"
	"log"
	"math"
	"reflect"
	"testing"

	"github.com/codershangfeng/vote-service/app/internal/api/models"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/votes"
)

func TestGetHealthHandler(t *testing.T) {
	got := GetHealthHandler(probe.NewGetHealthParams())

	if got.(*probe.GetHealthOK) == probe.NewGetHealthOK() {
		t.Errorf("Expected NewGetHealthOK response, but got: {%T}", got)
	}
}

func TestShouldReturnOKWhenVoteCanBeFound(t *testing.T) {
	votes := [...]models.Vote{
		{ID: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"},
		{ID: 2, Options: []string{"Noodle", "Dumpling"}, Topic: "Which food do you prefer?"},
		{ID: 3, Options: []string{"Basketball", "Billiards"}, Topic: "Which sports do you prefer?"},
		{ID: 4, Options: []string{"Beethoven", "Mozart"}, Topic: "Which artist do you prefer?"},
	}
	for _, v := range votes {
		params := vote.NewGetVoteByIDParams()
		params.VoteID = v.ID
		got := GetVoteByIDHandler(params)

		expect := vote.NewGetVoteByIDOK().WithPayload(&v)

		if !reflect.DeepEqual(got.(*vote.GetVoteByIDOK), expect) {
			t.Errorf("Expected get vote by ID: \n{%s}, but got: \n{%s}\n", marshal(expect), marshal(got))
		}
	}
}

func TestShouldReturnNotFoundWhenVoteCanNotFound(t *testing.T) {
	params := vote.NewGetVoteByIDParams()
	params.VoteID = math.MaxInt64
	got := GetVoteByIDHandler(params)

	expect := vote.NewGetVoteByIDNotFound()

	if !reflect.DeepEqual(got.(*vote.GetVoteByIDNotFound), expect) {
		t.Errorf("Expected get vote by ID: \n{%s}, but got: \n{%s}\n", marshal(expect), marshal(got))
	}
}

func TestShouldReturnOKWhenVotesCanBeFound(t *testing.T) {
	vs := models.Votes{
		{ID: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"},
		{ID: 2, Options: []string{"Noodle", "Dumpling"}, Topic: "Which food do you prefer?"},
		{ID: 3, Options: []string{"Basketball", "Billiards"}, Topic: "Which sports do you prefer?"},
		{ID: 4, Options: []string{"Beethoven", "Mozart"}, Topic: "Which artist do you prefer?"},
	}

	params := votes.NewGetVotesParams()
	got := GetVotes(params)

	expect := votes.NewGetVotesOK().WithPayload(vs)

	if !reflect.DeepEqual(got.(*votes.GetVotesOK), expect) {
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

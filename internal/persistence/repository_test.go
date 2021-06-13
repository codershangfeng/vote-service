package persistence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockVote = VoteEntity{ID: 1, Options: []string{"apple"}, Topic: "What's your favorite fruit?"}

func TestShouldSaveVoteSuccessfully(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := Repository{database: db}

	repo.SaveVote(mockVote)

	assert.NotEmpty(t, db)
	assert.Equal(t, 1, len(db))
	assert.Equal(t, int64(1), db[1].ID)
	assert.Equal(t, []string{"apple"}, db[1].Options)
	assert.Equal(t, "What's your favorite fruit?", db[1].Topic)
}

func TestShouldGetVoteByIDSuccessfully(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := Repository{database: db}
	db[mockVote.ID] = mockVote

	got := repo.GetVote(int64(1))

	assert.NotNil(t, got)
	assert.Equal(t, int64(1), got.ID)
	assert.Equal(t, []string{"apple"}, got.Options)
	assert.Equal(t, "What's your favorite fruit?", got.Topic)
}

func TestShouldReturnNilWhenGetVoteByIDAndVoteDoesNotExist(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := Repository{database: db}

	got := repo.GetVote(int64(1))

	assert.Nil(t, got)
}

func TestShouldNotImpactOriginalEntityWhenModifyItemReturnedFromGetVote(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := Repository{database: db}
	db[mockVote.ID] = mockVote

	got := repo.GetVote(int64(1))
	origin := db[mockVote.ID]
	origin.ID = 3
	origin.Options = []string{"banana"}
	origin.Topic = "What's yours?"

	assert.NotNil(t, got)
	assert.Equal(t, int64(1), got.ID)
	assert.Equal(t, []string{"apple"}, db[1].Options)
	assert.Equal(t, "What's your favorite fruit?", db[1].Topic)
}

func TestShouldGetVotesSuccessfully(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := Repository{database: db}
	db[int64(1)] = VoteEntity{ID: 1, Options: []string{"apple"}, Topic: "What's your favorite fruit?"}
	db[int64(2)] = VoteEntity{ID: 2, Options: []string{"basketball"}, Topic: "What's your favorite sports?"}

	got := repo.GetVotes()

	assert.NotEmpty(t, got)
	assert.Len(t, got, 2)
	assert.Equal(t, VoteEntity{ID: 1, Options: []string{"apple"}, Topic: "What's your favorite fruit?"}, got[0])
	assert.Equal(t, VoteEntity{ID: 2, Options: []string{"basketball"}, Topic: "What's your favorite sports?"}, got[1])
}

func TestShouldReturnEmptyListWhenGetVotesAndVoteDoesNotExist(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := Repository{database: db}

	got := repo.GetVotes()

	assert.Nil(t, got)
}

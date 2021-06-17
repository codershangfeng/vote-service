package persistence

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestShouldSaveVoteSuccessfully(t *testing.T) {
	mockIDGenerator := new(MockIDGenerator)
	db := make(map[int64]VoteEntity)
	mockIDGenerator.On("IncAndGet").Return(int64(123)).Once()
	repo := RepositoryImpl{database: db, idGenerator: mockIDGenerator}

	repo.SaveVoteEntity(VoteEntity{Options: []string{"apple"}, Topic: "What's your favorite fruit?"})

	assert.NotEmpty(t, db)
	assert.Equal(t, 1, len(db))
	assert.Equal(t, int64(123), db[123].ID)
	assert.Equal(t, []string{"apple"}, db[123].Options)
	assert.Equal(t, "What's your favorite fruit?", db[123].Topic)
}

func TestShouldGetVoteByIDSuccessfully(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := RepositoryImpl{database: db}
	db[int64(1)] = VoteEntity{ID: 1, Options: []string{"apple"}, Topic: "What's your favorite fruit?"}

	got := repo.GetVoteEntity(int64(1))

	assert.NotNil(t, got)
	assert.Equal(t, int64(1), got.ID)
	assert.Equal(t, []string{"apple"}, got.Options)
	assert.Equal(t, "What's your favorite fruit?", got.Topic)
}

func TestShouldReturnNilWhenGetVoteByIDAndVoteDoesNotExist(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := RepositoryImpl{database: db}

	got := repo.GetVoteEntity(int64(1))

	assert.Nil(t, got)
}

func TestShouldNotImpactOriginalEntityWhenModifyItemReturnedFromGetVote(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := RepositoryImpl{database: db}
	db[int64(1)] = VoteEntity{ID: 1, Options: []string{"apple"}, Topic: "What's your favorite fruit?"}

	got := repo.GetVoteEntity(int64(1))
	origin := db[int64(1)]
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
	repo := RepositoryImpl{database: db}
	db[int64(1)] = VoteEntity{ID: 1, Options: []string{"apple"}, Topic: "What's your favorite fruit?"}
	db[int64(2)] = VoteEntity{ID: 2, Options: []string{"basketball"}, Topic: "What's your favorite sports?"}

	got := repo.GetVoteEntities()

	assert.NotEmpty(t, got)
	assert.Len(t, got, 2)
	assert.Equal(t, VoteEntity{ID: 1, Options: []string{"apple"}, Topic: "What's your favorite fruit?"}, got[0])
	assert.Equal(t, VoteEntity{ID: 2, Options: []string{"basketball"}, Topic: "What's your favorite sports?"}, got[1])
}

func TestShouldDeleteVoteSuccessfully(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := RepositoryImpl{database: db}
	db[int64(1)] = VoteEntity{ID: 1, Options: []string{"apple"}, Topic: "What's your favorite fruit?"}

	err := repo.DeleteVoteEntity(int64(1))

	assert.Nil(t, err)
	_, ok := db[int64(1)]
	assert.False(t, ok)
}

func TestShouldReturnErrorWhenDeleteVoteFailed(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := RepositoryImpl{database: db}

	err := repo.DeleteVoteEntity(int64(1))

	assert.NotNil(t, err)
	assert.Equal(t, "vote with ID[1] does not exist", err.Error())
}

func TestShouldReturnEmptyListWhenGetVotesAndVoteDoesNotExist(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := RepositoryImpl{database: db}

	got := repo.GetVoteEntities()

	assert.Nil(t, got)
}

func TestShouldUpdateSuccuessfully(t *testing.T) {
	db := make(map[int64]VoteEntity)
	repo := RepositoryImpl{database: db}
	db[int64(1)] = VoteEntity{ID: 1, Options: []string{"apple"}, Topic: "What's your favorite fruit?"}

	err := repo.UpdateVoteEntity(VoteEntity{ID: 1, Options: []string{"football"}, Topic: "Which sport do you prefer?"})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(db))
	assert.Equal(t, VoteEntity{ID: 1, Options: []string{"football"}, Topic: "Which sport do you prefer?"}, db[1])
}

type MockIDGenerator struct {
	mock.Mock
}

func (o *MockIDGenerator) IncAndGet() int64 {
	args := o.Called()
	return args.Get(0).(int64)
}

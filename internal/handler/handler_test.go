package handler

import (
	"errors"
	"testing"

	"github.com/codershangfeng/vote-service/app/internal/api/models"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/vote"
	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/votes"
	"github.com/codershangfeng/vote-service/app/internal/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetHealthHandler(t *testing.T) {
	got := GetHealthHandler(probe.NewGetHealthParams())

	assert.Equal(t, probe.NewGetHealthOK(), got.(*probe.GetHealthOK))
}

func TestShouldReturnOKWhenVoteCanBeFound(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("GetVoteEntity", int64(1)).Return(&persistence.VoteEntity{ID: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"}).Once()
	InitRepository(mockRepo)

	params := vote.NewGetVoteByIDParams()
	params.VoteID = 1
	got := GetVoteByIDHandler(params)

	assert.Equal(t, vote.NewGetVoteByIDOK().WithPayload(&models.Vote{ID: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"}), got.(*vote.GetVoteByIDOK))
	mockRepo.AssertExpectations(t)
}

func TestShouldReturnNotFoundWhenVoteCanNotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("GetVoteEntity", int64(1)).Return(nil).Once()
	InitRepository(mockRepo)

	params := vote.NewGetVoteByIDParams()
	params.VoteID = 1
	got := GetVoteByIDHandler(params)

	assert.Equal(t, vote.NewGetVoteByIDNotFound(), got.(*vote.GetVoteByIDNotFound))
	mockRepo.AssertExpectations(t)
}

func TestShouldReturnOKWhenVotesCanBeFound(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("GetVoteEntities").Return([]persistence.VoteEntity{
		{ID: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"},
		{ID: 2, Options: []string{"Noodle", "Dumpling"}, Topic: "Which food do you prefer?"},
	}).Once()
	InitRepository(mockRepo)

	got := GetVotes(votes.NewGetVotesParams())

	assert.Equal(t, votes.NewGetVotesOK().WithPayload([]*models.Vote{
		{ID: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"},
		{ID: 2, Options: []string{"Noodle", "Dumpling"}, Topic: "Which food do you prefer?"},
	}), got.(*votes.GetVotesOK))
	mockRepo.AssertExpectations(t)
}

func TestShouldReturnCreatedWhenVoteCanBeSavedSuccessfully(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("SaveVoteEntity", persistence.VoteEntity{ID: int64(1), Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"}).Once()
	InitRepository(mockRepo)

	params := votes.NewSaveVoteParams()
	params.Vote = &models.Vote{ID: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"}
	got := SaveVote(params)

	assert.Equal(t, votes.NewSaveVoteCreated(), got.(*votes.SaveVoteCreated))
	mockRepo.AssertExpectations(t)
}

func TestShouldReturnOKWhenVoteCanBeDeletedSuccessfully(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("DeleteVoteEntity", int64(1)).Return(nil).Once()
	InitRepository(mockRepo)

	params := vote.NewDeleteVoteByIDParams()
	params.VoteID = 1
	got := DeleteVote(params)

	assert.Equal(t, vote.NewDeleteVoteByIDOK(), got.(*vote.DeleteVoteByIDOK))
	mockRepo.AssertExpectations(t)
}

func TestShouldReturnNotFoundWhenDeleteGotError(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("DeleteVoteEntity", int64(1)).Return(errors.New("some error")).Once()
	InitRepository(mockRepo)

	params := vote.NewDeleteVoteByIDParams()
	params.VoteID = 1
	got := DeleteVote(params)

	assert.Equal(t, vote.NewDeleteVoteByIDNotFound(), got.(*vote.DeleteVoteByIDNotFound))
	mockRepo.AssertExpectations(t)
}

type MockRepo struct {
	mock.Mock
}

func (o *MockRepo) GetVoteEntity(id int64) *persistence.VoteEntity {
	args := o.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*persistence.VoteEntity)
}

func (o *MockRepo) GetVoteEntities() []persistence.VoteEntity {
	args := o.Called()
	return args.Get(0).([]persistence.VoteEntity)
}

func (o *MockRepo) SaveVoteEntity(v persistence.VoteEntity) {
	o.Called(v)
}

func (o *MockRepo) DeleteVoteEntity(id int64) error {
	args := o.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

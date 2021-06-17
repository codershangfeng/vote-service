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

	assert.Equal(t, vote.NewGetVoteByIDOK().WithPayload(&models.VoteOutgoing{Vid: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"}), got.(*vote.GetVoteByIDOK))
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

	assert.Equal(t, votes.NewGetVotesOK().WithPayload([]*models.VoteOutgoing{
		{Vid: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"},
		{Vid: 2, Options: []string{"Noodle", "Dumpling"}, Topic: "Which food do you prefer?"},
	}), got.(*votes.GetVotesOK))
	mockRepo.AssertExpectations(t)
}

func TestShouldReturnCreatedWhenVoteCanBeSavedSuccessfully(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("SaveVoteEntity", persistence.VoteEntity{Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"}).Return(persistence.VoteEntity{ID: int64(1), Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"}).Once()
	InitRepository(mockRepo)

	params := votes.NewSaveVoteParams()
	topic := "Which song do you prefer?"
	params.Vote = &models.VoteIncoming{Options: []string{"Innocence", "Firework"}, Topic: &topic}
	got := SaveVote(params)

	assert.Equal(t, votes.NewSaveVoteCreated().WithPayload(&models.VoteOutgoing{Vid: 1, Options: []string{"Innocence", "Firework"}, Topic: "Which song do you prefer?"}), got.(*votes.SaveVoteCreated))
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

func TestShouldReturnOKWhenModifedExistedVote(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("UpdateVoteEntity", persistence.VoteEntity{ID: 1, Topic: "which sports do you perfer?", Options: []string{"basketball"}}).Return(nil).Once()
	InitRepository(mockRepo)

	params := vote.NewUpdateVoteByIDParams()
	params.VoteID = 1
	topic := "which sports do you perfer?"
	params.Vote = &models.VoteIncoming{Topic: &topic, Options: []string{"basketball"}}
	got := UpdateVote(params)

	assert.Equal(t, vote.NewUpdateVoteByIDOK(), got.(*vote.UpdateVoteByIDOK))
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

func (o *MockRepo) SaveVoteEntity(v persistence.VoteEntity) persistence.VoteEntity {
	args := o.Called(v)
	return args.Get(0).(persistence.VoteEntity)
}

func (o *MockRepo) DeleteVoteEntity(id int64) error {
	args := o.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (o *MockRepo) UpdateVoteEntity(v persistence.VoteEntity) error {
	args := o.Called(v)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

package tests

import (
	"context"
	"findsafe/backend/models/mocks/repository"
	"findsafe/backend/models/models"
	"findsafe/backend/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTeamService(t *testing.T) {
	// Creating a new instance of the MockTeamRepository
	mockRepo := new(mocks.MockTeamRepository)
	// Creating a new instance of the TeamService with the mocked repository
	service := &services.TeamService{TeamRepository: mockRepo}

	// Test UUID
	testUID := uuid.New()

	t.Run("Test Get", func(t *testing.T) {
		team := &models.Team{ID: testUID}
		mockRepo.On("FindTeamByID", mock.Anything, testUID).Return(team, nil)

		result, err := service.Get(context.Background(), testUID)

		assert.NoError(t, err)
		assert.Equal(t, team, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test UpdateTeam", func(t *testing.T) {
		team := &models.Team{ID: testUID}
		mockRepo.On("UpdateTeam", mock.Anything, testUID, team).Return(nil)

		err := service.Update(context.Background(), testUID, team)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test DeleteTeam", func(t *testing.T) {
		mockRepo.On("DeleteTeam", mock.Anything, testUID).Return(nil)

		err := service.Delete(context.Background(), testUID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetAllinSearch", func(t *testing.T) {
		team := &models.Team{ID: testUID}
		mockRepo.On("FindBySearch", mock.Anything, testUID).Return([]*models.Team{team}, nil)

		result, err := service.GetAllinSearch(context.Background(), testUID)

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, team, result[0])
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetAllWithoutSortie", func(t *testing.T) {
		team := &models.Team{ID: testUID}
		mockRepo.On("FindByEmptySortie", mock.Anything).Return([]*models.Team{team}, nil)

		result, err := service.GetAllWithoutSortie(context.Background())

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, team, result[0])
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetBySortie", func(t *testing.T) {
		team := &models.Team{ID: testUID}
		mockRepo.On("FindBySortie", mock.Anything, testUID).Return(team, nil)

		result, err := service.GetBySortie(context.Background(), testUID)

		assert.NoError(t, err)
		assert.Equal(t, team, result)
		mockRepo.AssertExpectations(t)
	})
}

package tests

import (
	"context"
	mocks2 "findsafe/account/models/mocks/repository"
	"findsafe/account/models/models"
	"findsafe/account/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestResourceService(t *testing.T) {
	// Mock repository instance
	mockRepo := new(mocks2.MockResourceRepository)

	// Creating a new instance of ResourceService with the mocked repository
	service := &services.ResourceService{ResourceRepository: mockRepo}

	// Test UUIDs
	testUID := uuid.New()
	testOwnerID := uuid.New()
	testTeamID := uuid.New()

	t.Run("Test Get", func(t *testing.T) {
		// Setting up the mock response
		resource := &models.Resource{ID: testUID}
		mockRepo.On("FindResourceByID", mock.Anything, testUID).Return(resource, nil)

		// Calling the service method
		result, err := service.Get(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, resource, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test UpdateResource", func(t *testing.T) {
		resource := &models.Resource{ID: testUID}
		mockRepo.On("UpdateResource", mock.Anything, testUID, resource).Return(nil)

		// Call UpdateResource method
		err := service.Update(context.Background(), testUID, resource)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test DeleteResource", func(t *testing.T) {
		mockRepo.On("DeleteResource", mock.Anything, testUID).Return(nil)

		// Call DeleteResource method
		err := service.Delete(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetByOwnerID", func(t *testing.T) {
		resource := &models.Resource{ID: testUID}
		mockRepo.On("FindByOwnerID", mock.Anything, &testOwnerID).Return([]*models.Resource{resource}, nil)

		// Calling GetByOwnerID method
		result, err := service.GetByOwnerID(context.Background(), testOwnerID)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)             // Ensure that exactly 1 result is returned
		assert.Equal(t, resource, result[0]) // Ensure the returned resource is the same
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetByTeamID", func(t *testing.T) {
		resource := &models.Resource{ID: testUID}
		mockRepo.On("FindByTeamID", mock.Anything, &testTeamID).Return([]*models.Resource{resource}, nil)

		// Calling GetByTeamID method
		result, err := service.GetByTeamID(context.Background(), testTeamID)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)             // Ensure that exactly 1 result is returned
		assert.Equal(t, resource, result[0]) // Ensure the returned resource is the same
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetByAvailable", func(t *testing.T) {
		resource := &models.Resource{ID: testUID}
		mockRepo.On("FindAvailable", mock.Anything, &testUID).Return([]*models.Resource{resource}, nil)

		// Calling GetByAvailable method
		result, err := service.GetByAvailable(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)             // Ensure that exactly 1 result is returned
		assert.Equal(t, resource, result[0]) // Ensure the returned resource is the same
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetUnreturned", func(t *testing.T) {
		resource := &models.Resource{ID: testUID}
		mockRepo.On("FindUnreturned", mock.Anything, &testUID).Return([]*models.Resource{resource}, nil)

		// Calling GetUnreturned method
		result, err := service.GetUnreturned(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)             // Ensure that exactly 1 result is returned
		assert.Equal(t, resource, result[0]) // Ensure the returned resource is the same
		mockRepo.AssertExpectations(t)
	})
}

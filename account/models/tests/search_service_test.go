package tests

import (
	"context"
	"findsafe/account/models"
	mocks "findsafe/account/models/mocks/repository"
	"findsafe/account/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSearchService(t *testing.T) {
	// Mock repository instance
	mockRepo := new(mocks.MockSearchRepository)

	// Creating a new instance of SearchService with the mocked repository
	service := &services.SearchService{SearchRepository: mockRepo}

	// Test UUIDs
	testUID := uuid.New()
	testOrgID := uuid.New()

	t.Run("Test Get", func(t *testing.T) {
		// Setting up the mock response
		search := &models.Searches{ID: testUID}
		mockRepo.On("FindByID", mock.Anything, testUID).Return(search, nil)

		// Calling the service method
		result, err := service.Get(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, search, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetAll", func(t *testing.T) {
		// Setting up the mock response
		search := &models.Searches{ID: testUID}
		mockRepo.On("FindAll", mock.Anything).Return([]*models.Searches{search}, nil)

		// Calling the service method
		result, err := service.GetAll(context.Background())

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)           // Ensure exactly 1 result is returned
		assert.Equal(t, search, result[0]) // Ensure the returned search is the same
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetAllBySubject", func(t *testing.T) {
		// Setting up the mock response
		search := &models.Searches{ID: testUID}
		mockRepo.On("FindAllBySubject", mock.Anything, testUID).Return(search, nil)

		// Calling the service method
		result, err := service.GetAllBySubject(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, search, result) // Ensure the returned search is the same
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetAllByOrg", func(t *testing.T) {
		// Setting up the mock response
		search := &models.Searches{ID: testUID}
		mockRepo.On("FindAllByOrg", mock.Anything, testOrgID).Return([]*models.Searches{search}, nil)

		// Calling the service method
		result, err := service.GetAllByOrg(context.Background(), testOrgID)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)           // Ensure exactly 1 result is returned
		assert.Equal(t, search, result[0]) // Ensure the returned search is the same
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test Update", func(t *testing.T) {
		// Setting up the mock response
		search := &models.Searches{ID: testUID}
		mockRepo.On("Update", mock.Anything, testUID, search).Return(nil)

		// Calling the service method
		err := service.Update(context.Background(), testUID, search)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test Delete", func(t *testing.T) {
		// Setting up the mock response
		mockRepo.On("Delete", mock.Anything, testUID).Return(nil)

		// Calling the service method
		err := service.Delete(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

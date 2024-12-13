package tests

import (
	"context"
	"findsafe/account/models"
	mocks2 "findsafe/account/models/mocks/repository"
	"findsafe/account/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOrgService(t *testing.T) {
	// Mock repository instance
	mockRepo := new(mocks2.MockOrgRepository)

	// Creating a new instance of OrgService with the mocked repository
	service := &services.OrgService{OrgRepository: mockRepo}

	// Test UUID
	testUID := uuid.New()

	t.Run("Test Get", func(t *testing.T) {
		// Setting up the mock response
		org := &models.Organization{ID: testUID}
		mockRepo.On("FindByID", mock.Anything, testUID).Return(org, nil)

		// Calling the service method
		result, err := service.Get(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, org, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test Update", func(t *testing.T) {
		org := &models.Organization{ID: testUID}
		mockRepo.On("Update", mock.Anything, testUID, org).Return(nil)

		// Call Update method
		err := service.Update(context.Background(), testUID, org)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test Delete", func(t *testing.T) {
		mockRepo.On("Delete", mock.Anything, testUID).Return(nil)

		// Call Delete method
		err := service.Delete(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetAll", func(t *testing.T) {
		org := &models.Organization{ID: testUID}
		mockRepo.On("FindAll", mock.Anything).Return([]*models.Organization{org}, nil)

		// Calling GetAll method
		result, err := service.GetAll(context.Background())

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)        // Ensure that exactly 1 result is returned
		assert.Equal(t, org, result[0]) // Ensure the returned org is the same
		mockRepo.AssertExpectations(t)
	})
}

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

func TestCertService(t *testing.T) {
	// Mock repository instance
	mockRepo := new(mocks2.MockCertRepository)

	// Creating a new instance of CertService with the mocked repository
	service := &services.CertService{CertRepository: mockRepo}

	// Test UUIDs
	testUID := uuid.New()
	testOrgID := uuid.New()

	t.Run("Test Get", func(t *testing.T) {
		// Setting up the mock response
		cert := &models.Certification{ID: testUID}
		mockRepo.On("FindByID", mock.Anything, testUID).Return(cert, nil)

		// Calling the service method
		result, err := service.Get(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, cert, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test Update", func(t *testing.T) {
		cert := &models.Certification{ID: testUID}
		mockRepo.On("Update", mock.Anything, testUID, cert).Return(nil)

		// Call Update method
		err := service.Update(context.Background(), testUID, cert)

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

	t.Run("Test GetByUserID", func(t *testing.T) {
		cert := &models.Certification{ID: testUID}
		mockRepo.On("FindByUserID", mock.Anything, testUID).Return([]*models.Certification{cert}, nil)

		// Calling GetByUserID method
		result, err := service.GetByUserID(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)         // Ensure that exactly 1 result is returned
		assert.Equal(t, cert, result[0]) // Ensure the returned cert is the same
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test GetByAccreditingOrg", func(t *testing.T) {
		cert := &models.Certification{ID: testUID}
		mockRepo.On("FindByAccreditingOrg", mock.Anything, testOrgID).Return([]*models.Certification{cert}, nil)

		// Calling GetByAccreditingOrg method
		result, err := service.GetByAccreditingOrg(context.Background(), testOrgID)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)         // Ensure that exactly 1 result is returned
		assert.Equal(t, cert, result[0]) // Ensure the returned cert is the same
		mockRepo.AssertExpectations(t)
	})
}

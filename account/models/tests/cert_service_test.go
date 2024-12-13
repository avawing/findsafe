package tests

import (
	"context"
	"findsafe/account/models"
	mocks "findsafe/account/models/mocks/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCertService(t *testing.T) {
	// Mock service instance
	mockCertService := new(mocks.MockCertService)

	// UUID for test
	testUID := uuid.New()
	testOrgID := uuid.New()

	t.Run("Test Get", func(t *testing.T) {
		// Setting up the mock response
		cert := &models.Certification{ID: testUID}
		mockCertService.On("Get", mock.Anything, testUID).Return(cert, nil)

		// Calling the service method
		result, err := mockCertService.Get(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, cert, result)
		mockCertService.AssertExpectations(t)
	})

	t.Run("Test Update", func(t *testing.T) {
		cert := &models.Certification{ID: testUID}
		mockCertService.On("Update", mock.Anything, testUID, cert).Return(nil)

		// Call Update method
		err := mockCertService.Update(context.Background(), testUID, cert)

		// Assertions
		assert.NoError(t, err)
		mockCertService.AssertExpectations(t)
	})

	t.Run("Test Delete", func(t *testing.T) {
		mockCertService.On("Delete", mock.Anything, testUID).Return(nil)

		// Call Delete method
		err := mockCertService.Delete(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		mockCertService.AssertExpectations(t)
	})

	t.Run("Test GetByUserID", func(t *testing.T) {
		cert := &models.Certification{ID: testUID}
		mockCertService.On("GetByUserID", mock.Anything, testUID).Return([]*models.Certification{cert}, nil)

		// Calling GetByUserID method
		result, err := mockCertService.GetByUserID(context.Background(), testUID)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)         // Ensure that exactly 1 result is returned
		assert.Equal(t, cert, result[0]) // Ensure the returned cert is the same
		mockCertService.AssertExpectations(t)
	})

	t.Run("Test GetByAccreditingOrg", func(t *testing.T) {
		cert := &models.Certification{ID: testUID}
		mockCertService.On("GetByAccreditingOrg", mock.Anything, testOrgID).Return([]*models.Certification{cert}, nil)

		// Calling GetByAccreditingOrg method
		result, err := mockCertService.GetByAccreditingOrg(context.Background(), testOrgID)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, result, 1)         // Ensure that exactly 1 result is returned
		assert.Equal(t, cert, result[0]) // Ensure the returned cert is the same
		mockCertService.AssertExpectations(t)
	})
}

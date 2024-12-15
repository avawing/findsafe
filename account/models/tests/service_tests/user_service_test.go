package tests

import (
	"context"
	"findsafe/account/models/mocks/repository"
	"findsafe/account/models/models"
	"findsafe/account/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUserService(t *testing.T) {
	t.Run("Get User", func(t *testing.T) {
		// Create a mock repository
		mockRepo := new(mocks.MockUserRepository)
		userService := services.NewUserService(&services.USConfig{UserRepository: mockRepo})

		uid := uuid.New()
		expectedUser := &models.User{ID: uid, FirstName: "John", LastName: "Doe"}

		// Mock the FindByCertID method to return a user
		mockRepo.On("FindUserByID", mock.Anything, uid).Return(expectedUser, nil)

		// Call the Get method
		user, err := userService.Get(context.Background(), uid)

		// Assert the results
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)

		// Assert that the FindByCertID method was called with the correct arguments
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateUser User", func(t *testing.T) {
		// Create a mock repository
		mockRepo := new(mocks.MockUserRepository)
		userService := services.NewUserService(&services.USConfig{UserRepository: mockRepo})

		uid := uuid.New()
		userToUpdate := &models.User{ID: uid, FirstName: "John", LastName: "Doe"}

		// Mock the UpdateUser method
		mockRepo.On("UpdateUser", mock.Anything, uid, userToUpdate).Return(nil)

		// Call the UpdateUser method
		err := userService.Update(context.Background(), uid, userToUpdate)

		// Assert the results
		assert.NoError(t, err)

		// Assert that the UpdateUser method was called with the correct arguments
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteUser User", func(t *testing.T) {
		// Create a mock repository
		mockRepo := new(mocks.MockUserRepository)
		userService := services.NewUserService(&services.USConfig{UserRepository: mockRepo})

		uid := uuid.New()

		// Mock the DeleteUser method
		mockRepo.On("DeleteUser", mock.Anything, uid).Return(nil)

		// Call the DeleteUser method
		err := userService.Delete(context.Background(), uid)

		// Assert the results
		assert.NoError(t, err)

		// Assert that the DeleteUser method was called with the correct arguments
		mockRepo.AssertExpectations(t)
	})

	t.Run("Get All Users in Org", func(t *testing.T) {
		// Create a mock repository
		mockRepo := new(mocks.MockUserRepository)
		userService := services.NewUserService(&services.USConfig{UserRepository: mockRepo})

		orgID := uuid.New()
		expectedUsers := []*models.User{
			{ID: uuid.New(), FirstName: "John", LastName: "Doe"},
			{ID: uuid.New(), FirstName: "Jane", LastName: "Doe"},
		}

		// Mock the FindByOrgID method
		mockRepo.On("FindByOrgID", mock.Anything, orgID).Return(expectedUsers, nil)

		// Call the GetAllInOrg method
		users, err := userService.GetAllInOrg(context.Background(), orgID)

		// Assert the results
		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)

		// Assert that the FindByOrgID method was called with the correct arguments
		mockRepo.AssertExpectations(t)
	})

	t.Run("Get All Users in Search", func(t *testing.T) {
		// Create a mock repository
		mockRepo := new(mocks.MockUserRepository)
		userService := services.NewUserService(&services.USConfig{UserRepository: mockRepo})

		searchID := uuid.New()
		expectedUsers := []*models.User{
			{ID: uuid.New(), FirstName: "John", LastName: "Doe"},
			{ID: uuid.New(), FirstName: "Jane", LastName: "Doe"},
		}

		// Mock the FindUsersBySearchID method
		mockRepo.On("FindUsersBySearchID", mock.Anything, searchID).Return(expectedUsers, nil)

		// Call the GetAllinSearch method
		users, err := userService.GetAllinSearch(context.Background(), searchID)

		// Assert the results
		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)

		// Assert that the FindUsersBySearchID method was called with the correct arguments
		mockRepo.AssertExpectations(t)
	})

	t.Run("Get All Users in Sortie", func(t *testing.T) {
		// Create a mock repository
		mockRepo := new(mocks.MockUserRepository)
		userService := services.NewUserService(&services.USConfig{UserRepository: mockRepo})

		sortieID := uuid.New()
		expectedUsers := []*models.User{
			{ID: uuid.New(), FirstName: "John", LastName: "Doe"},
			{ID: uuid.New(), FirstName: "Jane", LastName: "Doe"},
		}

		// Mock the FindBySortieID method
		mockRepo.On("FindBySortieID", mock.Anything, sortieID).Return(expectedUsers, nil)

		// Call the GetAllInSortie method
		users, err := userService.GetAllInSortie(context.Background(), sortieID)

		// Assert the results
		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)

		// Assert that the FindBySortieID method was called with the correct arguments
		mockRepo.AssertExpectations(t)
	})
}

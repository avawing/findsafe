package mocks

import (
	"context"
	"findsafe/account/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockResourceRepository struct {
	mock.Mock
}

func (m *MockResourceRepository) FindByID(c context.Context, uid uuid.UUID) (*models.Resource, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Resource), args.Error(1)
}

func (m *MockResourceRepository) Update(c context.Context, uid uuid.UUID, user *models.Resource) error {
	args := m.Called(c, uid, user)
	return args.Error(0)
}

func (m *MockResourceRepository) Delete(c context.Context, uid uuid.UUID) error {
	args := m.Called(c, uid)
	return args.Error(0)
}

func (m *MockResourceRepository) FindByOwnerID(c context.Context, ownerID uuid.UUID) ([]*models.Resource, error) {
	args := m.Called(c, &ownerID)
	return args.Get(0).([]*models.Resource), args.Error(1)
}

func (m *MockResourceRepository) FindByIssuedID(c context.Context, issuedID uuid.UUID) ([]*models.Resource, error) {
	args := m.Called(c, &issuedID)
	return args.Get(0).([]*models.Resource), args.Error(1)
}

func (m *MockResourceRepository) FindByTeamID(c context.Context, teamID uuid.UUID) ([]*models.Resource, error) {
	args := m.Called(c, &teamID)
	return args.Get(0).([]*models.Resource), args.Error(1)
}

func (m *MockResourceRepository) FindAvailable(c context.Context, searchID uuid.UUID) ([]*models.Resource, error) {
	args := m.Called(c, &searchID)
	return args.Get(0).([]*models.Resource), args.Error(1)
}

func (m *MockResourceRepository) FindUnreturned(c context.Context, uid uuid.UUID) ([]*models.Resource, error) {
	args := m.Called(c, &uid)
	return args.Get(0).([]*models.Resource), args.Error(1)
}

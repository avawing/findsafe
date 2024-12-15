package mocks

import (
	"context"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockResourceService struct {
	mock.Mock
}

func (m *MockResourceService) Get(c context.Context, uid uuid.UUID) (*models.Resource, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Resource), args.Error(1)
}

func (m *MockResourceService) Update(c context.Context, uid uuid.UUID, user *models.Resource) error {
	args := m.Called(c, uid, user)
	return args.Error(0)
}

func (m *MockResourceService) Delete(c context.Context, uid uuid.UUID) error {
	args := m.Called(c, uid)
	return args.Error(0)
}

func (m *MockResourceService) GetByOwnerID(c context.Context, ownerID uuid.UUID) ([]*models.Resource, error) {
	args := m.Called(c, ownerID)
	return args.Get(0).([]*models.Resource), args.Error(1)
}

func (m *MockResourceService) GetByIssuedID(c context.Context, issuedID uuid.UUID) ([]*models.Resource, error) {
	args := m.Called(c, issuedID)
	return args.Get(0).([]*models.Resource), args.Error(1)
}

func (m *MockResourceService) GetByTeamID(c context.Context, teamID uuid.UUID) ([]*models.Resource, error) {
	args := m.Called(c, teamID)
	return args.Get(0).([]*models.Resource), args.Error(1)
}

func (m *MockResourceService) GetByAvailable(c context.Context, searchID uuid.UUID) ([]*models.Resource, error) {
	args := m.Called(c, searchID)
	return args.Get(0).([]*models.Resource), args.Error(1)
}

func (m *MockResourceService) GetUnreturned(c context.Context, searchID uuid.UUID) ([]*models.Resource, error) {
	args := m.Called(c, searchID)
	return args.Get(0).([]*models.Resource), args.Error(1)
}

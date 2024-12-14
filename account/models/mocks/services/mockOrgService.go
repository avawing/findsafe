package mocks

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockOrgService struct {
	mock.Mock
}

func (m *MockOrgService) GetAll(c context.Context) ([]*models.Organization, error) {
	args := m.Called(c)
	return args.Get(0).([]*models.Organization), args.Error(1)
}

func (m *MockOrgService) Get(c context.Context, uid uuid.UUID) (*models.Organization, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Organization), args.Error(1)
}

func (m *MockOrgService) Update(c context.Context, uid uuid.UUID, user *models.Organization) error {
	args := m.Called(c, uid, user)
	return args.Error(0)
}

func (m *MockOrgService) Delete(c context.Context, uid uuid.UUID) error {
	args := m.Called(c, uid)
	return args.Error(0)
}

func (m *MockOrgService) GetAllInSearch(c context.Context, uid uuid.UUID) ([]*models.Organization, error) {
	args := m.Called(c, uid)
	return args.Get(0).([]*models.Organization), args.Error(1)
}

package mocks

import (
	"context"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockOrgRepository struct {
	mock.Mock
}

func (m *MockOrgRepository) FindOrgByID(c context.Context, uid uuid.UUID) (*models.Organization, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Organization), args.Error(1)
}

func (m *MockOrgRepository) FindAllOrgs(c context.Context) ([]*models.Organization, error) {
	args := m.Called(c)
	return args.Get(0).([]*models.Organization), args.Error(1)
}

func (m *MockOrgRepository) UpdateOrg(c context.Context, uid uuid.UUID, user *models.Organization) error {
	args := m.Called(c, uid, user)
	return args.Error(0)
}

func (m *MockOrgRepository) DeleteOrg(c context.Context, uid uuid.UUID) error {
	args := m.Called(c, uid)
	return args.Error(0)
}

func (m *MockOrgRepository) FindAllInSearch(c context.Context, uid uuid.UUID) ([]*models.Organization, error) {
	args := m.Called(c, uid)
	return args.Get(0).([]*models.Organization), args.Error(1)
}

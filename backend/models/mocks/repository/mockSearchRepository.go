package mocks

import (
	"context"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockSearchRepository struct {
	mock.Mock
}

func (m *MockSearchRepository) FindBySearchID(c context.Context, uid uuid.UUID) (*models.Searches, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Searches), args.Error(1)
}
func (m *MockSearchRepository) FindAllSearches(c context.Context) ([]*models.Searches, error) {
	args := m.Called(c)
	return args.Get(0).([]*models.Searches), args.Error(1)
}
func (m *MockSearchRepository) FindAllBySubject(c context.Context, uid uuid.UUID) (*models.Searches, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Searches), args.Error(1)
}
func (m *MockSearchRepository) FindAllByOrg(c context.Context, orgID uuid.UUID) ([]*models.Searches, error) {
	args := m.Called(c, orgID)
	return args.Get(0).([]*models.Searches), args.Error(1)
}
func (m *MockSearchRepository) UpdateSearch(c context.Context, uid uuid.UUID, user *models.Searches) error {
	args := m.Called(c, uid, user)
	return args.Error(0)
}
func (m *MockSearchRepository) DeleteSearch(c context.Context, uid uuid.UUID) error {
	args := m.Called(c, uid)
	return args.Error(0)
}

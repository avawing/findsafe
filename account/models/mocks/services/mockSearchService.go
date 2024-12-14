package mocks

import (
	"context"
	"findsafe/account/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockSearchService struct {
	mock.Mock
}

func (m *MockSearchService) Get(c context.Context, uid uuid.UUID) (*models.Searches, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Searches), args.Error(1)
}
func (m *MockSearchService) GetAll(c context.Context) ([]*models.Searches, error) {
	args := m.Called(c)
	return args.Get(0).([]*models.Searches), args.Error(1)
}
func (m *MockSearchService) GetAllBySubject(c context.Context, uid uuid.UUID) (*models.Searches, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Searches), args.Error(1)
}
func (m *MockSearchService) GetAllByOrg(c context.Context, orgID uuid.UUID) ([]*models.Searches, error) {
	args := m.Called(c, orgID)
	return args.Get(0).([]*models.Searches), args.Error(1)
}
func (m *MockSearchService) Update(c context.Context, uid uuid.UUID, user *models.Searches) error {
	args := m.Called(c, uid, user)
	return args.Error(0)
}
func (m *MockSearchService) Delete(c context.Context, uid uuid.UUID) error {
	args := m.Called(c, uid)
	return args.Error(0)
}

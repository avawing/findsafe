package mocks

import (
	"context"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Get(c context.Context, uid uuid.UUID) (*models.User, error) {
	ret := m.Called(c, uid)
	return ret.Get(0).(*models.User), ret.Error(1)
}

func (m *MockUserService) Update(c context.Context, uid uuid.UUID, user *models.User) error {
	ret := m.Called(c, uid, user)
	return ret.Error(0)
}

func (m *MockUserService) Delete(c context.Context, uid uuid.UUID) error {
	ret := m.Called(c, uid)
	return ret.Error(0)
}

func (m *MockUserService) GetAllInOrg(c context.Context, orgID uuid.UUID) ([]*models.User, error) {
	ret := m.Called(c, orgID)
	return ret.Get(0).([]*models.User), ret.Error(1)
}

func (m *MockUserService) GetAllinSearch(c context.Context, searchID uuid.UUID) ([]*models.User, error) {
	ret := m.Called(c, searchID)
	return ret.Get(0).([]*models.User), ret.Error(1)
}

func (m *MockUserService) GetAllInSortie(c context.Context, sortieID uuid.UUID) ([]*models.User, error) {
	ret := m.Called(c, sortieID)
	return ret.Get(0).([]*models.User), ret.Error(1)
}

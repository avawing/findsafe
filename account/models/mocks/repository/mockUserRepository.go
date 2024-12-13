package mocks

import (
	"context"
	"findsafe/account/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(c context.Context, uid uuid.UUID) (*models.User, error) {
	ret := m.Called(c, uid)
	return ret.Get(0).(*models.User), ret.Error(1)
}
func (m *MockUserRepository) Update(c context.Context, uid uuid.UUID, user *models.User) error {
	ret := m.Called(c, uid, user)
	return ret.Error(0)
}
func (m *MockUserRepository) Delete(c context.Context, uid uuid.UUID) error {
	ret := m.Called(c, uid)
	return ret.Error(0)
}
func (m *MockUserRepository) FindByOrgID(c context.Context, orgID uuid.UUID) ([]*models.User, error) {
	ret := m.Called(c, orgID)
	return ret.Get(0).([]*models.User), ret.Error(1)
}
func (m *MockUserRepository) FindBySearchID(c context.Context, searchID uuid.UUID) ([]*models.User, error) {
	ret := m.Called(c, searchID)
	return ret.Get(0).([]*models.User), ret.Error(1)
}
func (m *MockUserRepository) FindBySortieID(c context.Context, sortieID uuid.UUID) ([]*models.User, error) {
	ret := m.Called(c, sortieID)
	return ret.Get(0).([]*models.User), ret.Error(1)
}
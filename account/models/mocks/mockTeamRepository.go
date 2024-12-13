package mocks

import (
	"context"
	"findsafe/account/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockTeamRepository struct {
	mock.Mock
}

func (m *MockTeamRepository) FindByID(c context.Context, uid uuid.UUID) (*models.Team, error) {
	ret := m.Called(c, uid)
	return ret.Get(0).(*models.Team), ret.Error(1)
}
func (m *MockTeamRepository) Update(c context.Context, uid uuid.UUID, user *models.Team) error {
	ret := m.Called(c, uid, user)
	return ret.Error(0)
}
func (m *MockTeamRepository) Delete(c context.Context, uid uuid.UUID) error {
	ret := m.Called(c, uid)
	return ret.Error(0)
}
func (m *MockTeamRepository) FindBySearch(c context.Context, uid uuid.UUID) ([]*models.Team, error) {
	ret := m.Called(c, uid)
	return ret.Get(0).([]*models.Team), ret.Error(1)
}
func (m *MockTeamRepository) FindByEmptySortie(c context.Context) ([]*models.Team, error) {
	ret := m.Called(c)
	return ret.Get(0).([]*models.Team), ret.Error(1)
}
func (m *MockTeamRepository) FindBySortie(c context.Context, uid uuid.UUID) (*models.Team, error) {
	ret := m.Called(c, uid)
	return ret.Get(0).(*models.Team), ret.Error(1)
}

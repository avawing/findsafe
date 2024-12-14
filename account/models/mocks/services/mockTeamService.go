package mocks

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockTeamService struct {
	mock.Mock
}

func (m *MockTeamService) Get(c context.Context, uid uuid.UUID) (*models.Team, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Team), args.Error(1)
}

func (m *MockTeamService) Update(c context.Context, uid uuid.UUID, user *models.Team) error {
	args := m.Called(c, uid, user)
	return args.Error(0)
}

func (m *MockTeamService) Delete(c context.Context, uid uuid.UUID) error {
	args := m.Called(c, uid)
	return args.Error(0)
}

func (m *MockTeamService) GetAllinSearch(c context.Context, uid uuid.UUID) ([]*models.Team, error) {
	args := m.Called(c, uid)
	return args.Get(0).([]*models.Team), args.Error(1)
}

func (m *MockTeamService) GetAllWithoutSortie(c context.Context) ([]*models.Team, error) {
	args := m.Called(c)
	return args.Get(0).([]*models.Team), args.Error(1)
}

func (m *MockTeamService) GetBySortie(c context.Context, uid uuid.UUID) (*models.Team, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Team), args.Error(1)
}

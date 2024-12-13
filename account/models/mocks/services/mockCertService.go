package mocks

import (
	"context"
	"findsafe/account/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCertService struct {
	mock.Mock
}

func (m *MockCertService) Get(c context.Context, uid uuid.UUID) (*models.Certification, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Certification), args.Error(1)
}

func (m *MockCertService) Update(c context.Context, uid uuid.UUID, user *models.Certification) error {
	args := m.Called(c, uid, user)
	return args.Error(0)
}

func (m *MockCertService) Delete(c context.Context, uid uuid.UUID) error {
	args := m.Called(c, uid)
	return args.Error(0)
}

func (m *MockCertService) GetByUserID(c context.Context, uid uuid.UUID) ([]*models.Certification, error) {
	args := m.Called(c, uid)
	return args.Get(0).([]*models.Certification), args.Error(1)
}

func (m *MockCertService) GetByAccreditingOrg(c context.Context, org uuid.UUID) ([]*models.Certification, error) {
	args := m.Called(c, org)
	return args.Get(0).([]*models.Certification), args.Error(1)
}

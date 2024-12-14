package mocks

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCertRepository struct {
	mock.Mock
}

func (m *MockCertRepository) FindByCertID(c context.Context, uid uuid.UUID) (*models.Certification, error) {
	args := m.Called(c, uid)
	return args.Get(0).(*models.Certification), args.Error(1)
}

func (m *MockCertRepository) UpdateCert(c context.Context, uid uuid.UUID, user *models.Certification) error {
	args := m.Called(c, uid, user)
	return args.Error(0)
}

func (m *MockCertRepository) DeleteByCertID(c context.Context, uid uuid.UUID) error {
	args := m.Called(c, uid)
	return args.Error(0)
}

func (m *MockCertRepository) FindByUserID(c context.Context, uid uuid.UUID) ([]*models.Certification, error) {
	args := m.Called(c, uid)
	return args.Get(0).([]*models.Certification), args.Error(1)
}

func (m *MockCertRepository) FindByAccreditingOrg(c context.Context, org uuid.UUID) ([]*models.Certification, error) {
	args := m.Called(c, org)
	return args.Get(0).([]*models.Certification), args.Error(1)
}

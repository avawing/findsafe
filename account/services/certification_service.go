package services

import (
	"context"
	"findsafe/account/models"
	"findsafe/account/models/interfaces"
	"github.com/google/uuid"
)

// CertService acts as a struct for injecting an implementation of UserRepository
// for use in service methods.
type CertService struct {
	CertRepository interfaces.CertRepository
}

// NewCertService is a factory function for initializing a UserService with
// its repository layer dependencies.
func NewCertService(c *USConfig) *CertService {
	return &CertService{
		CertRepository: c.CertRepository,
	}
}

func (s *CertService) Get(c context.Context, uid uuid.UUID) (*models.Certification, error) {
	return s.CertRepository.FindByID(c, uid)
}
func (s *CertService) Update(c context.Context, uid uuid.UUID, user *models.Certification) error {
	return s.CertRepository.Update(c, uid, user)
}
func (s *CertService) Delete(c context.Context, uid uuid.UUID) error {
	return s.CertRepository.Delete(c, uid)
}
func (s *CertService) GetByUserID(c context.Context, uid uuid.UUID) ([]*models.Certification, error) {
	return s.CertRepository.FindByUserID(c, uid)
}
func (s *CertService) GetByAccreditingOrg(c context.Context, org uuid.UUID) ([]*models.Certification, error) {
	return s.CertRepository.FindByAccreditingOrg(c, org)
}

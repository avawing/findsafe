package services

import (
	"context"
	"findsafe/account/models/interfaces"
	"findsafe/account/models/models"
	"github.com/google/uuid"
)

// USConfig will hold repositories that will eventually be injected into
// this service layer.
type USConfig struct {
	UserRepository     interfaces.UserRepository
	TeamRepository     interfaces.TeamRepository
	ResourceRepository interfaces.ResourceRepository
	OrgRepository      interfaces.OrgRepository
	CertRepository     interfaces.CertRepository
	SearchRepository   interfaces.SearchRepository
}

// UserService acts as a struct for injecting an implementation of UserRepository
// for use in service methods.
type UserService struct {
	UserRepository interfaces.UserRepository
}

// NewUserService is a factory function for initializing a UserService with
// its repository layer dependencies.
func NewUserService(c *USConfig) *UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

func (s *UserService) Get(c context.Context, uid uuid.UUID) (*models.User, error) {
	return s.UserRepository.FindUserByID(c, uid)
}
func (s *UserService) Update(c context.Context, uid uuid.UUID, user *models.User) error {
	return s.UserRepository.UpdateUser(c, uid, user)
}
func (s *UserService) Delete(c context.Context, uid uuid.UUID) error {
	return s.UserRepository.DeleteUser(c, uid)
}
func (s *UserService) GetAllInOrg(c context.Context, orgID uuid.UUID) ([]*models.User, error) {
	return s.UserRepository.FindByOrgID(c, orgID)
}
func (s *UserService) GetAllinSearch(c context.Context, searchID uuid.UUID) ([]*models.User, error) {
	return s.UserRepository.FindUsersBySearchID(c, searchID)
}
func (s *UserService) GetAllInSortie(c context.Context, sortieID uuid.UUID) ([]*models.User, error) {
	return s.UserRepository.FindBySortieID(c, sortieID)
}

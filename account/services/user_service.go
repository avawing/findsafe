package services

import (
	"context"
	"findsafe/account/models"
	"findsafe/account/models/interfaces"
	"github.com/google/uuid"
)

// UserService acts as a struct for injecting an implementation of UserRepository
// for use in service methods.
type UserService struct {
	UserRepository interfaces.UserRepository
}

// USConfig will hold repositories that will eventually be injected into
// this service layer.
type USConfig struct {
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
	return s.UserRepository.FindByID(c, uid)
}
func (s *UserService) Update(c context.Context, uid uuid.UUID, user *models.User) error {
	return s.UserRepository.Update(c, uid, user)
}
func (s *UserService) Delete(c context.Context, uid uuid.UUID) error {
	return s.UserRepository.Delete(c, uid)
}
func (s *UserService) GetAllInOrg(c context.Context, orgID uuid.UUID) ([]*models.User, error) {
	return s.UserRepository.FindByOrgID(c, orgID)
}
func (s *UserService) GetAllinSearch(c context.Context, searchID uuid.UUID) ([]*models.User, error) {
	return s.UserRepository.FindBySearchID(c, searchID)
}
func (s *UserService) GetAllInSortie(c context.Context, sortieID uuid.UUID) ([]*models.User, error) {
	return s.UserRepository.FindBySortieID(c, sortieID)
}

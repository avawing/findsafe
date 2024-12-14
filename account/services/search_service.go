package services

import (
	"context"
	"findsafe/account/models"
	"findsafe/account/models/interfaces"
	"github.com/google/uuid"
)

// SearchService acts as a struct for injecting an implementation of UserRepository
// for use in service methods.
type SearchService struct {
	SearchRepository interfaces.SearchRepository
}

// NewSearchService is a factory function for initializing a UserService with
// its repository layer dependencies.
func NewSearchService(c *USConfig) *SearchService {
	return &SearchService{
		SearchRepository: c.SearchRepository,
	}
}

func (s *SearchService) Get(c context.Context, uid uuid.UUID) (*models.Searches, error) {
	return s.SearchRepository.FindByID(c, uid)
}
func (s *SearchService) GetAll(c context.Context) ([]*models.Searches, error) {
	return s.SearchRepository.FindAll(c)
}

func (s *SearchService) GetAllBySubject(c context.Context, uid uuid.UUID) (*models.Searches, error) {
	return s.SearchRepository.FindAllBySubject(c, uid)
}

func (s *SearchService) GetAllByOrg(c context.Context, orgID uuid.UUID) ([]*models.Searches, error) {
	return s.SearchRepository.FindAllByOrg(c, orgID)
}

func (s *SearchService) Update(c context.Context, uid uuid.UUID, user *models.Searches) error {
	return s.SearchRepository.Update(c, uid, user)
}

func (s *SearchService) Delete(c context.Context, uid uuid.UUID) error {
	return s.SearchRepository.Delete(c, uid)
}

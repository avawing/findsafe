package services

import (
	"context"
	"findsafe/backend/models/interfaces"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
)

// TeamService acts as a struct for injecting an implementation of UserRepository
// for use in service methods.
type TeamService struct {
	TeamRepository interfaces.TeamRepository
}

// NewTeamService is a factory function for initializing a UserService with
// its repository layer dependencies.
func NewTeamService(c *USConfig) *TeamService {
	return &TeamService{
		TeamRepository: c.TeamRepository,
	}
}

func (t *TeamService) Get(c context.Context, uid uuid.UUID) (*models.Team, error) {
	return t.TeamRepository.FindTeamByID(c, uid)
}

func (t *TeamService) Update(c context.Context, uid uuid.UUID, user *models.Team) error {
	return t.TeamRepository.UpdateTeam(c, uid, user)
}

func (t *TeamService) Delete(c context.Context, uid uuid.UUID) error {
	return t.TeamRepository.DeleteTeam(c, uid)
}

func (t *TeamService) GetAllinSearch(c context.Context, uid uuid.UUID) ([]*models.Team, error) {
	return t.TeamRepository.FindBySearch(c, uid)
}

func (t *TeamService) GetAllWithoutSortie(c context.Context) ([]*models.Team, error) {
	return t.TeamRepository.FindByEmptySortie(c)
}

func (t *TeamService) GetBySortie(c context.Context, uid uuid.UUID) (*models.Team, error) {
	return t.TeamRepository.FindBySortie(c, uid)
}

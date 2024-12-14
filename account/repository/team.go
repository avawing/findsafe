package repository

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewTeamRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindTeamByID(c context.Context, uid uuid.UUID) (*models.Team, error) {
	return &models.Team{}, nil
}
func (r *Repository) UpdateTeam(c context.Context, uid uuid.UUID, user *models.Team) error {
	return nil
}
func (r *Repository) DeleteTeam(c context.Context, uid uuid.UUID) error {
	return nil
}
func (r *Repository) FindBySearch(c context.Context, uid uuid.UUID) ([]*models.Team, error) {
	return []*models.Team{}, nil
}
func (r *Repository) FindByEmptySortie(c context.Context) ([]*models.Team, error) {
	return []*models.Team{}, nil
}
func (r *Repository) FindBySortie(c context.Context, uid uuid.UUID) (*models.Team, error) {
	return &models.Team{}, nil
}

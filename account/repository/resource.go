package repository

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewResourceRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindResourceByID(c context.Context, uid uuid.UUID) (*models.Resource, error) {
	return &models.Resource{}, nil
}
func (r *Repository) UpdateResource(c context.Context, uid uuid.UUID, resource *models.Resource) error {
	return nil
}
func (r *Repository) DeleteResource(c context.Context, uid uuid.UUID) error {
	return nil
}
func (r *Repository) FindByOwnerID(c context.Context, ownerID uuid.UUID) ([]*models.Resource, error) {
	return []*models.Resource{}, nil
}
func (r *Repository) FindByIssuedID(c context.Context, issuedID uuid.UUID) ([]*models.Resource, error) {
	return []*models.Resource{}, nil
}
func (r *Repository) FindByTeamID(c context.Context, teamID uuid.UUID) ([]*models.Resource, error) {
	return []*models.Resource{}, nil
}
func (r *Repository) FindAvailable(c context.Context, searchID uuid.UUID) ([]*models.Resource, error) {
	return []*models.Resource{}, nil
}
func (r *Repository) FindUnreturned(c context.Context, uid uuid.UUID) ([]*models.Resource, error) {
	return []*models.Resource{}, nil
}

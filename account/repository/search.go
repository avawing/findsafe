package repository

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewSearchRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindBySearchID(c context.Context, uid uuid.UUID) (*models.Searches, error) {
	return &models.Searches{}, nil
}
func (r *Repository) FindAllSearches(c context.Context) ([]*models.Searches, error) {
	return []*models.Searches{}, nil
}
func (r *Repository) FindAllBySubject(c context.Context, uid uuid.UUID) (*models.Searches, error) {
	return &models.Searches{}, nil
}
func (r *Repository) FindAllByOrg(c context.Context, orgID uuid.UUID) ([]*models.Searches, error) {
	return []*models.Searches{}, nil
}
func (r *Repository) UpdateSearch(c context.Context, uid uuid.UUID, user *models.Searches) error {
	return nil
}
func (r *Repository) DeleteSearch(c context.Context, uid uuid.UUID) error {
	return nil
}

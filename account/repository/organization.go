package repository

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewOrgRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (t *Repository) FindOrgByID(c context.Context, uid uuid.UUID) (*models.Organization, error) {
	return &models.Organization{}, nil
}

func (t *Repository) FindAllOrgs(c context.Context) ([]*models.Organization, error) {
	return []*models.Organization{}, nil
}

func (t *Repository) FindAllInSearch(c context.Context, uid uuid.UUID) ([]*models.Organization, error) {
	return []*models.Organization{}, nil
}

func (t *Repository) UpdateOrg(c context.Context, uid uuid.UUID, user *models.Organization) error {
	return nil
}

func (t *Repository) DeleteOrg(c context.Context, uid uuid.UUID) error {
	return nil
}

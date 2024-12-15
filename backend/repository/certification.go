package repository

import (
	"context"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewCertRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindByCertID(c context.Context, uid uuid.UUID) (*models.Certification, error) {
	return &models.Certification{}, nil
}
func (r *Repository) UpdateCert(c context.Context, uid uuid.UUID, user *models.Certification) error {
	return nil
}
func (r *Repository) DeleteByCertID(c context.Context, uid uuid.UUID) error {
	return nil
}
func (r *Repository) FindByUserID(c context.Context, uid uuid.UUID) ([]*models.Certification, error) {
	return []*models.Certification{}, nil
}
func (r *Repository) FindByAccreditingOrg(c context.Context, org uuid.UUID) ([]*models.Certification, error) {
	return []*models.Certification{}, nil
}

package repository

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindUserByID(c context.Context, uid uuid.UUID) (*models.User, error) {
	return &models.User{}, nil
}
func (r *Repository) UpdateUser(c context.Context, uid uuid.UUID, user *models.User) error {
	return nil
}
func (r *Repository) DeleteUser(c context.Context, uid uuid.UUID) error {
	return nil
}
func (r *Repository) FindByOrgID(c context.Context, orgID uuid.UUID) ([]*models.User, error) {
	return []*models.User{}, nil
}
func (r *Repository) FindUsersBySearchID(c context.Context, searchID uuid.UUID) ([]*models.User, error) {
	return []*models.User{}, nil
}
func (r *Repository) FindBySortieID(c context.Context, sortieID uuid.UUID) ([]*models.User, error) {
	return []*models.User{}, nil
}

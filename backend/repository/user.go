package repository

import (
	"context"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindUserByID(c context.Context, uid uuid.UUID) (*models.User, error) {
	var user models.User
	if result := r.DB.WithContext(c).Model(&models.User{}).Where("id = ?", uid).First(&user); result.Error != nil {
		return &models.User{}, result.Error
	}
	return &user, nil
}
func (r *Repository) UpdateUser(c context.Context, uid uuid.UUID, user *models.User) error {
	return nil
}
func (r *Repository) DeleteUser(c context.Context, uid uuid.UUID) error {
	if result := r.DB.WithContext(c).Model(&models.User{}).Where("id = ?", uid).Updates(models.User{
		Model: gorm.Model{
			DeletedAt: gorm.DeletedAt{
				Time:  time.Now().UTC(),
				Valid: true,
			},
		},
	}); result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *Repository) FindByOrgID(c context.Context, orgID uuid.UUID) ([]*models.User, error) {
	var users []*models.User
	if result := r.DB.WithContext(c).Where("org_id = ?", orgID).Find(&users); result.Error != nil {
		return []*models.User{}, result.Error
	}
	return users, nil
}
func (r *Repository) FindUsersBySearchID(c context.Context, searchID uuid.UUID) ([]*models.User, error) {
	var users []*models.User
	if result := r.DB.WithContext(c).Where("search_id = ?", searchID).Find(&users); result.Error != nil {
		return []*models.User{}, result.Error
	}
	return users, nil
}
func (r *Repository) FindBySortieID(c context.Context, sortieID uuid.UUID) ([]*models.User, error) {
	var users []*models.User
	if result := r.DB.WithContext(c).Where("sortie_id = ?", sortieID).Find(&users); result.Error != nil {
		return []*models.User{}, result.Error
	}
	return users, nil
}

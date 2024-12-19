package repository

import (
	"context"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func NewSearchRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindBySearchID(c context.Context, uid uuid.UUID) (*models.Searches, error) {
	var search models.Searches
	if result := r.DB.WithContext(c).Model(&models.Searches{}).Where("id = ?", uid).First(&search); result.Error != nil {
		return &models.Searches{}, result.Error
	}
	return &search, nil
}
func (r *Repository) FindAllSearches(c context.Context) ([]*models.Searches, error) {
	var searches []*models.Searches
	if result := r.DB.WithContext(c).Find(&searches); result.Error != nil {
		return nil, result.Error
	}
	return searches, nil
}
func (r *Repository) FindAllBySubject(c context.Context, uid uuid.UUID) (*models.Searches, error) {
	return &models.Searches{}, nil
}
func (r *Repository) FindAllByOrg(c context.Context, orgID uuid.UUID) ([]*models.Searches, error) {
	var searches []*models.Searches
	if result := r.DB.WithContext(c).Where("organization_id", orgID).Find(&searches); result.Error != nil {
		return nil, result.Error
	}
	return searches, nil
}
func (r *Repository) UpdateSearch(c context.Context, uid uuid.UUID, user *models.Searches) error {
	return nil
}
func (r *Repository) DeleteSearch(c context.Context, uid uuid.UUID) error {
	if result := r.DB.WithContext(c).Model(&models.Searches{}).Where("id = ?", uid).Updates(models.Searches{
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

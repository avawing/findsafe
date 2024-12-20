package repository

import (
	"context"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func NewTeamRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindTeamByID(c context.Context, uid uuid.UUID) (*models.Team, error) {
	var team models.Team
	if result := r.DB.WithContext(c).Model(&models.Team{}).Where("id = ?", uid).First(&team); result.Error != nil {
		return &models.Team{}, result.Error
	}
	return &team, nil
}
func (r *Repository) UpdateTeam(c context.Context, uid uuid.UUID, user *models.Team) error {
	return nil
}
func (r *Repository) DeleteTeam(c context.Context, uid uuid.UUID) error {
	if result := r.DB.WithContext(c).Model(&models.Team{}).Where("id = ?", uid).Updates(models.Team{
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
func (r *Repository) FindBySearch(c context.Context, uid uuid.UUID) ([]*models.Team, error) {
	var teams []*models.Team
	if result := r.DB.WithContext(c).Model(&models.Team{}).Where("search_id", uid).Find(&teams).Error; result.Error != nil {
		return []*models.Team{}, nil
	}
	return teams, nil
}
func (r *Repository) FindByEmptySortie(c context.Context) ([]*models.Team, error) {
	return []*models.Team{}, nil
}
func (r *Repository) FindBySortie(c context.Context, uid uuid.UUID) (*models.Team, error) {
	return &models.Team{}, nil
}

package interfaces

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
)

// TeamService defines methods the handler layer expects
// any service it interacts with to implement
type TeamService interface {
	Get(c context.Context, uid uuid.UUID) (*models.Team, error)
	Update(c context.Context, uid uuid.UUID, user *models.Team) error
	Delete(c context.Context, uid uuid.UUID) error
	GetAllinSearch(c context.Context, uid uuid.UUID) ([]*models.Team, error)
	GetAllWithoutSortie(c context.Context) ([]*models.Team, error)
	GetBySortie(c context.Context, uid uuid.UUID) (*models.Team, error)
}

// TeamRepository defines methods the service layer expects
// any repository it interacts with to implement
type TeamRepository interface {
	FindTeamByID(c context.Context, uid uuid.UUID) (*models.Team, error)
	UpdateTeam(c context.Context, uid uuid.UUID, user *models.Team) error
	DeleteTeam(c context.Context, uid uuid.UUID) error
	FindBySearch(c context.Context, uid uuid.UUID) ([]*models.Team, error)
	FindByEmptySortie(c context.Context) ([]*models.Team, error)
	FindBySortie(c context.Context, uid uuid.UUID) (*models.Team, error)
}

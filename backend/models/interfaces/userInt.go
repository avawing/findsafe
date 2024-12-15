package interfaces

import (
	"context"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
)

// UserService defines methods the handler layer expects
// any service it interacts with to implement
type UserService interface {
	Get(c context.Context, uid uuid.UUID) (*models.User, error)
	Update(c context.Context, uid uuid.UUID, user *models.User) error
	Delete(c context.Context, uid uuid.UUID) error
	GetAllInOrg(c context.Context, orgID uuid.UUID) ([]*models.User, error)
	GetAllinSearch(c context.Context, searchID uuid.UUID) ([]*models.User, error)
	GetAllInSortie(c context.Context, sortID uuid.UUID) ([]*models.User, error)
}

// UserRepository defines methods the service layer expects
// any repository it interacts with to implement
type UserRepository interface {
	FindUserByID(c context.Context, uid uuid.UUID) (*models.User, error)
	UpdateUser(c context.Context, uid uuid.UUID, user *models.User) error
	DeleteUser(c context.Context, uid uuid.UUID) error
	FindByOrgID(c context.Context, orgID uuid.UUID) ([]*models.User, error)
	FindUsersBySearchID(c context.Context, searchID uuid.UUID) ([]*models.User, error)
	FindBySortieID(c context.Context, sortieID uuid.UUID) ([]*models.User, error)
}

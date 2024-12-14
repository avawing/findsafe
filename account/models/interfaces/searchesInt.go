package interfaces

import (
	"context"
	"findsafe/account/models"
	"github.com/google/uuid"
)

// SearchService defines methods the handler layer expects
// any service it interacts with to implement
type SearchService interface {
	Get(c context.Context, uid uuid.UUID) (*models.Searches, error)
	GetAll(c context.Context) ([]*models.Searches, error)
	GetAllBySubject(c context.Context, uid uuid.UUID) (*models.Searches, error)
	GetAllByOrg(c context.Context, orgID uuid.UUID) ([]*models.Searches, error)
	Update(c context.Context, uid uuid.UUID, user *models.Searches) error
	Delete(c context.Context, uid uuid.UUID) error
}

// SearchRepository defines methods the service layer expects
// any repository it interacts with to implement
type SearchRepository interface {
	FindByID(c context.Context, uid uuid.UUID) (*models.Searches, error)
	FindAll(c context.Context) ([]*models.Searches, error)
	FindAllBySubject(c context.Context, uid uuid.UUID) (*models.Searches, error)
	FindAllByOrg(c context.Context, orgID uuid.UUID) ([]*models.Searches, error)
	Update(c context.Context, uid uuid.UUID, user *models.Searches) error
	Delete(c context.Context, uid uuid.UUID) error
}

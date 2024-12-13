package interfaces

import (
	"context"
	"findsafe/account/models"
	"github.com/google/uuid"
)

// OrgService defines methods the handler layer expects
// any service it interacts with to implement
type OrgService interface {
	GetAll(c context.Context) ([]*models.Organization, error)
	Get(c context.Context, uid uuid.UUID) (*models.Organization, error)
	Update(c context.Context, uid uuid.UUID, user *models.Organization) error
	Delete(c context.Context, uid uuid.UUID) error
}

// OrgRepository defines methods the service layer expects
// any repository it interacts with to implement
type OrgRepository interface {
	FindByID(c context.Context, uid uuid.UUID) (*models.Organization, error)
	FindAll(c context.Context) ([]*models.Organization, error)
	Update(c context.Context, uid uuid.UUID, user *models.Organization) error
	Delete(c context.Context, uid uuid.UUID) error
}

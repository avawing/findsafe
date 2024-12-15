package interfaces

import (
	"context"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
)

// OrgService defines methods the handler layer expects
// any service it interacts with to implement
type OrgService interface {
	GetAll(c context.Context) ([]*models.Organization, error)
	Get(c context.Context, uid uuid.UUID) (*models.Organization, error)
	GetAllInSearch(c context.Context, uid uuid.UUID) ([]*models.Organization, error)
	Update(c context.Context, uid uuid.UUID, user *models.Organization) error
	Delete(c context.Context, uid uuid.UUID) error
}

// OrgRepository defines methods the service layer expects
// any repository it interacts with to implement
type OrgRepository interface {
	FindOrgByID(c context.Context, uid uuid.UUID) (*models.Organization, error)
	FindAllOrgs(c context.Context) ([]*models.Organization, error)
	FindAllInSearch(c context.Context, uid uuid.UUID) ([]*models.Organization, error)
	UpdateOrg(c context.Context, uid uuid.UUID, user *models.Organization) error
	DeleteOrg(c context.Context, uid uuid.UUID) error
}

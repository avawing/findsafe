package interfaces

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
)

// ResourceService defines methods the handler layer expects
// any service it interacts with to implement
type ResourceService interface {
	Get(c context.Context, uid uuid.UUID) (*models.Resource, error)
	Update(c context.Context, uid uuid.UUID, user *models.Resource) error
	Delete(c context.Context, uid uuid.UUID) error
	GetByOwnerID(c context.Context, ownerID uuid.UUID) ([]*models.Resource, error)
	GetByIssuedID(c context.Context, issuedID uuid.UUID) ([]*models.Resource, error)
	GetByTeamID(c context.Context, teamID uuid.UUID) ([]*models.Resource, error)
	GetByAvailable(c context.Context, searchID uuid.UUID) ([]*models.Resource, error)
	GetUnreturned(c context.Context, searchID uuid.UUID) ([]*models.Resource, error)
}

// ResourceRepository defines methods the service layer expects
// any repository it interacts with to implement
type ResourceRepository interface {
	FindResourceByID(c context.Context, uid uuid.UUID) (*models.Resource, error)
	UpdateResource(c context.Context, uid uuid.UUID, user *models.Resource) error
	DeleteResource(c context.Context, uid uuid.UUID) error
	FindByOwnerID(c context.Context, ownerID uuid.UUID) ([]*models.Resource, error)
	FindByIssuedID(c context.Context, issuedID uuid.UUID) ([]*models.Resource, error)
	FindByTeamID(c context.Context, teamID uuid.UUID) ([]*models.Resource, error)
	FindAvailable(c context.Context, searchID uuid.UUID) ([]*models.Resource, error)
	FindUnreturned(c context.Context, uid uuid.UUID) ([]*models.Resource, error)
}

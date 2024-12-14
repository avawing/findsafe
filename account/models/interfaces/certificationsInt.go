package interfaces

import (
	"context"
	"findsafe/account/models/models"
	"github.com/google/uuid"
)

// CertService defines methods the handler layer expects
// any service it interacts with to implement
type CertService interface {
	Get(c context.Context, uid uuid.UUID) (*models.Certification, error)
	Update(c context.Context, uid uuid.UUID, user *models.Certification) error
	Delete(c context.Context, uid uuid.UUID) error
	GetByUserID(c context.Context, uid uuid.UUID) ([]*models.Certification, error)
	GetByAccreditingOrg(c context.Context, org uuid.UUID) ([]*models.Certification, error)
}

// CertRepository defines methods the service layer expects
// any repository it interacts with to implement
type CertRepository interface {
	FindByCertID(c context.Context, uid uuid.UUID) (*models.Certification, error)
	UpdateCert(c context.Context, uid uuid.UUID, user *models.Certification) error
	DeleteByCertID(c context.Context, uid uuid.UUID) error
	FindByUserID(c context.Context, uid uuid.UUID) ([]*models.Certification, error)
	FindByAccreditingOrg(c context.Context, org uuid.UUID) ([]*models.Certification, error)
}

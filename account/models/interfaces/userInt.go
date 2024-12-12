package interfaces

import (
	"context"
	"findsafe/account/models"
	"github.com/google/uuid"
)

// UserService defines methods the handler layer expects
// any service it interacts with to implement
type UserService interface {
	Get(c context.Context, uid uuid.UUID) (*models.User, error)
	Update(c context.Context, uid uuid.UUID, user *models.User) error
	Delete(c context.Context, uid uuid.UUID) error
}

// UserRepository defines methods the service layer expects
// any repository it interacts with to implement
type UserRepository interface {
	FindByID(c context.Context, uid uuid.UUID) (*models.User, error)
	Update(c context.Context, uid uuid.UUID, user *models.User) error
	Delete(c context.Context, uid uuid.UUID) error
}

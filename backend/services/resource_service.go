package services

import (
	"context"
	"findsafe/backend/models/interfaces"
	"findsafe/backend/models/models"
	"github.com/google/uuid"
)

// ResourceService acts as a struct for injecting an implementation of UserRepository
// for use in service methods.
type ResourceService struct {
	ResourceRepository interfaces.ResourceRepository
}

// NewResourceService is a factory function for initializing a UserService with
// its repository layer dependencies.
func NewResourceService(c *USConfig) *ResourceService {
	return &ResourceService{
		ResourceRepository: c.ResourceRepository,
	}
}

func (r *ResourceService) Get(c context.Context, uid uuid.UUID) (*models.Resource, error) {
	return r.ResourceRepository.FindResourceByID(c, uid)
}
func (r *ResourceService) Update(c context.Context, uid uuid.UUID, resource *models.Resource) error {
	return r.ResourceRepository.UpdateResource(c, uid, resource)
}
func (r *ResourceService) Delete(c context.Context, uid uuid.UUID) error {
	return r.ResourceRepository.DeleteResource(c, uid)
}
func (r *ResourceService) GetByOwnerID(c context.Context, ownerID uuid.UUID) ([]*models.Resource, error) {
	return r.ResourceRepository.FindByOwnerID(c, ownerID)
}
func (r *ResourceService) GetByIssuedID(c context.Context, issuedID uuid.UUID) ([]*models.Resource, error) {
	return r.ResourceRepository.FindByIssuedID(c, issuedID)
}
func (r *ResourceService) GetByTeamID(c context.Context, teamID uuid.UUID) ([]*models.Resource, error) {
	return r.ResourceRepository.FindByTeamID(c, teamID)
}
func (r *ResourceService) GetByAvailable(c context.Context, searchID uuid.UUID) ([]*models.Resource, error) {
	return r.ResourceRepository.FindAvailable(c, searchID)
}
func (r *ResourceService) GetUnreturned(c context.Context, searchID uuid.UUID) ([]*models.Resource, error) {
	return r.ResourceRepository.FindUnreturned(c, searchID)
}

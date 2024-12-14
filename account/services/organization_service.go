package services

import (
	"context"
	"findsafe/account/models"
	"findsafe/account/models/interfaces"
	"github.com/google/uuid"
)

// OrgService acts as a struct for injecting an implementation of OrgRepository
// for use in service methods.
type OrgService struct {
	OrgRepository interfaces.OrgRepository
}

func (o *OrgService) Get(c context.Context, uid uuid.UUID) (*models.Organization, error) {
	return o.OrgRepository.FindByID(c, uid)
}
func (o *OrgService) Update(c context.Context, uid uuid.UUID, user *models.Organization) error {
	return o.OrgRepository.Update(c, uid, user)
}
func (o *OrgService) Delete(c context.Context, uid uuid.UUID) error {
	return o.OrgRepository.Delete(c, uid)
}
func (o *OrgService) GetAll(c context.Context) ([]*models.Organization, error) {
	return o.OrgRepository.FindAll(c)
}

func (o *OrgService) GetAllInSearch(c context.Context, uid uuid.UUID) ([]*models.Organization, error) {
	return o.OrgRepository.FindAllInSearch(c, uid)
}

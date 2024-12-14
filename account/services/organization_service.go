package services

import (
	"context"
	"findsafe/account/models/interfaces"
	"findsafe/account/models/models"
	"github.com/google/uuid"
)

// OrgService acts as a struct for injecting an implementation of OrgRepository
// for use in service methods.
type OrgService struct {
	OrgRepository interfaces.OrgRepository
}

func (o *OrgService) Get(c context.Context, uid uuid.UUID) (*models.Organization, error) {
	return o.OrgRepository.FindOrgByID(c, uid)
}
func (o *OrgService) Update(c context.Context, uid uuid.UUID, user *models.Organization) error {
	return o.OrgRepository.UpdateOrg(c, uid, user)
}
func (o *OrgService) Delete(c context.Context, uid uuid.UUID) error {
	return o.OrgRepository.DeleteOrg(c, uid)
}
func (o *OrgService) GetAll(c context.Context) ([]*models.Organization, error) {
	return o.OrgRepository.FindAllOrgs(c)
}

func (o *OrgService) GetAllInSearch(c context.Context, uid uuid.UUID) ([]*models.Organization, error) {
	return o.OrgRepository.FindAllInSearch(c, uid)
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Organization struct {
	gorm.Model
	ID                    uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name                  string
	PhoneNumber           string
	DateFounded           time.Time
	About                 string
	PrimaryContactID      uuid.UUID `gorm:"type:uuid"`
	PrimaryContact        User      `gorm:"foreignKey:PrimaryContactID"`
	Website               string
	Members               []*User          `gorm:"many2many:user_organizations;"`
	ActiveSearches        []*Searches      `gorm:"many2many:searches_organizations;"`
	CertificationsOffered []*Certification `gorm:"foreignKey:OfferedByOrgID"`
	Resources             []*Resource      `gorm:"foreignKey:OwnerOrganizationID"`
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Organization struct {
	gorm.Model
	ID                    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name                  string
	PhoneNumber           string
	DateFounded           time.Time
	About                 string
	PrimaryContactID      uuid.UUID `gorm:"type:uuid"`
	PrimaryContact        User      `gorm:"foreignKey:PrimaryContactID"`
	Website               string
	Members               []*User         `gorm:"many2many:user_languages;"`
	ActiveSearches        []*uuid.UUID    // many2many searches
	CertificationsOffered []Certification `gorm:"foreignKey:OfferedByOrgID"`
}

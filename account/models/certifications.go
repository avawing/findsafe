package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Certification struct {
	gorm.Model
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name           string    `gorm:"not null"` // TODO: enums provided by orgs ?
	FileLocation   string
	DateGranted    time.Time `gorm:"not null"`
	ExpirationDate time.Time `gorm:"not null"`
	UserID         uuid.UUID
	User           User      `gorm:"notnull;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AccreditingOrg uuid.UUID // TODO: Certifications Offered by Organizations (I.E. NASAR)
}

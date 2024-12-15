package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Certification struct {
	gorm.Model
	ID               uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name             string
	FileLocation     string
	DateGranted      time.Time
	ExpirationDate   time.Time
	UserID           uuid.UUID
	User             User         `gorm:"foreignKey:UserID"`
	AccreditingOrgID uuid.UUID    `gorm:"type:uuid"`                                                                  // Add this field to store the foreign key
	AccreditingOrg   Organization `gorm:"foreignKey:AccreditingOrgID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Define the foreign key constraint
	OfferedByOrgID   uuid.UUID
	OfferedByOrg     Organization `gorm:"foreignKey:OfferedByOrgID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

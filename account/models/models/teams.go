package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	ID             uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name           string
	CurrentLat     string
	CurrentLng     string
	TeamLeadID     *uuid.UUID `gorm:"type:uuid;index"`
	ActiveSortie   string
	ActiveSearchID *uuid.UUID `gorm:"type:uuid;index"`
	TeamLead       User       `gorm:"foreignKey:TeamLeadID;constraint:OnDelete:SET NULL;OnUpdate:CASCADE"`
}

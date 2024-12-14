package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name           string    `gorm:"unique,not null"`
	CurrentLat     *string
	CurrentLng     *string
	TeamLeadID     uuid.UUID
	TeamLead       User `gorm:"foreignKey:TeamLeadID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ActiveSortie   *uuid.UUID
	ActiveSearchID uuid.UUID
	TeamMembers    []User     `gorm:"foreignKey:ActiveTeamID"`
	TeamResources  []Resource `gorm:"foreignKey:IssuedToTeamID"`
}

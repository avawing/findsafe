package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Resource struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Name                string
	OwnerID             uuid.UUID
	Owner               User `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OwnerOrganizationID *uuid.UUID
	OwnerOrganization   Organization `gorm:"foreignKey:OwnerOrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key to Organization
	IssuedToUserID      *uuid.UUID
	IssuedToUser        User `gorm:"foreignKey:IssuedToUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IssuedToTeamID      *uuid.UUID
	IssuedToTeam        Team `gorm:"foreignKey:IssuedToTeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IssuedAt            time.Time
	ReturnedAt          time.Time
	ActiveSearchID      *uuid.UUID
	ActiveSearch        Searches `gorm:"foreignKey:ActiveSearchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ImageFileLocation   string   // AWS S3 Bucket
	ResourceType        string   //TODO: enums of resource types
	Information         string
	DesignatedPurpose   string
}

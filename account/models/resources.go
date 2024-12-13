package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Resource struct {
	gorm.Model
	ID                uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name              string
	OwnerID           uuid.UUID
	Owner             User `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IssuedToUserID    *uuid.UUID
	IssuedToUser      User `gorm:"foreignKey:IssuedToUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IssuedToTeamID    *uuid.UUID
	IssuedToTeam      User `gorm:"foreignKey:IssuedToTeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IssuedAt          time.Time
	ReturnedAt        time.Time
	ImageFileLocation string // AWS S3 Bucket
	ResourceType      string //TODO: enums of resource types
	Information       string
	DesignatedPurpose string
}
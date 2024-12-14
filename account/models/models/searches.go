package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Searches struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Subjects           uuid.UUID // Foreign Key to Subjects
	BaseAddress        string
	BaseCity           string
	BaseState          string
	BasePostCode       string
	Lat                string
	Lon                string
	SearchResult       string //TODO: enums
	OrganizationID     uuid.UUID
	Organization       Organization
	PointOfContactID   uuid.UUID
	PointOfContact     User
	Internet           string
	InternetAccess     string
	MapStorageLocation string
	StartDate          time.Time
	EndDate            time.Time
	Organizations      []*Organization `gorm:"many2many:searches_organizations;"`
	Resources          []*Resource     `gorm:"foreignKey:OwnerOrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	//Evidence
}

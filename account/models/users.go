package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                uuid.UUID       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName         string          `gorm:"not null, type:varchar(100)" json:"first_name"`
	LastName          string          `gorm:"not null, type:varchar(100)" json:"last_name"`
	City              string          `json:"city"`
	State             string          `json:"state"`
	Email             string          `gorm:"not null, type:varchar(100)" json:"email"`
	Phone             string          `gorm:"not null, type:varchar(20)" json:"phone"`
	Certifications    []Certification `gorm:"foreignKey:UserID"`
	Resources         []Resource      `gorm:"foreignKey:OwnerID"`
	AssignedResources []Resource      `gorm:"foreignKey:IssuedToUserID"`
	// Foreign Key
	ActiveSearchID *uuid.UUID `json:"active_search_id"`
	ActiveSortieID *uuid.UUID `json:"active_sortie_id"`
	ActiveTeamID   *uuid.UUID `json:"active_team_id"`
	Team           *Team      `gorm:"foreignKey:ActiveTeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TeamRole       *string    `json:"team_role"`

	//Searches       int
	//Sorties        int
	//Evidence       int
	//Resources      int
}

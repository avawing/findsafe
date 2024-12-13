package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName string    `gorm:"not null, type:varchar(100)" json:"first_name"`
	LastName  string    `gorm:"not null, type:varchar(100)" json:"last_name"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Email     string    `gorm:"not null, type:varchar(100)" json:"email"`
	Phone     string    `gorm:"not null, type:varchar(20)" json:"phone"`
	// Foreign Key
	ActiveSearchID *uuid.UUID `json:"active_search"`
	ActiveSortieID *uuid.UUID `json:"active_sortie"`
	ActiveTeamID   *uuid.UUID `json:"active_team"`
	Team           *Team      `gorm:"foreignKey:ActiveTeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	//Searches       int
	//Sorties        int
	//Evidence       int
	//Resources      int
	//Certifications int
}

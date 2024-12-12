package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null, type:varchar(100)" json:"first_name"`
	LastName  string `gorm:"not null, type:varchar(100)" json:"last_name"`
	City      string `json:"city"`
	State     string `json:"state"`
	// Foreign Key
	ActiveSearch *uuid.UUID `json:"active_search"`
	//Searches       int
	//Sorties        int
	//Evidence       int
	//Resources      int
	//Certifications int
}

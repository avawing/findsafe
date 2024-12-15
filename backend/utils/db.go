package utils

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	if db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

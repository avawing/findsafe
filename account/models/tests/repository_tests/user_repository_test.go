package repository_tests

import (
	"context"
	"findsafe/account/models/models"
	"findsafe/account/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"os/exec"
	"testing"
)

var gdb *gorm.DB

func dbSetup() *repository.Repository {
	return &repository.Repository{
		DB: gdb.Debug().Begin(),
	}
}

func TestMain(m *testing.M) {
	var err error
	// database file name
	dbName := "database_test.db"
	// remove old database
	exec.Command("rm", "-f", dbName)

	// open and create a new database
	gdb, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// migrate tables
	gdb.AutoMigrate(
		models.Certification{},
		models.Organization{},
		models.Resource{},
		models.Searches{},
		models.Team{},
		models.User{},
	)
	//gdb.AutoMigrate(&schema.Rating{})
	//gdb.AutoMigrate(&schema.Dish{})
	//
	//// add mock data
	//gdb.Create(&users)
	//gdb.Create(&dishes)

	// run tests
	os.Exit(m.Run())
}

func TestGetUsers(t *testing.T) {
	db := dbSetup()
	t.Run("Success - Get User By Nickname", func(t *testing.T) {
		u, err := db.FindByUserID(context.Background(), uuid.New())
		assert.Nil(t, err)
		assert.NotNil(t, u)

	})
	//t.Run("Fail - Get User By Nickname", func(t *testing.T) {
	//	u, err := db.GetUserByUsername("Drogo")
	//	assert.Equal(t, gorm.ErrRecordNotFound, err)
	//	assert.Equal(t, (*schema.User)(nil), u)
	//})
}

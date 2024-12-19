package repository_tests

import (
	"context"
	"findsafe/backend/models/models"
	"findsafe/backend/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

var gdb *gorm.DB

func dbSetup() *repository.Repository {
	return &repository.Repository{
		DB: gdb.Debug().Begin(),
	}
}

var orgsuuids = []uuid.UUID{
	uuid.Must(uuid.NewRandom()),
}

var searchuuids = []uuid.UUID{
	uuid.Must(uuid.NewRandom()),
	uuid.Must(uuid.NewRandom()),
}

var teamuuids = []uuid.UUID{
	uuid.Must(uuid.NewRandom()),
	uuid.Must(uuid.NewRandom()),
}

var resources = []models.Resource{
	{
		ID:                  uuid.MustParse("daefd7cd-0b82-403f-babe-871548f7e976"),
		Name:                "Resource A",
		OwnerID:             users[0].ID,                                             // Owner: Alice
		Owner:               users[0],                                                // Alice is the owner
		OwnerOrganizationID: nil,                                                     // No organization for now
		OwnerOrganization:   models.Organization{},                                   // Empty organization (can be filled if needed)
		IssuedToUserID:      &users[1].ID,                                            // Issued to Bob
		IssuedToTeamID:      nil,                                                     // No team assigned
		IssuedToTeam:        models.Team{},                                           // No team user associated
		IssuedAt:            time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC), // Fixed timestamp
		ReturnedAt:          time.Time{},                                             // Empty time (not returned yet)
		ActiveSearchID:      nil,                                                     // Not associated with any search
		ActiveSearch:        models.Searches{},                                       // Empty search
		ImageFileLocation:   "s3://bucket/resourceA",                                 // Example image file location
		ResourceType:        "TypeA",                                                 // Specify the resource type
		Information:         "Sample resource A",                                     // Sample information
		DesignatedPurpose:   "Testing",                                               // Designated purpose for the resource
	},
	{
		ID:                  uuid.MustParse("6352951c-1fcb-4351-a2b3-acdc264f18b3"),
		Name:                "Resource B",
		OwnerID:             users[2].ID,                                             // Owner: Charlie
		OwnerOrganizationID: nil,                                                     // No organization for now
		OwnerOrganization:   models.Organization{},                                   // Empty organization (can be filled if needed)
		IssuedToUserID:      &users[3].ID,                                            // Issued to David
		IssuedToTeamID:      nil,                                                     // No team assigned
		IssuedToTeam:        models.Team{},                                           // No team user associated
		IssuedAt:            time.Date(2024, time.December, 2, 0, 0, 0, 0, time.UTC), // Fixed timestamp
		ReturnedAt:          time.Time{},                                             // Empty time (not returned yet)
		ActiveSearchID:      nil,                                                     // Not associated with any search
		ActiveSearch:        models.Searches{},                                       // Empty search
		ImageFileLocation:   "s3://bucket/resourceB",                                 // Example image file location
		ResourceType:        "TypeB",                                                 // Specify the resource type
		Information:         "Sample resource B",                                     // Sample information
		DesignatedPurpose:   "Analysis",                                              // Designated purpose for the resource
	},
}

// Create user instances with unique IDs
var users = []models.User{
	{
		ID:        uuid.MustParse("daefd7cd-0b82-403f-babe-871548f7e976"),
		FirstName: "Alice",
		LastName:  "Smith",
		City:      "New York",
		State:     "NY",
		Email:     "alice.smith@example.com",
		Phone:     "+1-234-567-8901",
	},
	{
		ID:        uuid.MustParse("6352951c-1fcb-4351-a2b3-acdc264f18b3"),
		FirstName: "Bob",
		LastName:  "Johnson",
		City:      "Los Angeles",
		State:     "CA",
		Email:     "bob.johnson@example.com",
		Phone:     "+1-234-567-8902",
	},
	{
		ID:        uuid.MustParse("d3367954-e5a8-4df4-821a-640b4d6d4360"),
		FirstName: "Charlie",
		LastName:  "Davis",
		City:      "Chicago",
		State:     "IL",
		Email:     "charlie.davis@example.com",
		Phone:     "+1-234-567-8903",
	},
	{
		ID:        uuid.MustParse("9847966c-8f21-45fa-b72d-a40206ba5199"),
		FirstName: "David",
		LastName:  "Miller",
		City:      "San Francisco",
		State:     "CA",
		Email:     "david.miller@example.com",
		Phone:     "+1-234-567-8904",
	},
}

// Create search instances with unique IDs, assigning users as point of contact
var searches = []models.Searches{
	{
		ID:                 searchuuids[0],
		Subjects:           uuid.Must(uuid.NewRandom()), // Subject UUID, you can link it to another model if necessary
		BaseAddress:        "123 Main St",
		BaseCity:           "New York",
		BaseState:          "NY",
		BasePostCode:       "10001",
		Lat:                "40.7128",
		Lon:                "-74.0060",
		SearchResult:       "Success",
		OrganizationID:     orgsuuids[0], // Assuming you have an org UUID to link here
		Organization:       models.Organization{ID: orgsuuids[0]},
		PointOfContactID:   users[0].ID, // Alice is the point of contact
		Internet:           "www.example.com",
		InternetAccess:     "Wi-Fi",
		MapStorageLocation: "s3://maps/location1",
		StartDate:          time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC),
		EndDate:            time.Date(2024, time.December, 5, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:                 searchuuids[1],
		Subjects:           uuid.Must(uuid.NewRandom()), // Subject UUID
		BaseAddress:        "456 Elm St",
		BaseCity:           "Los Angeles",
		BaseState:          "CA",
		BasePostCode:       "90001",
		Lat:                "34.0522",
		Lon:                "-118.2437",
		SearchResult:       "Failed",
		OrganizationID:     orgsuuids[0],
		Organization:       models.Organization{ID: orgsuuids[0]},
		PointOfContactID:   users[1].ID, // Bob is the point of contact
		Internet:           "www.anotherexample.com",
		InternetAccess:     "Ethernet",
		MapStorageLocation: "s3://maps/location2",
		StartDate:          time.Date(2024, time.December, 10, 0, 0, 0, 0, time.UTC),
		EndDate:            time.Date(2024, time.December, 15, 0, 0, 0, 0, time.UTC),
	},
}

// Create teams with unique IDs and associate users as team leads
var teams = []models.Team{
	{
		ID:           teamuuids[0],
		Name:         "Alpha Team",
		CurrentLat:   "37.7749",
		CurrentLng:   "-122.4194",
		TeamLeadID:   &users[0].ID, // Alice is the team lead
		ActiveSortie: "Sortie-001",
		TeamLead:     users[0],
	},
	{
		ID:           teamuuids[1],
		Name:         "Bravo Team",
		CurrentLat:   "34.0522",
		CurrentLng:   "-118.2437",
		TeamLeadID:   &users[2].ID, // Charlie is the team lead
		ActiveSortie: "Sortie-002",
		TeamLead:     users[2],
	},
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

	gdb.Create(&searches)
	gdb.Create(&users)
	gdb.Create(&teams)
	gdb.Create(&resources)
	// run tests
	os.Exit(m.Run())
}

// Resources
func TestGetResources(t *testing.T) {
	db := dbSetup()
	t.Run("Fail - Resources Not Found", func(t *testing.T) {
		u, err := db.FindResourceByID(context.Background(), uuid.New())
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Equal(t, u, &models.Resource{})
	})
	t.Run("Success - Get Resources By ID", func(t *testing.T) {
		u, err := db.FindResourceByID(context.Background(), resources[0].ID)
		assert.Equal(t, nil, err)
		assert.Equal(t, resources[0].ID, u.ID)
	})
}

// Searches
func TestGetSearches(t *testing.T) {
	db := dbSetup()
	t.Run("Fail - Search Not Found", func(t *testing.T) {
		u, err := db.FindBySearchID(context.Background(), uuid.New())
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Equal(t, u, &models.Searches{})
	})
	t.Run("Success - Get Search By ID", func(t *testing.T) {
		u, err := db.FindBySearchID(context.Background(), searches[0].ID)
		assert.Equal(t, nil, err)
		assert.Equal(t, searches[0].ID, u.ID)
	})
}

// Teams
func TestGetTeams(t *testing.T) {
	db := dbSetup()
	t.Run("Fail - Team Not Found", func(t *testing.T) {
		u, err := db.FindTeamByID(context.Background(), uuid.New())
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Equal(t, u, &models.Team{})
	})
	t.Run("Success - Get Team By ID", func(t *testing.T) {
		u, err := db.FindTeamByID(context.Background(), teams[0].ID)
		assert.Equal(t, nil, err)
		assert.Equal(t, teams[0].ID, u.ID)
	})
}

// Users
func TestGetUsers(t *testing.T) {
	db := dbSetup()
	t.Run("Fail - User Not Found", func(t *testing.T) {
		u, err := db.FindUserByID(context.Background(), uuid.New())
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Equal(t, u, &models.User{})
	})
	t.Run("Success - Get User By ID", func(t *testing.T) {
		u, err := db.FindUserByID(context.Background(), users[0].ID)
		assert.Equal(t, nil, err)
		assert.Equal(t, users[0].ID, u.ID)
	})
}

// Certifications
// Organizations

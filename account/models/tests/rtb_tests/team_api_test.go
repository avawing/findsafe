package tests

import (
	"encoding/json"
	handlers "findsafe/account/handlers/rtb_api"
	"findsafe/account/models"
	"findsafe/account/models/apperrors"
	mocks2 "findsafe/account/models/mocks/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func TestGetTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockTeamResp := &models.Team{
			Model:           gorm.Model{},
			ID:              uid,
			Name:            "ALPHA_TEAM",
			CurrentLat:      nil,
			CurrentLng:      nil,
			TeamLeadID:      uuid.UUID{},
			TeamLead:        models.User{},
			RadioOperatorID: uuid.UUID{},
			RadioOperator:   models.User{},
			MedicID:         uuid.UUID{},
			Medic:           models.User{},
			ActiveSortie:    nil,
			ActiveSearchID:  uuid.UUID{},
		}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockTeamService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockTeamResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"team": mockTeamResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})

	t.Run("BadTeamID", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/%s", "uid.String()"), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		tf.mockTeamService.AssertNotCalled(t, "Get", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockTeamService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(&models.Team{}, fmt.Errorf("downstream error"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("team", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})
}
func TestGetTeams(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockTeamResp := []*models.Team{{
			Model:           gorm.Model{},
			ID:              uid,
			Name:            "ALPHA_TEAM",
			CurrentLat:      nil,
			CurrentLng:      nil,
			TeamLeadID:      uuid.UUID{},
			TeamLead:        models.User{},
			RadioOperatorID: uuid.UUID{},
			RadioOperator:   models.User{},
			MedicID:         uuid.UUID{},
			Medic:           models.User{},
			ActiveSortie:    nil,
			ActiveSearchID:  uuid.UUID{},
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockTeamService.On("GetAllinSearch", mock.AnythingOfType("*gin.Context"), uid).Return(mockTeamResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/search/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"teams": mockTeamResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})

	t.Run("BadTeamID", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/search/%s", "uid.String()"), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		tf.mockTeamService.AssertNotCalled(t, "GetAllinSearch", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockTeamService.On("GetAllinSearch", mock.AnythingOfType("*gin.Context"), uid).Return([]*models.Team{}, fmt.Errorf("downstream error"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/search/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("search id", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})
}
func TestGetUnassignedTeams(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockTeamResp := []*models.Team{{
			Model:           gorm.Model{},
			ID:              uid,
			Name:            "ALPHA_TEAM",
			CurrentLat:      nil,
			CurrentLng:      nil,
			TeamLeadID:      uuid.UUID{},
			TeamLead:        models.User{},
			RadioOperatorID: uuid.UUID{},
			RadioOperator:   models.User{},
			MedicID:         uuid.UUID{},
			Medic:           models.User{},
			ActiveSortie:    nil,
			ActiveSearchID:  uuid.UUID{},
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockTeamService.On("GetAllWithoutSortie", mock.AnythingOfType("*gin.Context")).Return(mockTeamResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/search/%s/unassigned", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"teams": mockTeamResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockTeamService.On("GetAllWithoutSortie", mock.AnythingOfType("*gin.Context")).Return([]*models.Team{}, fmt.Errorf("downstream error"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/search/%s/unassigned", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("teams", "unassigned")

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})
}
func TestGetSortieTeams(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockTeamResp := &models.Team{
			Model:           gorm.Model{},
			ID:              uid,
			Name:            "ALPHA_TEAM",
			CurrentLat:      nil,
			CurrentLng:      nil,
			TeamLeadID:      uuid.UUID{},
			TeamLead:        models.User{},
			RadioOperatorID: uuid.UUID{},
			RadioOperator:   models.User{},
			MedicID:         uuid.UUID{},
			Medic:           models.User{},
			ActiveSortie:    nil,
			ActiveSearchID:  uuid.UUID{},
		}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockTeamService.On("GetBySortie", mock.AnythingOfType("*gin.Context"), uid).Return(mockTeamResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/sortie/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"team": mockTeamResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})

	t.Run("BadTeamID", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/sortie/%s", "uid.String()"), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		tf.mockTeamService.AssertNotCalled(t, "GetBySortie", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockTeamService.On("GetBySortie", mock.AnythingOfType("*gin.Context"), uid).Return(&models.Team{}, fmt.Errorf("downstream error"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/teams/sortie/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("sortie id", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})
}

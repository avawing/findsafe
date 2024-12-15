package tests

import (
	"encoding/json"
	handlers "findsafe/backend/handlers/rtb_api"
	"findsafe/backend/models/apperrors"
	mocks2 "findsafe/backend/models/mocks/services"
	models2 "findsafe/backend/models/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func TestGetCertifications(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockTeamResp := []*models2.Certification{{
			Model: gorm.Model{},
			ID:    uid,
			Name:  "TEST_CERTIFICATE",
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
			mockCertService: new(mocks2.MockCertService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockCertService.On("GetByUserID", mock.AnythingOfType("*gin.Context"), uid).Return(mockTeamResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("certifications", &models2.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
			CertService: tf.mockCertService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s/certifications", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"certifications": mockTeamResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockCertService.AssertExpectations(t)
	})

	t.Run("BadTeamID", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
			mockCertService: new(mocks2.MockCertService),
		}
		tf.setup()

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models2.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
			CertService: tf.mockCertService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s/certifications", "uid.String()"), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		tf.mockCertService.AssertNotCalled(t, "GetByUserID", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
			mockCertService: new(mocks2.MockCertService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockCertService.On("GetByUserID", mock.AnythingOfType("*gin.Context"), uid).Return([]*models2.Certification{}, fmt.Errorf("some error"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models2.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
			CertService: tf.mockCertService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s/certifications", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("user", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockCertService.AssertExpectations(t)
	})
}
func TestGetCertification(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockTeamResp := &models2.Certification{
			Model: gorm.Model{},
			ID:    uid,
			Name:  "TEST_CERTIFICATE",
		}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
			mockCertService: new(mocks2.MockCertService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockCertService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockTeamResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("certifications", &models2.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
			CertService: tf.mockCertService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/certifications/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"certifications": mockTeamResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockCertService.AssertExpectations(t)
	})

	t.Run("BadTeamID", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
			mockCertService: new(mocks2.MockCertService),
		}
		tf.setup()

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models2.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
			CertService: tf.mockCertService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/certifications/%s", "uid.String()"), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		tf.mockCertService.AssertNotCalled(t, "Get", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
			mockCertService: new(mocks2.MockCertService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockCertService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(&models2.Certification{}, fmt.Errorf("some error"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models2.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
			CertService: tf.mockCertService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/certifications/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("user", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockCertService.AssertExpectations(t)
	})
}

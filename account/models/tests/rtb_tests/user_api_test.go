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
	"net/http/httptest"
	"testing"
)

type testFixture struct {
	router              *gin.Engine
	rr                  *httptest.ResponseRecorder
	mockUserService     *mocks2.MockUserService
	mockTeamService     *mocks2.MockTeamService
	mockCertService     *mocks2.MockCertService
	mockResourceService *mocks2.MockResourceService
}

func (tf *testFixture) setup() {
	tf.router = gin.Default()
	tf.rr = httptest.NewRecorder()
}

func TestMe(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := &models.User{
			Model:     gorm.Model{},
			ID:        uid,
			FirstName: "TEST_USER_FIRST_NAME",
			LastName:  "TEST_USER_LAST_NAME",
			City:      "TEST_CITY",
			State:     "CA",
		}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/users/me", nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"user": mockUserResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, 200, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})

	t.Run("NoContextUser", func(t *testing.T) {
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
			mockTeamService: new(mocks2.MockTeamService),
		}
		tf.setup()

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
			TeamService: tf.mockTeamService,
		})

		request, err := http.NewRequest(http.MethodGet, "/users/me", nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.Equal(t, http.StatusInternalServerError, tf.rr.Code)
		tf.mockUserService.AssertNotCalled(t, "Get", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("Get", mock.Anything, uid).Return(&models.User{}, fmt.Errorf("some error down call chain"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/users/me", nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("user", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})
}

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := &models.User{
			Model:     gorm.Model{},
			ID:        uid,
			FirstName: "TEST_USER_FIRST_NAME",
			LastName:  "TEST_USER_LAST_NAME",
			City:      "TEST_CITY",
			State:     "CA",
		}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"user": mockUserResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})

	t.Run("BadUserID", func(t *testing.T) {
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("Get", mock.Anything, mock.Anything).Return(nil, nil)

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", "uid.String()"), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		tf.mockUserService.AssertNotCalled(t, "Get", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("Get", mock.Anything, uid).Return(&models.User{}, fmt.Errorf("some error down call chain"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("user", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})
}

func TestGetUsersBySearch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := []*models.User{{
			Model:     gorm.Model{},
			ID:        uid,
			FirstName: "TEST_USER_FIRST_NAME",
			LastName:  "TEST_USER_LAST_NAME",
			City:      "TEST_CITY",
			State:     "CA",
			Email:     "TEST_USER@TEST_EMAIL.COM",
			Phone:     "1-555-555-5555",
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("GetAllinSearch", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/search/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"users": mockUserResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})

	t.Run("NoContextUser", func(t *testing.T) {
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("GetAllinSearch", mock.Anything, mock.Anything).Return(nil, nil)

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/search/%s", "uid.String()"), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		tf.mockUserService.AssertNotCalled(t, "GetAllinSearch", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("GetAllinSearch", mock.Anything, uid).Return([]*models.User{}, fmt.Errorf("some error down call chain"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/search/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("user", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})
}

func TestGetUsersByOrg(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := []*models.User{{
			Model:     gorm.Model{},
			ID:        uid,
			FirstName: "TEST_USER_FIRST_NAME",
			LastName:  "TEST_USER_LAST_NAME",
			City:      "TEST_CITY",
			State:     "CA",
			Email:     "TEST_USER@TEST_EMAIL.COM",
			Phone:     "1-555-555-5555",
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("GetAllInOrg", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/organization/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"users": mockUserResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})

	t.Run("NoContextUser", func(t *testing.T) {
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("GetAllInOrg", mock.Anything, mock.Anything).Return(nil, nil)

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/search/%s", "uid.String()"), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		tf.mockUserService.AssertNotCalled(t, "GetAllInOrg", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("GetAllInOrg", mock.Anything, uid).Return([]*models.User{}, fmt.Errorf("some error down call chain"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/organization/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("user", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})
}

func TestGetUsersBySortie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := []*models.User{{
			Model:     gorm.Model{},
			ID:        uid,
			FirstName: "TEST_USER_FIRST_NAME",
			LastName:  "TEST_USER_LAST_NAME",
			City:      "TEST_CITY",
			State:     "CA",
			Email:     "TEST_USER@TEST_EMAIL.COM",
			Phone:     "1-555-555-5555",
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("GetAllInSortie", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/sortie/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respBody, err := json.Marshal(gin.H{
			"users": mockUserResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})

	t.Run("BadUUID", func(t *testing.T) {
		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("GetAllInSortie", mock.Anything, mock.Anything).Return(nil, nil)

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/sortie/%s", "uid.String()"), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		tf.mockUserService.AssertNotCalled(t, "GetAllInSortie", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		tf := &testFixture{
			mockUserService: new(mocks2.MockUserService),
		}
		tf.setup()

		// Use the mock directly from the fixture
		tf.mockUserService.On("GetAllInSortie", mock.Anything, uid).Return([]*models.User{}, fmt.Errorf("some error down call chain"))

		tf.router.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				ID: uid,
			})
		})

		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			UserService: tf.mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/sortie/%s", uid.String()), nil)
		assert.NoError(t, err)

		tf.router.ServeHTTP(tf.rr, request)

		respErr := apperrors.NewNotFound("user", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockUserService.AssertExpectations(t)
	})
}

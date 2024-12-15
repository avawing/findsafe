package tests

import (
	"encoding/json"
	handlers "findsafe/backend/handlers/rtb_api"
	"findsafe/backend/models/apperrors"
	mocks2 "findsafe/backend/models/mocks/services"
	"findsafe/backend/models/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestGetUserResources(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockResourceResp := []*models.Resource{{
			ID:   uid,
			Name: "Test Resource",
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockResourceService: new(mocks2.MockResourceService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:               tf.router,
			ResourceService: tf.mockResourceService,
		})
		// Use the mock directly from the fixture
		tf.mockResourceService.On("GetByOwnerID", mock.Anything, uid).Return(mockResourceResp, nil)

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s/resources", uid.String()), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected response body
		respBody, err := json.Marshal(gin.H{
			"users": mockResourceResp,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockResourceService.AssertExpectations(t)
	})

	t.Run("BadRequest", func(t *testing.T) {
		// Test with invalid UUID
		invalidID := "invalid-uuid"

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockResourceService: new(mocks2.MockResourceService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:               tf.router,
			ResourceService: tf.mockResourceService,
		})
		// Create request with invalid UUID
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s/resources", invalidID), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewBadRequest("invalid user id")
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockResourceService.AssertNotCalled(t, "GetByOwnerID", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockResourceService: new(mocks2.MockResourceService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:               tf.router,
			ResourceService: tf.mockResourceService,
		})
		// Setup mock to return an error
		tf.mockResourceService.On("GetByOwnerID", mock.Anything, uid).Return([]*models.Resource{}, fmt.Errorf("not found"))

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s/resources", uid.String()), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewNotFound("user", uid.String())
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusNotFound, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockResourceService.AssertExpectations(t)
	})
}
func TestGetResource(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockResourceResp := &models.Resource{
			ID:   uid,
			Name: "Test Resource",
		}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockResourceService: new(mocks2.MockResourceService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:               tf.router,
			ResourceService: tf.mockResourceService,
		})
		// Use the mock directly from the fixture
		tf.mockResourceService.On("Get", mock.Anything, uid).Return(mockResourceResp, nil)

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/resources/%s", uid.String()), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected response body
		respBody, err := json.Marshal(gin.H{
			"users": mockResourceResp,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockResourceService.AssertExpectations(t)
	})

	t.Run("BadRequest", func(t *testing.T) {
		// Test with invalid UUID
		invalidID := "invalid-uuid"

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockResourceService: new(mocks2.MockResourceService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:               tf.router,
			ResourceService: tf.mockResourceService,
		})
		// Create request with invalid UUID
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/resources/%s", invalidID), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewBadRequest("invalid user id")
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockResourceService.AssertNotCalled(t, "Get", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockResourceService: new(mocks2.MockResourceService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:               tf.router,
			ResourceService: tf.mockResourceService,
		})
		// Setup mock to return an error
		tf.mockResourceService.On("Get", mock.Anything, uid).Return(&models.Resource{}, fmt.Errorf("not found"))

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/resources/%s", uid.String()), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewNotFound("user", uid.String())
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusNotFound, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockResourceService.AssertExpectations(t)
	})
}

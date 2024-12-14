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
	"net/http"
	"testing"
)

func TestGetSearch(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockSearchResp := &models.Searches{
			ID: uid,
		}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockSearchService: new(mocks2.MockSearchService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:             tf.router,
			SearchService: tf.mockSearchService,
		})
		// Use the mock directly from the fixture
		tf.mockSearchService.On("Get", mock.Anything, uid).Return(mockSearchResp, nil)

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/search/%s", uid), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected response body
		respBody, err := json.Marshal(gin.H{
			"search": mockSearchResp,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockSearchService.AssertExpectations(t)
	})

	t.Run("BadRequest", func(t *testing.T) {
		// Test with invalid UUID
		invalidID := "invalid-uuid"

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockSearchService: new(mocks2.MockSearchService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:               tf.router,
			ResourceService: tf.mockResourceService,
		})
		// Create request with invalid UUID
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/search/%s", invalidID), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewBadRequest("invalid search id")
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockSearchService.AssertNotCalled(t, "Get", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid := uuid.New()
		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockSearchService: new(mocks2.MockSearchService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:             tf.router,
			SearchService: tf.mockSearchService,
		})
		// Setup mock to return an error
		tf.mockSearchService.On("Get", mock.Anything, uid).Return(&models.Searches{}, fmt.Errorf("not found"))

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/search/%s", uid), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewNotFound("search", uid.String())
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusNotFound, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockSearchService.AssertExpectations(t)
	})
}
func TestGetSearches(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockSearchResp := []*models.Searches{{
			ID: uid,
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockSearchService: new(mocks2.MockSearchService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:             tf.router,
			SearchService: tf.mockSearchService,
		})
		// Use the mock directly from the fixture
		tf.mockSearchService.On("GetAll", mock.Anything).Return(mockSearchResp, nil)

		// Create request
		request, err := http.NewRequest(http.MethodGet, "/search/", nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected response body
		respBody, err := json.Marshal(gin.H{
			"searches": mockSearchResp,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockSearchService.AssertExpectations(t)
	})

	t.Run("NotFound", func(t *testing.T) {
		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockSearchService: new(mocks2.MockSearchService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:             tf.router,
			SearchService: tf.mockSearchService,
		})
		// Setup mock to return an error
		tf.mockSearchService.On("GetAll", mock.Anything).Return([]*models.Searches{}, fmt.Errorf("not found"))

		// Create request
		request, err := http.NewRequest(http.MethodGet, "/search/", nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewNotFound("searches", "not found")
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusNotFound, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockSearchService.AssertExpectations(t)
	})
}

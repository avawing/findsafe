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

//func (h *Handler) GetOrgCertifications(c *gin.Context) {
//	id := c.Param("id")
//	uid, err := uuid.Parse(id)
//	if err != nil {
//		log.Printf("Unable to generate UUID from request: %v\n", err)
//		e := apperrors.NewBadRequest("invalid user id")
//		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
//		return
//	}
//	if t, uErr := h.CertService.GetByAccreditingOrg(c, uid); uErr != nil {
//		log.Printf("Unable to find organization: %v\n%v", uid, uErr)
//		e := apperrors.NewNotFound("organization", uid.String())
//		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"organization": t,
//		})
//	}
//}

func TestGetOrgs(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockOrgResp := []*models.Organization{{
			ID:   uid,
			Name: "Test Org",
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService: new(mocks2.MockOrgService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:          tf.router,
			OrgService: tf.mockOrgService,
		})
		// Use the mock directly from the fixture
		tf.mockOrgService.On("GetAll", mock.Anything).Return(mockOrgResp, nil)

		// Create request
		request, err := http.NewRequest(http.MethodGet, "/organizations/", nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected response body
		respBody, err := json.Marshal(gin.H{
			"organizations": mockOrgResp,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockOrgService.AssertExpectations(t)
	})

	t.Run("NotFound", func(t *testing.T) {
		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService: new(mocks2.MockOrgService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:          tf.router,
			OrgService: tf.mockOrgService,
		})
		// Setup mock to return an error
		tf.mockOrgService.On("GetAll", mock.Anything).Return([]*models.Organization{}, fmt.Errorf("not found"))

		// Create request
		request, err := http.NewRequest(http.MethodGet, "/organizations/", nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewNotFound("organizations", "not found")
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusNotFound, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockOrgService.AssertExpectations(t)
	})
}
func TestGetOrgsInSearch(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockOrgResp := []*models.Organization{{
			ID:   uid,
			Name: "Test Org",
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService: new(mocks2.MockOrgService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:          tf.router,
			OrgService: tf.mockOrgService,
		})
		// Use the mock directly from the fixture
		tf.mockOrgService.On("GetAllInSearch", mock.Anything, uid).Return(mockOrgResp, nil)

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/organizations/search/%s", uid), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected response body
		respBody, err := json.Marshal(gin.H{
			"organization": mockOrgResp,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockOrgService.AssertExpectations(t)
	})

	t.Run("BadRequest", func(t *testing.T) {
		// Test with invalid UUID
		invalidID := "invalid-uuid"

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService: new(mocks2.MockOrgService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:               tf.router,
			ResourceService: tf.mockResourceService,
		})
		// Create request with invalid UUID
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/organizations/search/%s", invalidID), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewBadRequest("invalid organization id")
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockOrgService.AssertNotCalled(t, "GetAllInSearch", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid := uuid.New()
		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService: new(mocks2.MockOrgService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:          tf.router,
			OrgService: tf.mockOrgService,
		})
		// Setup mock to return an error
		tf.mockOrgService.On("GetAllInSearch", mock.Anything, uid).Return([]*models.Organization{}, fmt.Errorf("not found"))

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/organizations/search/%s", uid), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewNotFound("organization", uid.String())
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusNotFound, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockOrgService.AssertExpectations(t)
	})
}
func TestGetOrg(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockOrgResp := &models.Organization{
			ID:   uid,
			Name: "Test Org",
		}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService: new(mocks2.MockOrgService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:          tf.router,
			OrgService: tf.mockOrgService,
		})
		// Use the mock directly from the fixture
		tf.mockOrgService.On("Get", mock.Anything, uid).Return(mockOrgResp, nil)

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/organizations/%s", uid), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected response body
		respBody, err := json.Marshal(gin.H{
			"organization": mockOrgResp,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockOrgService.AssertExpectations(t)
	})

	t.Run("BadRequest", func(t *testing.T) {
		// Test with invalid UUID
		invalidID := "invalid-uuid"

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService: new(mocks2.MockOrgService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:               tf.router,
			ResourceService: tf.mockResourceService,
		})
		// Create request with invalid UUID
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/organizations/%s", invalidID), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewBadRequest("invalid organization id")
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockOrgService.AssertNotCalled(t, "Get", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid := uuid.New()
		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService: new(mocks2.MockOrgService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:          tf.router,
			OrgService: tf.mockOrgService,
		})
		// Setup mock to return an error
		tf.mockOrgService.On("Get", mock.Anything, uid).Return(&models.Organization{}, fmt.Errorf("not found"))

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/organizations/%s", uid), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewNotFound("organization", uid.String())
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusNotFound, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockOrgService.AssertExpectations(t)
	})
}
func TestGetOrgCertifications(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockOrgResp := []*models.Certification{{
			ID:   uid,
			Name: "Test Org",
		}}

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService:  new(mocks2.MockOrgService),
			mockCertService: new(mocks2.MockCertService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			OrgService:  tf.mockOrgService,
			CertService: tf.mockCertService,
		})
		// Use the mock directly from the fixture
		tf.mockCertService.On("GetByAccreditingOrg", mock.Anything, uid).Return(mockOrgResp, nil)

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/organizations/%s/certifications", uid), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected response body
		respBody, err := json.Marshal(gin.H{
			"organization": mockOrgResp,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusOK, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockOrgService.AssertExpectations(t)
	})

	t.Run("BadRequest", func(t *testing.T) {
		// Test with invalid UUID
		invalidID := "invalid-uuid"

		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService:  new(mocks2.MockOrgService),
			mockCertService: new(mocks2.MockCertService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:               tf.router,
			ResourceService: tf.mockResourceService,
			CertService:     tf.mockCertService,
		})
		// Create request with invalid UUID
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/organizations/%s/certifications", invalidID), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewBadRequest("invalid organization id")
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusBadRequest, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockCertService.AssertNotCalled(t, "GetByAccreditingOrg", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid := uuid.New()
		// Initialize the test fixture with the mock service
		tf := &testFixture{
			mockOrgService:  new(mocks2.MockOrgService),
			mockCertService: new(mocks2.MockCertService),
		}
		tf.setup()
		handlers.NewHandler(&handlers.Config{
			R:           tf.router,
			OrgService:  tf.mockOrgService,
			CertService: tf.mockCertService,
		})
		// Setup mock to return an error
		tf.mockCertService.On("GetByAccreditingOrg", mock.Anything, uid).Return([]*models.Certification{}, fmt.Errorf("not found"))

		// Create request
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/organizations/%s/certifications", uid), nil)
		assert.NoError(t, err)

		// Call the handler via router
		tf.router.ServeHTTP(tf.rr, request)

		// Prepare expected error response
		respErr := apperrors.NewNotFound("organization", uid.String())
		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		// Check assertions
		assert.Equal(t, http.StatusNotFound, tf.rr.Code)
		assert.Equal(t, respBody, tf.rr.Body.Bytes())
		tf.mockCertService.AssertExpectations(t)
	})

}

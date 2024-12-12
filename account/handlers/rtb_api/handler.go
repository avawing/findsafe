package handlers

import (
	"findsafe/account/models"
	"findsafe/account/models/apperrors"
	"findsafe/account/models/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// Handler struct holds required services for handler to function
type Handler struct {
	UserService interfaces.UserService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R           *gin.Engine
	UserService interfaces.UserService
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create an account group
	// Create a handler (which will later have injected services)
	h := &Handler{
		UserService: c.UserService,
	} // currently has no properties

	// Create a group, or base url for all routes
	g := c.R.Group("/users")
	g.GET("/me", h.Me)
	g.GET("/:id", h.GetUser)
	g.PUT("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)
	g.GET("/organization/:id", h.GetUsersInOrg)
	// Search Admin
	g.GET("/search/:id", h.GetUsersInSearch)
	g.GET("/sortie/:id", h.GetUsersInSortie)
	g.POST("/resources", h.CreateUserResource)
}

// Me handler calls services for getting
// a user's details
func (h *Handler) Me(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := apperrors.NewInternal()
		c.AbortWithStatusJSON(err.Status(), gin.H{"error": err})
	}

	uid := user.(*models.User).ID
	if u, err := h.UserService.Get(c, uid); err != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user": u,
		})
	}
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if u, uErr := h.UserService.Get(c, uid); uErr != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user": u,
		})
	}
}

func (h *Handler) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) GetUsersInOrg(c *gin.Context) {
	_ = c.DefaultQuery("Type", "")
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if u, uErr := h.UserService.GetAllInOrg(c, uid); uErr != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"users": u,
		})
	}
}

func (h *Handler) GetUsersInSearch(c *gin.Context) {
	_ = c.DefaultQuery("Type", "")
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if u, uErr := h.UserService.GetAllinSearch(c, uid); uErr != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"users": u,
		})
	}
}

func (h *Handler) GetUsersInSortie(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if u, uErr := h.UserService.GetAllInSortie(c, uid); uErr != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"users": u,
		})
	}
}

func (h *Handler) CreateUserResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

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
	UserService     interfaces.UserService
	TeamService     interfaces.TeamService
	CertService     interfaces.CertService
	ResourceService interfaces.ResourceService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R               *gin.Engine
	UserService     interfaces.UserService
	TeamService     interfaces.TeamService
	CertService     interfaces.CertService
	ResourceService interfaces.ResourceService
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create an account group
	// Create a handler (which will later have injected services)
	h := &Handler{
		UserService:     c.UserService,
		TeamService:     c.TeamService,
		CertService:     c.CertService,
		ResourceService: c.ResourceService,
	}

	// Create a group, or base url for all routes
	g := c.R.Group("/users")
	g.GET("/me", h.Me)
	g.GET("/:id", h.GetUser)
	g.PUT("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)
	g.GET("/organization/:id", h.GetUsersInOrg)
	g.GET("/:id/certifications", h.GetUserCertifications)

	n := g.Group("/certifications")
	n.GET("/:id", h.GetUserCertification)
	n.PUT("/:id", h.UpdateUserCertification)
	n.DELETE("/:id", h.DeleteUserCertification)
	// Search Admin
	g.GET("/search/:id", h.GetUsersInSearch)
	g.GET("/sortie/:id", h.GetUsersInSortie)
	g.POST("/resources", h.CreateUserResource)
	g.GET("/:id/resources", h.GetUserResources)

	t := c.R.Group("/teams")
	t.GET("/:id", h.GetTeam)
	t.PUT("/:id", h.UpdateTeam)
	t.DELETE("/:id", h.DeleteTeam)

	t.GET("/search/:id", h.GetTeams)
	t.GET("/search/:id/unassigned", h.GetUnassigned)
	t.GET("/sortie/:id", h.GetSortie)

	r := c.R.Group("/resources")
	r.GET("/:id", h.GetResource)
	r.PUT("/:id", h.UpdateUserResource)
	r.DELETE("/:id", h.DeleteUserResource)

}

func (h *Handler) Me(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := apperrors.NewInternal()
		c.AbortWithStatusJSON(err.Status(), gin.H{"error": err})
	}

	us, ok := user.(*models.User)
	if !ok {
		log.Printf("User in context is not of type *models.User: %v\n", user)
		err := apperrors.NewInternal()
		c.AbortWithStatusJSON(err.Status(), gin.H{"error": err})
		return // Ensure we don't continue further in this handler
	}

	if u, err := h.UserService.Get(c, us.ID); err != nil {
		log.Printf("Unable to find user: %v\n%v", us.ID, err)
		e := apperrors.NewNotFound("user", us.ID.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user": u,
		})
	}
}

// Users
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

// Resources
func (h *Handler) CreateUserResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}
func (h *Handler) UpdateUserResource(c *gin.Context) {}
func (h *Handler) DeleteUserResource(c *gin.Context) {}
func (h *Handler) GetResource(c *gin.Context) {
	_ = c.DefaultQuery("Type", "")
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if u, uErr := h.ResourceService.Get(c, uid); uErr != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"users": u,
		})
	}
}
func (h *Handler) GetUserResources(c *gin.Context) {
	_ = c.DefaultQuery("Type", "")
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if u, uErr := h.ResourceService.GetByOwnerID(c, uid); uErr != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"users": u,
		})
	}
}

// Certifications
func (h *Handler) GetUserCertifications(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if u, uErr := h.CertService.GetByUserID(c, uid); uErr != nil {
		log.Printf("Unable to find certifications: %v\n%v", uid, uErr)
		e := apperrors.NewNotFound("user", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"certifications": u,
		})
	}
}
func (h *Handler) CreateUserCertification(c *gin.Context) {}
func (h *Handler) UpdateUserCertification(c *gin.Context) {}
func (h *Handler) DeleteUserCertification(c *gin.Context) {}
func (h *Handler) GetUserCertification(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if u, uErr := h.CertService.Get(c, uid); uErr != nil {
		log.Printf("Unable to find certifications: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"certifications": u,
		})
	}
}

// Teams
func (h *Handler) GetTeam(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if t, uErr := h.TeamService.Get(c, uid); uErr != nil {
		log.Printf("Unable to find team: %v\n%v", uid, uErr)
		e := apperrors.NewNotFound("team", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"team": t,
		})
	}
}
func (h *Handler) UpdateTeam(c *gin.Context) {}
func (h *Handler) DeleteTeam(c *gin.Context) {}
func (h *Handler) GetTeams(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if t, tErr := h.TeamService.GetAllinSearch(c, uid); tErr != nil {
		log.Printf("Unable to find team: %v\n%v", uid, err)
		e := apperrors.NewNotFound("search id", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"teams": t,
		})
	}
}
func (h *Handler) GetSortie(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Unable to generate UUID from request: %v\n", err)
		e := apperrors.NewBadRequest("invalid user id")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
		return
	}
	if t, tErr := h.TeamService.GetBySortie(c, uid); tErr != nil {
		log.Printf("Unable to find team: %v\n%v", uid, err)
		e := apperrors.NewNotFound("sortie id", uid.String())
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"team": t,
		})
	}
}
func (h *Handler) GetUnassigned(c *gin.Context) {
	if t, tErr := h.TeamService.GetAllWithoutSortie(c); tErr != nil {
		log.Printf("No Unassigned Teams: %v\n%v", tErr)
		e := apperrors.NewNotFound("teams", "unassigned")
		c.AbortWithStatusJSON(e.Status(), gin.H{"error": e})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"teams": t,
		})
	}
}

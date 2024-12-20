package main

import (
	"context"
	handlers "findsafe/backend/handlers/rtb_api"
	"findsafe/backend/repository"
	"findsafe/backend/services"
	"findsafe/backend/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//	@title			FindSafe Main API Server
//	@version		1.0
//	@description	FindSafe API -

//	@contact.name	Ava Wingfield
//	@contact.email	avelinewingfield@gmail.com

//	@host		localhost:8080
//	@BasePath	/
//  @securityDefinitions.apikey ApiKeyAuth
//  @in header
//  @name Authorization
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func inject() *gin.Engine {
	router := gin.Default()
	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf(err.Error())
	}

	handlers.NewHandler(&handlers.Config{
		R: router,
		CertService: services.NewCertService(&services.USConfig{
			CertRepository: repository.NewCertRepository(db),
		}),
		OrgService: services.NewOrgService(&services.USConfig{
			OrgRepository: repository.NewOrgRepository(db),
		}),
		ResourceService: services.NewResourceService(&services.USConfig{
			ResourceRepository: repository.NewResourceRepository(db),
		}),
		SearchService: services.NewSearchService(&services.USConfig{
			SearchRepository: repository.NewSearchRepository(db),
		}),
		TeamService: services.NewTeamService(&services.USConfig{
			TeamRepository: repository.NewTeamRepository(db),
		}),
		UserService: services.NewUserService(&services.USConfig{
			UserRepository: repository.NewUserRepository(db),
		}),
	})
	return router
}

func main() {
	r := inject()

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}

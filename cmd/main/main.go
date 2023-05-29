package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/surysatriah/go-dashboard-app/internal/database"
	"github.com/surysatriah/go-dashboard-app/internal/handlers"
	"github.com/surysatriah/go-dashboard-app/internal/handlers/middleware"
	"github.com/surysatriah/go-dashboard-app/internal/handlers/router"
)

func main() {

	database.ConnectDatabase()
	go handlers.ConnectMqtt()
	serveApplication()

}

func serveApplication() {
	r := gin.Default()

	publicRoutes := r.Group("/")
	router.PublicRoutes(publicRoutes)

	protectedRoutes := r.Group("/")
	protectedRoutes.Use(middleware.Authentication())
	router.ProtectedRoutes(protectedRoutes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	r.Run(":" + port)
}

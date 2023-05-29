package router

import (
	"github.com/gin-gonic/gin"
	"github.com/surysatriah/go-dashboard-app/internal/handlers/controller"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.POST("/register", controller.UserRegister)
	g.POST("/login", controller.UserLogin)
}

func ProtectedRoutes(g *gin.RouterGroup) {
	g.GET("/protected", controller.Protected)
}

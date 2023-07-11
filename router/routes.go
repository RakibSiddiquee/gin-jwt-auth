package router

import (
	"github.com/RakibSiddiquee/gin-jwt-auth/controllers"
	"github.com/RakibSiddiquee/gin-jwt-auth/middleware"
	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
}

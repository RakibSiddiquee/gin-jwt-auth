package main

import (
	"github.com/RakibSiddiquee/gin-jwt-auth/controllers"
	"github.com/RakibSiddiquee/gin-jwt-auth/initializers"
	"github.com/RakibSiddiquee/gin-jwt-auth/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
package main

import (
	"github.com/RakibSiddiquee/gin-jwt-auth/initializers"
	"github.com/RakibSiddiquee/gin-jwt-auth/router"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	router.GetRoutes(r)

	r.Run()
}

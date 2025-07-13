package main

import (
	"GO_AUTH_JWT/controllers"
	"GO_AUTH_JWT/initializers"
	"GO_AUTH_JWT/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.EnvLoadVariables()
	initializers.ConnectTodatabase()
}

func main() {
	r := gin.Default()
	r.POST("/Signup", controllers.Signup)
	r.POST("/Login", controllers.Login)
	r.GET("/Validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}

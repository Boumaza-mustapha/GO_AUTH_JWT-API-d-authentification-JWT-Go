package main

import (
	"GO_AUTH_JWT/initializers"
	"GO_AUTH_JWT/models"
)

func init() {
	initializers.EnvLoadVariables()
	initializers.ConnectTodatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}

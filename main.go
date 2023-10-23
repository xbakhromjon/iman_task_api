package main

import (
	"github.com/gin-gonic/gin"
	"iman_task_api/controllers"
	"iman_task_api/initializers"
	"iman_task_api/middleware"
)

func main() {
	initialize()
	router := gin.Default()
	router.GET("/day-difference", middleware.RequireAuth, controllers.CalculateDateDifference)
	router.Run("localhost:8080")
}

func initialize() {
	initializers.LoadEnvVariables()
}

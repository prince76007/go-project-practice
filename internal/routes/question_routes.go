package routes

import (
	"github.com/gin-gonic/gin"
	"your_project/internal/handlers"
	"your_project/internal/middleware"
)

func SetupQuestionRoutes(router *gin.Engine) {
	questionGroup := router.Group("/questions")
	questionGroup.Use(middleware.AuthMiddleware())

	questionGroup.GET("/", handlers.GetQuestions)
	questionGroup.POST("/", handlers.CreateQuestion)
	questionGroup.PUT("/:id", handlers.UpdateQuestion)
	questionGroup.DELETE("/:id", handlers.DeleteQuestion)
}
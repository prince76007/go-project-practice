package routes

import (
	"go-project-practice/internal/database"
	"go-project-practice/internal/handlers"
	"go-project-practice/internal/repositories"
	"go-project-practice/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterQuestionRoutes(router *gin.Engine) {
	db := database.GetDB()
	questionRepo := repositories.NewQuestionRepository(db)
	questionService := services.NewQuestionService(*questionRepo)
	questionHandler := handlers.NewQuestionHandler(*questionService)

	questionRoutes := router.Group("/api/questions")
	{
		questionRoutes.GET("", questionHandler.GetQuestions)
		questionRoutes.POST("", questionHandler.CreateQuestion)
		questionRoutes.GET("/:id", questionHandler.GetQuestion)
		questionRoutes.PUT("/:id", questionHandler.UpdateQuestion)
		questionRoutes.DELETE("/:id", questionHandler.DeleteQuestion)
	}
}

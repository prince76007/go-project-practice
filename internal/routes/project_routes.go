package routes

import (
	"go-project-practice/internal/database"
	"go-project-practice/internal/handlers"
	"go-project-practice/internal/repositories"
	"go-project-practice/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(router *gin.Engine) {
	db := database.GetDB()
	projectRepo := repositories.NewProjectRepository(db)
	projectService := services.NewProjectService(*projectRepo)
	projectHandler := handlers.NewProjectHandler(*projectService)

	projectRoutes := router.Group("/api/projects")
	{
		projectRoutes.POST("", projectHandler.CreateProject)
		projectRoutes.GET("/:id", projectHandler.GetProject)
		projectRoutes.PATCH("/:id", projectHandler.UpdateProject)
		projectRoutes.DELETE("/:id", projectHandler.DeleteProject)
	}
}

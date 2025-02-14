package routes

import (
	"go-project-practice/internal/database"
	"go-project-practice/internal/handlers"
	"go-project-practice/internal/repositories"
	"go-project-practice/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProjectRoutes(router *mux.Router) {
	db := database.GetDB()
	projectRepo := repositories.NewProjectRepository(db)
	projectService := services.NewProjectService(*projectRepo)
	projectHandler := handlers.NewProjectHandler(*projectService)

	projectRouter := router.PathPrefix("/api/projects").Subrouter()

	projectRouter.HandleFunc("", projectHandler.CreateProject).Methods(http.MethodPost)
	projectRouter.HandleFunc("/{id}", projectHandler.GetProject).Methods(http.MethodGet)
	projectRouter.HandleFunc("/{id}", projectHandler.UpdateProject).Methods(http.MethodPut)
	projectRouter.HandleFunc("/{id}", projectHandler.DeleteProject).Methods(http.MethodDelete)
}

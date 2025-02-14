package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"your_project/internal/handlers"
	"your_project/internal/middleware"
)

func RegisterProjectRoutes(r *mux.Router) {
	projectRouter := r.PathPrefix("/api/projects").Subrouter()
	projectRouter.Use(middleware.AuthMiddleware)

	projectRouter.HandleFunc("", handlers.CreateProject).Methods(http.MethodPost)
	projectRouter.HandleFunc("", handlers.GetProjects).Methods(http.MethodGet)
	projectRouter.HandleFunc("/{id}", handlers.GetProject).Methods(http.MethodGet)
	projectRouter.HandleFunc("/{id}", handlers.UpdateProject).Methods(http.MethodPut)
	projectRouter.HandleFunc("/{id}", handlers.DeleteProject).Methods(http.MethodDelete)
}
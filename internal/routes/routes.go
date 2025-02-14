package routes

import (
	"github.com/gorilla/mux"
)

// SetupRouter initializes the router and registers all routes.
func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	RegisterProjectRoutes(router)
	return router
}

package routes

import (
	"go-project-practice/internal/database"
	"go-project-practice/internal/handlers"
	"go-project-practice/internal/repositories"
	"go-project-practice/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func RegisterQuestionRoutes(router *mux.Router) {
	db := database.GetDB()
	questionRepo := repositories.NewQuestionRepository(db)
	questionService := services.NewQuestionService(*questionRepo)
	questionHandler := handlers.NewQuestionHandler(*questionService)

	questionRouter := router.PathPrefix("/api/questions").Subrouter()

	questionRouter.HandleFunc("", ginHandlerWrapper(questionHandler.GetQuestions)).Methods(http.MethodGet)
	questionRouter.HandleFunc("", ginHandlerWrapper(questionHandler.CreateQuestion)).Methods(http.MethodPost)
	questionRouter.HandleFunc("/{id}", ginHandlerWrapper(questionHandler.GetQuestion)).Methods(http.MethodGet)
	questionRouter.HandleFunc("/{id}", ginHandlerWrapper(questionHandler.UpdateQuestion)).Methods(http.MethodPut)
	questionRouter.HandleFunc("/{id}", ginHandlerWrapper(questionHandler.DeleteQuestion)).Methods(http.MethodDelete)
}

func ginHandlerWrapper(h func(c *gin.Context)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		h(c)
	}
}

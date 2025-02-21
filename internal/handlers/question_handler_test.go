package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-project-practice/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Service
type MockQuestionService struct {
	mock.Mock
}

func TestGetQuestions(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/projects/:project_id/questions", handler.GetQuestions)

	questions := []models.Question{
		{ID: 1, ProjectID: 1, Question: "What is Go?"},
		{ID: 2, ProjectID: 1, Question: "What is a goroutine?"},
	}

	mockService.On("GetQuestionsByProject", 1).Return(questions, nil)

	req, _ := http.NewRequest("GET", "/projects/1/questions", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response []models.Question
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, questions, response)
	mockService.AssertExpectations(t)
}

func TestCreateQuestion(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/projects/:project_id/questions", handler.CreateQuestion)

	question := &models.Question{
		ProjectID: 1,
		Question:  "What is Go?",
	}

	mockService.On("CreateQuestion", question).Return(nil)

	body, _ := json.Marshal(question)
	req, _ := http.NewRequest("POST", "/projects/1/questions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateQuestion(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/questions/:id", handler.UpdateQuestion)

	question := &models.Question{
		ID:        1,
		ProjectID: 1,
		Question:  "Updated Question",
	}

	mockService.On("UpdateQuestion", 1, question).Return(nil)

	body, _ := json.Marshal(question)
	req, _ := http.NewRequest("PUT", "/questions/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteQuestion(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/questions/:id", handler.DeleteQuestion)

	mockService.On("DeleteQuestion", 1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/questions/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockService.AssertExpectations(t)
}

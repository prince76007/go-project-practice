package handlers

import (
	"bytes"
	"encoding/json"
	"go-project-practice/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQuestionService struct {
	mock.Mock
}

func (m *MockQuestionService) CreateQuestion(question *models.Question) error {
	args := m.Called(question)
	return args.Error(0)
}

func (m *MockQuestionService) GetQuestion(id int) (*models.Question, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionService) UpdateQuestion(id int, question *models.Question) error {
	args := m.Called(id, question)
	return args.Error(0)
}

func (m *MockQuestionService) DeleteQuestion(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockQuestionService) GetQuestionsByProject(projectID int) ([]models.Question, error) {
	args := m.Called(projectID)
	return args.Get(0).([]models.Question), args.Error(1)
}

func TestGetQuestions(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	questions := []models.Question{
		{ID: 1, ProjectID: 1, Question: "What is Go?"},
		{ID: 2, ProjectID: 1, Question: "What is Gin?"},
	}

	mockService.On("GetQuestionsByProject", 1).Return(questions, nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/projects/:project_id/questions", handler.GetQuestions)

	req, _ := http.NewRequest(http.MethodGet, "/projects/1/questions", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response []models.Question
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, questions, response)
	mockService.AssertExpectations(t)
}

func TestCreateQuestion(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	question := &models.Question{
		ProjectID: 1,
		Question:  "What is Go?",
	}

	mockService.On("CreateQuestion", question).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/projects/:project_id/questions", handler.CreateQuestion)

	body, _ := json.Marshal(question)
	req, _ := http.NewRequest(http.MethodPost, "/projects/1/questions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateQuestion(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	question := &models.Question{
		ID:        1,
		ProjectID: 1,
		Question:  "What is Go?",
	}

	mockService.On("UpdateQuestion", 1, question).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/questions/:id", handler.UpdateQuestion)

	body, _ := json.Marshal(question)
	req, _ := http.NewRequest(http.MethodPut, "/questions/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteQuestion(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	mockService.On("DeleteQuestion", 1).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/questions/:id", handler.DeleteQuestion)

	req, _ := http.NewRequest(http.MethodDelete, "/questions/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

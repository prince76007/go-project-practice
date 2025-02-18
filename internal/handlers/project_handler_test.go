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

type MockProjectService struct {
	mock.Mock
}

func (m *MockProjectService) CreateProject(project *models.Project) error {
	args := m.Called(project)
	return args.Error(0)
}

func (m *MockProjectService) GetProject(id int) (*models.Project, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Project), args.Error(1)
}

func (m *MockProjectService) UpdateProject(project *models.Project) error {
	args := m.Called(project)
	return args.Error(0)
}

func (m *MockProjectService) DeleteProject(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProjectService) GetAllProjects(tenantID int) ([]models.Project, error) {
	args := m.Called(tenantID)
	return args.Get(0).([]models.Project), args.Error(1)
}

func TestCreateProject(t *testing.T) {
	mockService := new(MockProjectService)
	handler := NewProjectHandler(mockService)

	project := &models.Project{
		Name: "Test Project",
	}

	mockService.On("CreateProject", project).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/projects", handler.CreateProject)

	body, _ := json.Marshal(project)
	req, _ := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetProject(t *testing.T) {
	mockService := new(MockProjectService)
	handler := NewProjectHandler(mockService)

	project := &models.Project{
		ID:   1,
		Name: "Test Project",
	}

	mockService.On("GetProject", 1).Return(project, nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/projects/:id", handler.GetProject)

	req, _ := http.NewRequest(http.MethodGet, "/projects/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Project
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, project, &response)
	mockService.AssertExpectations(t)
}

func TestUpdateProject(t *testing.T) {
	mockService := new(MockProjectService)
	handler := NewProjectHandler(mockService)

	project := &models.Project{
		ID:   1,
		Name: "Test Project",
	}

	mockService.On("UpdateProject", project).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/projects/:id", handler.UpdateProject)

	body, _ := json.Marshal(project)
	req, _ := http.NewRequest(http.MethodPut, "/projects/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteProject(t *testing.T) {
	mockService := new(MockProjectService)
	handler := NewProjectHandler(mockService)

	mockService.On("DeleteProject", 1).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/projects/:id", handler.DeleteProject)

	req, _ := http.NewRequest(http.MethodDelete, "/projects/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

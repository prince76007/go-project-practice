/*package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-project-practice/internal/models"
	"go-project-practice/internal/services"
)

// Mock Service
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

	gin.SetMode*/
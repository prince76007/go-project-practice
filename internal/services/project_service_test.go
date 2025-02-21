package services

import (
	"testing"

	"go-project-practice/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Repository
type MockProjectRepo struct {
	mock.Mock
}

func (m *MockProjectRepo) Create(project *models.Project) error {
	args := m.Called(project)
	return args.Error(0)
}

func (m *MockProjectRepo) GetByID(id int) (*models.Project, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Project), args.Error(1)
}

func (m *MockProjectRepo) Update(project *models.Project) error {
	args := m.Called(project)
	return args.Error(0)
}

func (m *MockProjectRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProjectRepo) GetAllByTenantID(tenantID int) ([]models.Project, error) {
	args := m.Called(tenantID)
	return args.Get(0).([]models.Project), args.Error(1)
}

func TestCreateProject(t *testing.T) {
	mockRepo := new(MockProjectRepo)
	service := NewProjectService(mockRepo)

	project := &models.Project{
		Name: "Test Project",
	}

	mockRepo.On("Create", project).Return(nil)

	err := service.CreateProject(project)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetProject(t *testing.T) {
	mockRepo := new(MockProjectRepo)
	service := NewProjectService(mockRepo)

	project := &models.Project{
		ID:   1,
		Name: "Test Project",
	}

	mockRepo.On("GetByID", 1).Return(project, nil)

	result, err := service.GetProject(1)
	assert.NoError(t, err)
	assert.Equal(t, project, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProject(t *testing.T) {
	mockRepo := new(MockProjectRepo)
	service := NewProjectService(mockRepo)

	project := &models.Project{
		ID:   1,
		Name: "Updated Project",
	}

	mockRepo.On("Update", project).Return(nil)

	err := service.UpdateProject(project)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProject(t *testing.T) {
	mockRepo := new(MockProjectRepo)
	service := NewProjectService(mockRepo)

	mockRepo.On("Delete", 1).Return(nil)

	err := service.DeleteProject(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllProjects(t *testing.T) {
	mockRepo := new(MockProjectRepo)
	service := NewProjectService(mockRepo)

	projects := []models.Project{
		{ID: 1, Name: "Project 1"},
		{ID: 2, Name: "Project 2"},
	}

	mockRepo.On("GetAllByTenantID", 1).Return(projects, nil)

	result, err := service.GetAllProjects(1)
	assert.NoError(t, err)
	assert.Equal(t, projects, result)
	mockRepo.AssertExpectations(t)
}

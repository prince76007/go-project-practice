package services

import (
	"testing"

	"go-project-practice/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQuestionRepository struct {
	mock.Mock
}

func (m *MockQuestionRepository) Create(question *models.Question) error {
	args := m.Called(question)
	return args.Error(0)
}

func (m *MockQuestionRepository) GetByID(id int) (*models.Question, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionRepository) Update(question *models.Question) error {
	args := m.Called(question)
	return args.Error(0)
}

func (m *MockQuestionRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockQuestionRepository) GetByProjectID(projectID int) ([]models.Question, error) {
	args := m.Called(projectID)
	return args.Get(0).([]models.Question), args.Error(1)
}

func TestCreateQuestion(t *testing.T) {
	mockRepo := new(MockQuestionRepository)
	service := NewQuestionService(mockRepo)

	question := &models.Question{
		ProjectID: 1,
		Question:  "What is Go?",
	}

	mockRepo.On("Create", question).Return(nil)

	err := service.CreateQuestion(question)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetQuestion(t *testing.T) {
	mockRepo := new(MockQuestionRepository)
	service := NewQuestionService(mockRepo)

	question := &models.Question{
		ID:        1,
		ProjectID: 1,
		Question:  "What is Go?",
	}

	mockRepo.On("GetByID", 1).Return(question, nil)

	result, err := service.GetQuestion(1)
	assert.NoError(t, err)
	assert.Equal(t, question, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateQuestion(t *testing.T) {
	mockRepo := new(MockQuestionRepository)
	service := NewQuestionService(mockRepo)

	question := &models.Question{
		ID:        1,
		ProjectID: 1,
		Question:  "What is Go?",
	}

	mockRepo.On("Update", question).Return(nil)

	err := service.UpdateQuestion(1, question)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteQuestion(t *testing.T) {
	mockRepo := new(MockQuestionRepository)
	service := NewQuestionService(mockRepo)

	mockRepo.On("Delete", 1).Return(nil)

	err := service.DeleteQuestion(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetQuestionsByProject(t *testing.T) {
	mockRepo := new(MockQuestionRepository)
	service := NewQuestionService(mockRepo)

	questions := []models.Question{
		{ID: 1, ProjectID: 1, Question: "What is Go?"},
		{ID: 2, ProjectID: 1, Question: "What is Gin?"},
	}

	mockRepo.On("GetByProjectID", 1).Return(questions, nil)

	result, err := service.GetQuestionsByProject(1)
	assert.NoError(t, err)
	assert.Equal(t, questions, result)
	mockRepo.AssertExpectations(t)
}

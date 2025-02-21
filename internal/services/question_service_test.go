package services

import (
	"testing"

	"go-project-practice/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Repository
type MockQuestionRepo struct {
	mock.Mock
}

func (m *MockQuestionRepo) GetAll() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockQuestionRepo) Create(question *models.Question) error {
	args := m.Called(question)
	return args.Error(0)
}

func (m *MockQuestionRepo) GetByID(id int) (*models.Question, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionRepo) Update(question *models.Question) error {
	args := m.Called(question)
	return args.Error(0)
}

func (m *MockQuestionRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockQuestionRepo) GetByProjectID(projectID int) ([]models.Question, error) {
	args := m.Called(projectID)
	return args.Get(0).([]models.Question), args.Error(1)
}

func TestGetAllQuestions(t *testing.T) {
	mockRepo := new(MockQuestionRepo)
	mockService := QuestionService{Repo: mockRepo}

	mockRepo.On("GetAll").Return([]string{"Q1", "Q2"}, nil)

	questions, err := mockService.GetAllQuestions()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(questions))
}

func TestCreateQuestion(t *testing.T) {
	mockRepo := new(MockQuestionRepo)
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
	mockRepo := new(MockQuestionRepo)
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
	mockRepo := new(MockQuestionRepo)
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
	mockRepo := new(MockQuestionRepo)
	service := NewQuestionService(mockRepo)

	mockRepo.On("Delete", 1).Return(nil)

	err := service.DeleteQuestion(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetQuestionsByProject(t *testing.T) {
	mockRepo := new(MockQuestionRepo)
	service := NewQuestionService(mockRepo)

	questions := []models.Question{
		{ID: 1, ProjectID: 1, Question: "What is Go?"},
		{ID: 2, ProjectID: 1, Question: "What is a goroutine?"},
	}

	mockRepo.On("GetByProjectID", 1).Return(questions, nil)

	result, err := service.GetQuestionsByProject(1)
	assert.NoError(t, err)
	assert.Equal(t, questions, result)
	mockRepo.AssertExpectations(t)
}

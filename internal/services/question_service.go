package services

import (
	"errors"
	"time"

	"go-project/internal/models"
	"go-project/internal/repositories"
)

type QuestionService struct {
	repo repositories.QuestionRepository
}

func NewQuestionService(repo repositories.QuestionRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) CreateQuestion(question *models.Question) error {
	if question.ProjectID == 0 || question.Question == "" {
		return errors.New("invalid question data")
	}
	question.CreateDate = time.Now()
	question.UpdateDate = time.Now()
	return s.repo.Create(question)
}

func (s *QuestionService) GetQuestion(id int) (*models.Question, error) {
	return s.repo.GetByID(id)
}

func (s *QuestionService) UpdateQuestion(question *models.Question) error {
	if question.ID == 0 || question.Question == "" {
		return errors.New("invalid question data")
	}
	question.UpdateDate = time.Now()
	return s.repo.Update(question)
}

func (s *QuestionService) DeleteQuestion(id int) error {
	return s.repo.Delete(id)
}

func (s *QuestionService) GetQuestionsByProject(projectID int) ([]models.Question, error) {
	return s.repo.GetByProjectID(projectID)
}
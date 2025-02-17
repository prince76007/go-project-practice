package repositories

import (
	"go-project-practice/internal/models"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) Create(question *models.Question) error {
	return r.db.Create(question).Error
}

func (r *QuestionRepository) GetByID(id int) (*models.Question, error) {
	var question models.Question
	err := r.db.First(&question, id).Error
	return &question, err
}

func (r *QuestionRepository) Update(question *models.Question) error {
	return r.db.Save(question).Error
}

func (r *QuestionRepository) Delete(id int) error {
	return r.db.Delete(&models.Question{}, id).Error
}

func (r *QuestionRepository) GetAll() ([]models.Question, error) {
	var questions []models.Question
	err := r.db.Find(&questions).Error
	return questions, err
}

func (r *QuestionRepository) GetByProjectID(projectID int) ([]models.Question, error) {
	var questions []models.Question
	err := r.db.Where("project_id = ?", projectID).Find(&questions).Error
	return questions, err
}

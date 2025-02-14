package repositories

import (
	"database/sql"
	"time"

	"go-project-practice/internal/models"
)

type QuestionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) Create(question *models.Question) error {
	query := `INSERT INTO questions (project_id, question, type, create_date, update_date, language_code, tenant_id) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	return r.db.QueryRow(query, question.ProjectID, question.Question, question.Type, time.Now(), time.Now(), question.LanguageCode, question.TenantID).Scan(&question.ID)
}

func (r *QuestionRepository) GetByID(id int) (*models.Question, error) {
	query := `SELECT id, project_id, question, type, create_date, update_date, language_code, tenant_id 
			  FROM questions WHERE id = $1`
	question := &models.Question{}
	err := r.db.QueryRow(query, id).Scan(&question.ID, &question.ProjectID, &question.Question, &question.Type, &question.CreateDate, &question.UpdateDate, &question.LanguageCode, &question.TenantID)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (r *QuestionRepository) Update(question *models.Question) error {
	query := `UPDATE questions SET project_id = $1, question = $2, type = $3, update_date = $4, language_code = $5, tenant_id = $6 
			  WHERE id = $7`
	_, err := r.db.Exec(query, question.ProjectID, question.Question, question.Type, time.Now(), question.LanguageCode, question.TenantID, question.ID)
	return err
}

func (r *QuestionRepository) Delete(id int) error {
	query := `DELETE FROM questions WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *QuestionRepository) GetAll() ([]models.Question, error) {
	query := `SELECT id, project_id, question, type, create_date, update_date, language_code, tenant_id FROM questions`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var question models.Question
		if err := rows.Scan(&question.ID, &question.ProjectID, &question.Question, &question.Type, &question.CreateDate, &question.UpdateDate, &question.LanguageCode, &question.TenantID); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func (r *QuestionRepository) GetByProjectID(projectID int) ([]models.Question, error) {
	query := `SELECT id, project_id, question, type, create_date, update_date, language_code, tenant_id 
			  FROM questions WHERE project_id = $1`
	rows, err := r.db.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var question models.Question
		if err := rows.Scan(&question.ID, &question.ProjectID, &question.Question, &question.Type, &question.CreateDate, &question.UpdateDate, &question.LanguageCode, &question.TenantID); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}

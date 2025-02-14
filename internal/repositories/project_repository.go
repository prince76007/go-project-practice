package repositories

import (
	"database/sql"
	"errors"

	"go-project-practice/internal/models"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	query := `INSERT INTO projects (tenant_id, name, language_code) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, project.TenantID, project.Name, project.LanguageCode).Scan(&project.ID)
	return err
}

func (r *ProjectRepository) GetByID(id int) (*models.Project, error) {
	query := `SELECT id, tenant_id, name, language_code FROM projects WHERE id = $1`
	project := &models.Project{}
	err := r.db.QueryRow(query, id).Scan(&project.ID, &project.TenantID, &project.Name, &project.LanguageCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("project not found")
		}
		return nil, err
	}
	return project, nil
}

func (r *ProjectRepository) Update(project *models.Project) error {
	query := `UPDATE projects SET tenant_id = $1, name = $2, language_code = $3 WHERE id = $4`
	_, err := r.db.Exec(query, project.TenantID, project.Name, project.LanguageCode, project.ID)
	return err
}

func (r *ProjectRepository) Delete(id int) error {
	query := `DELETE FROM projects WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ProjectRepository) GetAll() ([]models.Project, error) {
	query := `SELECT id, tenant_id, name, language_code FROM projects`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.TenantID, &project.Name, &project.LanguageCode); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (r *ProjectRepository) GetAllByTenantID(tenant_id int) ([]models.Project, error) {
	query := `SELECT id, tenant_id, name, language_code FROM projects WHERE tenant_id = $1`
	rows, err := r.db.Query(query, tenant_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.TenantID, &project.Name, &project.LanguageCode); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

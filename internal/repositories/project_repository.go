package repositories

import (
	"go-project-practice/internal/models"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *ProjectRepository) GetByID(id int) (*models.Project, error) {
	var project models.Project
	err := r.db.First(&project, id).Error
	return &project, err
}

func (r *ProjectRepository) Update(project *models.Project) error {
	return r.db.Save(project).Error
}

func (r *ProjectRepository) Delete(id int) error {
	return r.db.Delete(&models.Project{}, id).Error
}

func (r *ProjectRepository) GetAll() ([]models.Project, error) {
	var projects []models.Project
	err := r.db.Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) GetAllByTenantID(tenantID int) ([]models.Project, error) {
	var projects []models.Project
	err := r.db.Where("tenant_id = ?", tenantID).Find(&projects).Error
	return projects, err
}

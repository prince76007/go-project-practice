package services

import (
	"errors"
	"time"

	"go-project-practice/internal/models"
	"go-project-practice/internal/repositories"
)

type ProjectService struct {
	repo repositories.ProjectRepository
}

func NewProjectService(repo repositories.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(project *models.Project) error {
	if project.Name == "" {
		return errors.New("project name is required")
	}
	project.CreateDate = time.Now()
	return s.repo.Create(project)
}

func (s *ProjectService) GetProject(id int) (*models.Project, error) {
	return s.repo.GetByID(id)
}

func (s *ProjectService) UpdateProject(project *models.Project) error {
	if project.ID == 0 {
		return errors.New("project ID is required")
	}
	project.UpdateDate = time.Now()
	return s.repo.Update(project)
}

func (s *ProjectService) DeleteProject(id int) error {
	return s.repo.Delete(id)
}

func (s *ProjectService) GetAllProjects(tenantID int) ([]models.Project, error) {
	return s.repo.GetAllByTenantID(tenantID)
}

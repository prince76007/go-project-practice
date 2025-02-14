package services

import (
	"errors"
	"time"

	"go-project/internal/models"
	"go-project/internal/repositories"
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

func (s *ProjectService) GetProject(id string) (*models.Project, error) {
	return s.repo.GetByID(id)
}

func (s *ProjectService) UpdateProject(project *models.Project) error {
	if project.ID == "" {
		return errors.New("project ID is required")
	}
	project.UpdateDate = time.Now()
	return s.repo.Update(project)
}

func (s *ProjectService) DeleteProject(id string) error {
	return s.repo.Delete(id)
}

func (s *ProjectService) GetAllProjects(tenantID string) ([]models.Project, error) {
	return s.repo.GetAllByTenantID(tenantID)
}
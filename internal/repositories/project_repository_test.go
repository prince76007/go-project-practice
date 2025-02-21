package repositories

import (
	"testing"

	"go-project-practice/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateProject(t *testing.T) {
	repo := NewProjectRepository(testDB)

	project := &models.Project{
		Name: "Test Project",
	}

	err := repo.Create(project)
	assert.NoError(t, err)
	assert.NotZero(t, project.ID)
}

func TestGetProjectByID(t *testing.T) {
	repo := NewProjectRepository(testDB)

	project := &models.Project{
		Name: "Test Project",
	}

	repo.Create(project)

	result, err := repo.GetByID(project.ID)
	assert.NoError(t, err)
	assert.Equal(t, project.Name, result.Name)
}

func TestUpdateProject(t *testing.T) {
	repo := NewProjectRepository(testDB)

	project := &models.Project{
		Name: "Test Project",
	}

	repo.Create(project)

	project.Name = "Updated Project"
	err := repo.Update(project)
	assert.NoError(t, err)

	result, err := repo.GetByID(project.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Project", result.Name)
}

func TestDeleteProject(t *testing.T) {
	repo := NewProjectRepository(testDB)

	project := &models.Project{
		Name: "Test Project",
	}

	repo.Create(project)

	err := repo.Delete(project.ID)
	assert.NoError(t, err)

	result, err := repo.GetByID(project.ID)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetAllProjectsByTenantID(t *testing.T) {
	repo := NewProjectRepository(testDB)

	project1 := &models.Project{
		Name:     "Project 1",
		TenantID: 1,
	}
	project2 := &models.Project{
		Name:     "Project 2",
		TenantID: 1,
	}

	repo.Create(project1)
	repo.Create(project2)

	projects, err := repo.GetAllByTenantID(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(projects))
}

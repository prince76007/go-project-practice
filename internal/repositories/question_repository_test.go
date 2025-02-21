package repositories

import (
	"testing"

	"go-project-practice/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestGetAllQuestions(t *testing.T) {
	repo := NewQuestionRepository(testDB)

	questions, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Greater(t, len(questions), 0)
}

func TestCreateQuestion(t *testing.T) {
	repo := NewQuestionRepository(testDB)

	question := &models.Question{
		ProjectID: 1,
		Question:  "Is this a test question?",
	}

	err := repo.Create(question)
	assert.NoError(t, err)
	assert.NotZero(t, question.ID)
}

func TestGetQuestionByID(t *testing.T) {
	repo := NewQuestionRepository(testDB)

	question := &models.Question{
		ProjectID: 1,
		Question:  "Is this a test question?",
	}

	repo.Create(question)

	result, err := repo.GetByID(question.ID)
	assert.NoError(t, err)
	assert.Equal(t, question.Question, result.Question)
}

func TestUpdateQuestion(t *testing.T) {
	repo := NewQuestionRepository(testDB)

	question := &models.Question{
		ProjectID: 1,
		Question:  "Is this a test question?",
	}

	repo.Create(question)

	question.Question = "Updated Question"
	err := repo.Update(question)
	assert.NoError(t, err)

	result, err := repo.GetByID(question.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Question", result.Question)
}

func TestDeleteQuestion(t *testing.T) {
	repo := NewQuestionRepository(testDB)

	question := &models.Question{
		ProjectID: 1,
		Question:  "Is this a test question?",
	}

	repo.Create(question)

	err := repo.Delete(question.ID)
	assert.NoError(t, err)

	result, err := repo.GetByID(question.ID)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetQuestionsByProjectID(t *testing.T) {
	repo := NewQuestionRepository(testDB)

	question1 := &models.Question{
		ProjectID: 1,
		Question:  "Is this a test question?",
	}
	question2 := &models.Question{
		ProjectID: 1,
		Question:  "What is a goroutine?",
	}

	repo.Create(question1)
	repo.Create(question2)

	questions, err := repo.GetByProjectID(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(questions))
}

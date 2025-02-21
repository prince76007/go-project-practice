package handlers

import (
	"net/http"
	"strconv"

	"go-project-practice/internal/models"
	"go-project-practice/internal/services"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	service services.QuestionService
}

func NewQuestionHandler(service services.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: service}
}

func (h *QuestionHandler) GetQuestions(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project_id"})
		return
	}
	questions, err := h.service.GetQuestionsByProject(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, questions)
}

func (h *QuestionHandler) GetQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}
	question, err := h.service.GetQuestion(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}
	c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) CreateQuestion(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("project_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project_id"})
		return
	}

	//question := models.Question{ProjectID: projectID}
	var question models.Question = models.Question{ProjectID: projectID}
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.CreateQuestion(&question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (h *QuestionHandler) UpdateQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateErr := h.service.UpdateQuestion(id, &question)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateErr.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}
	if err := h.service.DeleteQuestion(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

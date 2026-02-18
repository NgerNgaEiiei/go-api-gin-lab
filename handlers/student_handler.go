package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/student-api/models"
	"example.com/student-api/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

func sendError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.Service.GetStudents()
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Failed to retrieve students")
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.Service.GetStudentByID(id)
	if err != nil {
		sendError(c, http.StatusNotFound, "Student not found")
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		sendError(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if student.Id == "" || student.Name == "" {
		sendError(c, http.StatusBadRequest, "id and name must not be empty")
		return
	}

	if student.GPA < 0.0 || student.GPA > 4.0 {
		sendError(c, http.StatusBadRequest, "gpa must be between 0.00 and 4.00")
		return
	}

	if err := h.Service.CreateStudent(student); err != nil {
		sendError(c, http.StatusInternalServerError, "Failed to create student")
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {

	id := c.Param("id")
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		sendError(c, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	updatedStudent, err := h.Service.UpdateStudent(id, student)
	if err != nil {
		sendError(c, http.StatusNotFound, "Student not found")
		return
	}

	c.JSON(http.StatusOK, updatedStudent)
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteStudent(id)
	if err != nil {
		sendError(c, http.StatusNotFound, "Student not found")
		return
	}

	c.Status(http.StatusNoContent)
}

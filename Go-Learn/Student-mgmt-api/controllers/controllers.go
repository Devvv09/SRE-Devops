package controllers

import (
	"log"
	"net/http"
	"students-mgmt-api/config"
	"students-mgmt-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllStudents retrieves all students from the database
func GetAllStudents(c *gin.Context) {
	var students []models.Student

	// Fetch students using GORM
	if err := config.DB.Find(&students).Error; err != nil {
		log.Printf("Error fetching students: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}

	c.IndentedJSON(http.StatusOK, students)
}

// AddStudent adds a new student to the database
func AddStudent(c *gin.Context) {
	var st models.Student

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Generate a new UUID for the student
	st.ID = uuid.New().String()

	// Insert into database using GORM
	if err := config.DB.Create(&st).Error; err != nil {
		log.Printf("Error inserting student: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student added!", "student": st})
}

// GetStudent retrieves a single student by ID
func GetStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	// Find student by ID using GORM
	if err := config.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, student)
}

// UpdateStudent updates a student's details
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	// Check if the student exists
	if err := config.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Update student in database
	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student updated", "student": student})
}

// DeleteStudent deletes a student by ID
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	// Delete student using GORM
	if err := config.DB.Where("id = ?", id).Delete(&models.Student{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}

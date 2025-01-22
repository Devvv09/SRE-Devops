package controllers

import (
	"fmt"
	"log"
	"net/http"
	"students-mgmt-api/config"
	"students-mgmt-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var st models.Student

func GetAllStudents(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	rows, err := config.DB.Query("SELECT id, name, age, grade, created_at FROM student")
	if err != nil {
		log.Printf("Error querying students: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	defer rows.Close()
	var students []models.Student
	for rows.Next() {
		var st models.Student
		if err := rows.Scan(&st.ID, &st.Name, &st.Age, &st.Grade, &st.CreatedAt); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process student records"})
			return
		}
		students = append(students, st)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error after iterating rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

func AddStudent(c *gin.Context) {
	if err := c.BindJSON(&st); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	id := uuid.New()
	stmt, err := config.DB.Prepare("INSERT INTO student (id, name, age, grade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Printf("Error preparing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add student"})
		return
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id, st.Name, st.Age, st.Grade); err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student Added!"})

}

func GetStudent(c *gin.Context) {
	id := c.Param("id")
	query := "SELECT id, name, age, grade, created_at FROM student WHERE id=$1"
	results, err := config.DB.Query(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching student"})
		return
	}
	defer results.Close()
	if results.Next() {
		err = results.Scan(&st.ID, &st.Name, &st.Age, &st.Grade, &st.CreatedAt)
		if err != nil {
			fmt.Println("Scan Error:", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning student record"})
			return
		}
		c.IndentedJSON(http.StatusOK, st)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
	}
}
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	if err := c.BindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	query := `UPDATE student SET name = $1, age = $2, grade = $3 WHERE id = $4`
	_, err := config.DB.Exec(query, st.Name, st.Age, st.Grade, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student updated"})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM student WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}

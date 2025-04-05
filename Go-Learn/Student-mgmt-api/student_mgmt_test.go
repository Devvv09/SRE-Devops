package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"students-mgmt-api/src/config"
	"students-mgmt-api/src/controllers"
	"students-mgmt-api/src/routes"
	"students-mgmt-api/src/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var handler *gin.Engine
var err error

func TestMain(m *testing.M) {
	_ = config.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	fmt.Println("Connected to the database:")
	handler = routes.Routes()
	os.Exit(m.Run())
}

func TestAddStudent(t *testing.T) {
	router := gin.Default()
	router.POST("/api/v1/students", controllers.AddStudent)

	student := models.Student{
		Name:  "Devv",
		Age:   21,
		Grade: "A",
	}

	body, _ := json.Marshal(student)
	req, _ := http.NewRequest("POST", "http://localhost:3000/api/v1/students", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetAllStudents(t *testing.T) {
	router := gin.Default()
	router.GET("/students", controllers.GetAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetStudent_NotFound(t *testing.T) {
	router := gin.Default()
	router.GET("/students/:id", controllers.GetStudent)

	req, _ := http.NewRequest("GET", "/students/non-existent-id", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestUpdateStudent_NotFound(t *testing.T) {
	router := gin.Default()
	router.PUT("/students/:id", controllers.UpdateStudent)

	student := models.Student{
		Name:  "Devv",
		Age:   21,
		Grade: "A",
	}
	body, _ := json.Marshal(student)
	req, _ := http.NewRequest("PUT", "/students/non-existent-id", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestDeleteStudent(t *testing.T) {
	router := gin.Default()
	router.DELETE("/students/:id", controllers.DeleteStudent)

	req, _ := http.NewRequest("DELETE", "/students/non-existent-id", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Contains(t, []int{http.StatusOK, http.StatusInternalServerError}, resp.Code)
}

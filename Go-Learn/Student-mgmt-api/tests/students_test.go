package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"students-mgmt-api/src/routes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllStudents(t *testing.T) {
	router := routes.Routes()

	// Create a test HTTP request
	req, _ := http.NewRequest("GET", "/students", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"students":`) // Check if response has students
}


func TestGetStudentByID(t *testing.T) {
	router := routes.Routes()

	// Test existing student
	req, _ := http.NewRequest("GET", "/students/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"id":1`)

	// Test non-existent student
	req, _ = http.NewRequest("GET", "/students/999", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}


func TestCreateStudent(t *testing.T) {
	router := routes.Routes()

	// Valid student
	payload := `{"name": "Alice", "email": "alice@example.com"}`
	req, _ := http.NewRequest("POST", "/students", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"Alice"`)

	// Invalid student (missing name)
	payload = `{"email": "invalid@example.com"}`
	req, _ = http.NewRequest("POST", "/students", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}


func TestUpdateStudent(t *testing.T) {
	router := routes.Routes()

	// Update existing student
	payload := `{"name": "Bob Updated", "email": "bob@updated.com"}`
	req, _ := http.NewRequest("PUT", "/students/1", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"Bob Updated"`)

	// Update non-existent student
	req, _ = http.NewRequest("PUT", "/students/999", strings.NewReader(payload))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}


func TestDeleteStudent(t *testing.T) {
	router := routes.Routes()

	// Delete existing student
	req, _ := http.NewRequest("DELETE", "/students/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Delete non-existent student
	req, _ = http.NewRequest("DELETE", "/students/999", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
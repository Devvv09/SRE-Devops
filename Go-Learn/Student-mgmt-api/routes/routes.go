package routes

import (
	"students-mgmt-api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {

	r := gin.Default()

	r.GET("/api/v1/students", controllers.GetAllStudents)

	r.GET("/api/v1/students/:id", controllers.GetStudent)

	r.POST("/api/v1/students", controllers.AddStudent )

	r.PUT("/api/v1/students/:id",controllers.UpdateStudent )

	r.DELETE("/api/v1/students/:id",controllers.DeleteStudent)

	return r

}

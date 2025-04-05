package routes

import (
	"students-mgmt-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {

	r := gin.Default()

	apiGroup := r.Group("/api/v1")
	{
		apiGroup.GET("/students", controllers.GetAllStudents)
		apiGroup.GET("/students/:id", controllers.GetStudent)
		apiGroup.POST("/students", controllers.AddStudent)
		apiGroup.PUT("/students/:id", controllers.UpdateStudent)
		apiGroup.DELETE("/students/:id", controllers.DeleteStudent)
	}
	return r
}

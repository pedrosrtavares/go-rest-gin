package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrosrtavares/go-rest-gin/database"
	"github.com/pedrosrtavares/go-rest-gin/models"
)

func GetAllStudents(c *gin.Context) {
	c.Header("Content-Type", "text/json")
	// TODO - FIX THIS ROUTE
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func RegisterNewStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)

}

// func FilterStudentsById(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	c.JSON(200, gin.H{

// 	})
// }

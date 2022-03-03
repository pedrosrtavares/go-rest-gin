package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrosrtavares/go-rest-gin/database"
	"github.com/pedrosrtavares/go-rest-gin/models"
)

func GetAllStudents(c *gin.Context) {
	c.Header("Content-Type", "text/json")
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

func FindStudentById(c *gin.Context) {
	var aluno models.Student
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Could not find a student with this id in the database",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeleteStudent(c *gin.Context) {
	var aluno models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, aluno)
}

func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func SearchStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Query("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, &student)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedrosrtavares/go-rest-gin/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetAllStudents)
	r.GET("/aluno/:id", controllers.FindStudentById)
	r.POST("/alunos", controllers.RegisterNewStudent)
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	r.Run()
}

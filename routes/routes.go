package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedrosrtavares/go-rest-gin/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetAllStudents)
	//Not Implemented \/
	// r.GET("/aluno/{id}")
	r.POST("/alunos", controllers.RegisterNewStudent)
	r.Run()
}

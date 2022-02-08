package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pedrosrtavares/go-rest-gin/models"
)

func GetAllStudents(c *gin.Context) {
	c.Header("Content-Type", "text/json")
	c.JSON(200, models.Students)
}

// func FilterStudentsById(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	c.JSON(200, gin.H{

// 	})
// }

package database

import (
	"log"

	"github.com/pedrosrtavares/go-rest-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	stringDeConexao := "host=172.18.0.2 port=5432 user=root password=root dbname=root"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Student{})
}

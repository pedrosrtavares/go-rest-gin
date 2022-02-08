package main

import (
	"github.com/pedrosrtavares/go-rest-gin/models"
	"github.com/pedrosrtavares/go-rest-gin/routes"
)

func main() {
	models.Students = []models.Student{
		{Name: "Pedro Tavares", CPF: "000.000.000-00", RG: "00.000.000-0"},
		{Name: "Darth Vader", CPF: "000.000.000-00", RG: "00.000.000-0"},
		{Name: "Master Yoda", CPF: "000.000.000-00", RG: "00.000.000-0"},
	}
	routes.HandleRequest()
}

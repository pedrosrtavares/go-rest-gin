package main

import (
	"github.com/pedrosrtavares/go-rest-gin/database"
	"github.com/pedrosrtavares/go-rest-gin/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequest()
}

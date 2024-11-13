package main

import (
	"restapi/api-test/routes"
	"restapi/db"

	"github.com/gin-gonic/gin"
)

func main() {
	//fmt.Println("33")
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}

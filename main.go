package main

import (
	"example.com/udemy_course/db"
	"example.com/udemy_course/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	db.InitDB()

	server := gin.Default()

	routes.InitializeRoute(server)

	server.Run(":8080")
}


package main

import (
	"example.com/event-booking-rest-api/db"
	"example.com/event-booking-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}

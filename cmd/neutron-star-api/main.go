package main

import (
	"log"
	"neutron-star-api/internal/database"
	"neutron-star-api/internal/server"
)

// @title           Neutron Star API
// @version         1.0
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        token

func main() {
	log.Println("neutron-star-router start")
	database.Connect()
	defer database.Close()
	server.Run()
}

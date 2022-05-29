package main

import (
	"log"
	"neutron-star-api/internal/database"
	"neutron-star-api/internal/server"
)

func main() {
	log.Println("neutron-star-api start")
	database.Connect()
	defer database.Close()
	server.Run()
}

package main

import (
	"log"
	"simple-api/config"
	"simple-api/data"
	"simple-api/database"
	"simple-api/handler"
	"simple-api/server"
)

func main() {
	configPath := "./config.json"
	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDatabase(*config)
	if err != nil {
		log.Fatal(err)
	}

	repo := data.NewProductRepository(db)

	handler := handler.NewHandler(repo)
	handler.Init()

	server := server.NewServer(config, handler.Router)

	log.Fatalln(server.Run())
}

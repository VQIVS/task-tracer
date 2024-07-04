package main

import (
	"fmt"
	"net/http"
	"task-tracker/api"
	"task-tracker/config"
	"task-tracker/pkg/database"
)

func main() {
	cfg, err := config.ReadConfig("./config") // Assuming your config directory is here
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	db, err := database.Init(cfg.DatabaseURL)
	if err != nil {
		panic(fmt.Errorf("database initialization failed: %s", err))
	}

	router := api.SetupRouter(db) // Setup router with database dependency

	fmt.Println("Starting server on port", cfg.ServerPort)
	http.ListenAndServe(":"+cfg.ServerPort, router)
}

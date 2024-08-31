package main

import (
	"fmt"
	models "github.com/BerkatPS/internal"
	"github.com/BerkatPS/internal/database"
	server2 "github.com/BerkatPS/internal/server"
	"github.com/BerkatPS/pkg/config"
	"log"
	"net/http"
)

func main() {

	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

	}
	fmt.Println("Connected to database")

	server := server2.NewServer(db)

	err = database.AutoMigrate(db,
		&models.Document{},
		&models.User{},
		&models.QualityCheck{},
		&models.Expense{},
		&models.Project{},
		&models.SafetyIncident{},
		&models.Task{},
		&models.Message{},
		&models.Report{},
		&models.Presence{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	fmt.Println("Auto migrated tables")

	log.Printf("starting server on %s", cfg.ServerAddress)

	if err := http.ListenAndServe(cfg.ServerAddress, server.Router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

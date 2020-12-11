package main

import (
	"log"
	"net/http"
	"quizy/database"
	"quizy/models"
	"quizy/routes"
	"time"
)

func main() {
	// Migrate()
	// setup server
	server := &http.Server{
		Addr:         ":8081",
		Handler:      routes.Router(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func Migrate() {
	con := database.Connect()
	db, _ := con.DB()
	defer db.Close()

	con.AutoMigrate(
		&models.User{},
		&models.Quiz{},
		&models.Question{},
		&models.Answer{},
	)
}

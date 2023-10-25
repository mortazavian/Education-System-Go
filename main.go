package main

import (
	"Education-System-Go/db_conections"
	"Education-System-Go/handlers"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

var db *sql.DB

func main() {

	db = db_conections.NewPostgres01()

	app := fiber.New()

	app.Get("/student/:id", handlers.HandleGetUser)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}

package main

// get all student/ teachers
// return student of a teacher
// return teacher of a student

import (
	"Education-System-Go/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	//db = db_conections.NewPostgres01()

	app := fiber.New()

	// Student Based APIs
	app.Get("/student/:id", handlers.HandleGetStudent)
	app.Get("/student", handlers.HandleGetStudents)

	// Teacher Based APIs
	app.Get("/teacher/:id", handlers.HandleGetTeacher)
	app.Get("/teacher", handlers.HandleGetTeachers)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}

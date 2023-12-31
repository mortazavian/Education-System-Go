package main

// get all student/ teachers --> DONE
// return student of a teacher --> DONE
// return teacher of a student --> DONE
// Update user information
// Login
// Reset password

import (
	"Education-System-Go/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// Student Based APIs
	app.Get("/student/:id", handlers.HandleGetStudent)
	app.Get("/student", handlers.HandleGetStudents)
	app.Get("/student-of/:id", handlers.HandleGetTeacherOfAStudent)
	app.Post("/student", handlers.HandlePostStudent)
	app.Put("/student/:id", handlers.HandlePutStudent)

	// Teacher Based APIs
	app.Get("/teacher/:id", handlers.HandleGetTeacher)
	app.Get("/teacher", handlers.HandleGetTeachers)
	app.Get("/teacher-of/:id", handlers.HandleGetStudentsByTeacherId)
	app.Post("teacher", handlers.HandlePostTeacher)
	app.Put("/teacher/:id", handlers.HandlePutTeacher)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}

package handlers

import (
	"Education-System-Go/database"
	"Education-System-Go/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func HandleGetStudent(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	fmt.Println("the id is:", id)

	student, err := database.GetUserByID(int64(id))
	if err != nil {
		return err
	}

	err = c.JSON(student)
	if err != nil {
		return err
	}

	return err
}

func HandleGetUsers(c *fiber.Ctx) error {
	var students []models.Student

	students, err := database.GetStudents()
	if err != nil {
		return err
	}

	err = c.JSON(students)
	if err != nil {
		return err
	}

	return err
}

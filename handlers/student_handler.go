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

	student, err := database.GetStudentByID(int64(id))
	if err != nil {
		return err
	}

	err = c.JSON(student)
	if err != nil {
		return err
	}

	return err
}

func HandleGetStudents(c *fiber.Ctx) error {
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

func HandleGetTeacherOfAStudent(c *fiber.Ctx) error {
	var teacher models.Teacher
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	teacher, err = database.GetTeacherByStudentId(int64(id))

	err = c.JSON(teacher)
	if err != nil {
		return err
	}
	return nil
}

func HandleputStudent(c *fiber.Ctx) error {

}

package handlers

import (
	"Education-System-Go/database"
	"Education-System-Go/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func HandleGetTeacher(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	fmt.Println("the id is:", id)

	teacher, err := database.GetTeacherByID(int64(id))
	if err != nil {
		return err
	}

	err = c.JSON(teacher)
	if err != nil {
		return err
	}

	return nil
}

func HandleGetTeachers(c *fiber.Ctx) error {
	var students []models.Teacher

	students, err := database.GetTeachers()
	if err != nil {
		return err
	}

	err = c.JSON(students)
	if err != nil {
		return err
	}

	return nil

}

func HandleGetStudentsByTeacherId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	studentsOfTheTeacher, err := database.GetStudentsByTeacherId(int64(id))

	if err != nil {
		return err
	}
	err = c.JSON(studentsOfTheTeacher)
	if err != nil {
		return err
	}

	return nil

}

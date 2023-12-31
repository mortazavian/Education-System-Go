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

func HandlePostTeacher(c *fiber.Ctx) error {
	var teacher models.Teacher
	if err := c.BodyParser(&teacher); err != nil {
		return err
	}

	fmt.Printf("%+v", teacher)

	if errors := teacher.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}

	insertedTeacher, err := database.PostTeacher(teacher)
	if err != nil {
		return err
	}
	return c.JSON(insertedTeacher)
}

func HandlePutTeacher(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	var teacher models.Teacher
	err = c.BodyParser(&teacher)
	if err != nil {
		return err
	}

	updatedStudent, err := database.PutTeacher(id, teacher)
	if err != nil {
		return err
	}

	_ = updatedStudent

	return nil
}

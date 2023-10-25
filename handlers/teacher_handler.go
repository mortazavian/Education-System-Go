package handlers

import (
	"Education-System-Go/database"
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

	return err
}

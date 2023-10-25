package handlers

import (
	"Education-System-Go/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func HandleGetUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	fmt.Println("the id is:", id)

	student, err := database.GetUserByID(int64(id))
	if err != nil {
		return err
	}

	//res, err := json.Marshal(student)
	//if err != nil {
	//	return err
	//}

	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(res)

	err = c.JSON(student)
	if err != nil {
		return err
	}

	return err
}

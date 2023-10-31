package models

import (
	"fmt"
)

type Teacher struct {
	Name     string
	LastName string
	Email    string
	Password string
}

func (teacher Teacher) Validate() map[string]string {
	fmt.Printf("%+v", teacher)
	errors := map[string]string{}
	if len(teacher.Name) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("firstName length should be at lease %d characters", minFirstNameLen)
	}
	if len(teacher.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("lastName length should be at lease %d characters", minLastNameLen)
	}
	if len(teacher.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at lease %d characters", minPasswordLen)
	}
	if !isEmailValid(teacher.Email) {
		errors["email"] = fmt.Sprintf("email is invalid")
	}
	return errors
}

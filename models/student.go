package models

import (
	"fmt"
	"regexp"
)

type Student struct {
	Name      string
	LastName  string
	Email     string
	Password  string
	TeacherId int64
}

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 7
)

func (student Student) Validate() map[string]string {
	fmt.Printf("%+v", student)
	errors := map[string]string{}
	if len(student.Name) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("firstName length should be at lease %d characters", minFirstNameLen)
	}
	if len(student.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("lastName length should be at lease %d characters", minLastNameLen)
	}
	if len(student.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at lease %d characters", minPasswordLen)
	}
	if !isEmailValid(student.Email) {
		errors["email"] = fmt.Sprintf("email is invalid")
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

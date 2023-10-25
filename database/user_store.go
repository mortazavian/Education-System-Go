package database

import (
	"Education-System-Go/db_conections"
	"Education-System-Go/models"
	"database/sql"
	"fmt"
)

var db *sql.DB

func GetUserByID(id int64) (models.Student, error) {
	db = db_conections.NewPostgres01()
	query := `
	SELECT *
	FROM students
	WHERE id =  $1
`
	rows, err := db.Query(query, id)
	if err != nil {
		return models.Student{}, err
	}
	defer rows.Close()

	var student models.Student
	for rows.Next() {

		err := rows.Scan(&id, &student.Name, &student.LastName, &student.TeacherId, &student.Email, &student.Password)
		if err != nil {
			return models.Student{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return models.Student{}, err
	}

	fmt.Printf("%+v", student)

	return student, nil
}

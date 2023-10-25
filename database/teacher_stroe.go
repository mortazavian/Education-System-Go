package database

import (
	"Education-System-Go/db_conections"
	"Education-System-Go/models"
	"fmt"
)

func GetTeacherByID(id int64) (models.Teacher, error) {
	db = db_conections.NewPostgres01()
	query := `
	SELECT *
	FROM teachers
	WHERE id =  $1
`
	rows, err := db.Query(query, id)
	if err != nil {
		return models.Teacher{}, err
	}
	defer rows.Close()

	var teacher models.Teacher
	for rows.Next() {

		err := rows.Scan(&id, &teacher.Name, &teacher.LastName, &teacher.Email, &teacher.Password)
		if err != nil {
			return models.Teacher{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return models.Teacher{}, err
	}

	fmt.Printf("%+v", teacher)

	return teacher, nil
}

func GetTeachers() ([]models.Teacher, error) {
	db = db_conections.NewPostgres01()
	query := `
	SELECT *
	FROM teachers`

	rows, err := db.Query(query)
	if err != nil {
		return []models.Teacher{}, err
	}
	defer rows.Close()

	var (
		teachers []models.Teacher
		teacher  models.Teacher
		id       int
	)
	for rows.Next() {

		err := rows.Scan(&id, &teacher.Name, &teacher.LastName, &teacher.Email, &teacher.Password)
		if err != nil {
			return []models.Teacher{}, err
		}
		teachers = append(teachers, teacher)

		if err := rows.Err(); err != nil {
			return []models.Teacher{}, err
		}
	}

	fmt.Printf("%+v", teachers)

	return teachers, nil

}

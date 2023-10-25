package database

import (
	"Education-System-Go/db_conections"
	"Education-System-Go/models"
	"database/sql"
)

var db *sql.DB

func GetStudentByID(id int64) (models.Student, error) {
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

	return student, nil
}

func GetStudents() ([]models.Student, error) {
	db = db_conections.NewPostgres01()
	query := `
	SELECT *
	FROM students`

	rows, err := db.Query(query)
	if err != nil {
		return []models.Student{}, err
	}
	defer rows.Close()

	var (
		students []models.Student
		student  models.Student
		id       int
	)
	for rows.Next() {

		err := rows.Scan(&id, &student.Name, &student.LastName, &student.TeacherId, &student.Email, &student.Password)
		if err != nil {
			return []models.Student{}, err
		}
		students = append(students, student)

		if err := rows.Err(); err != nil {
			return []models.Student{}, err
		}
	}

	return students, nil
}

func GetTeacherByStudentId(id int64) (models.Teacher, error) {
	db = db_conections.NewPostgres01()
	// Query to get the data of the student
	query := `
	SELECT *
	FROM students
	WHERE id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		return models.Teacher{}, err
	}
	var student models.Student
	for rows.Next() {

		err := rows.Scan(&id, &student.Name, &student.LastName, &student.TeacherId, &student.Email, &student.Password)
		if err != nil {
			return models.Teacher{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return models.Teacher{}, err
	}

	var teacherId int64 = student.TeacherId

	query = `
	SELECT *
	FROM teachers
	WHERE id = $1;
`
	rows, err = db.Query(query, teacherId)

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

	return teacher, nil
}

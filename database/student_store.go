package database

import (
	"Education-System-Go/db_conections"
	"Education-System-Go/models"
	"database/sql"
	"fmt"
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

func PostStudent(student models.Student) (models.Student, error) {
	db = db_conections.NewPostgres01()
	studentInsertQuery := "Insert INTO students (name, last_name, teacher_fk, email, password)  values ($1, $2, $3, $4,$5) ;"
	_, err := db.Exec(studentInsertQuery, student.Name, student.LastName, student.TeacherId, student.Email, student.Password)
	if err != nil {
		return models.Student{}, err
	}
	return student, nil
}

func PutStudent(id int, student models.Student) (models.Student, error) {
	db = db_conections.NewPostgres01()
	getTheStudentQuery := `Select * from students where id = $1`
	rows, err := db.Query(getTheStudentQuery, id)
	if err != nil {
		return models.Student{}, err
	}
	var studentToUpdate models.Student
	for rows.Next() {

		err := rows.Scan(&id, &studentToUpdate.Name, &studentToUpdate.LastName, &studentToUpdate.TeacherId, &studentToUpdate.Email, &studentToUpdate.Password)
		if err != nil {
			return models.Student{}, err
		}
	}

	UpdateStudentInformation(&studentToUpdate, student)

	fmt.Printf("%+v", studentToUpdate)

	putStudentQuery := "update students set name = $1, last_name = $2, teacher_fk = $3, email = $4,password = $5 where id = $6"

	_, err = db.Exec(putStudentQuery, studentToUpdate.Name, studentToUpdate.LastName, studentToUpdate.TeacherId, studentToUpdate.Email, studentToUpdate.Password, id)
	if err != nil {
		return models.Student{}, err
	}

	fmt.Printf("%+v", studentToUpdate)

	return studentToUpdate, err
}

func UpdateStudentInformation(oldStudent *models.Student, newStudent models.Student) {
	if newStudent.Name != "" {
		oldStudent.Name = newStudent.Name
	}
	if newStudent.LastName != "" {
		oldStudent.LastName = newStudent.LastName
	}
	if newStudent.Email != "" {
		oldStudent.Email = newStudent.Email
	}
	if newStudent.Password != "" {
		oldStudent.Password = newStudent.Password
	}
}

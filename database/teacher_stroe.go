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

func GetStudentsByTeacherId(id int64) ([]models.Student, error) {
	db = db_conections.NewPostgres01()
	query := `
	SELECT *
	FROM students
	WHERE teacher_fk =  $1
`
	rows, err := db.Query(query, id)
	if err != nil {
		return []models.Student{}, err
	}
	defer rows.Close()

	var (
		students []models.Student
		student  models.Student
	)

	for rows.Next() {

		err := rows.Scan(&id, &student.Name, &student.LastName, &student.TeacherId, &student.Email, &student.Password)
		if err != nil {
			return []models.Student{}, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return []models.Student{}, err
	}

	return students, nil

}

func PostTeacher(teacher models.Teacher) (models.Teacher, error) {
	db = db_conections.NewPostgres01()
	teacherInsertQuery := "Insert INTO teachers (name, last_name, email, password)  values ($1, $2, $3, $4) ;"
	_, err := db.Exec(teacherInsertQuery, teacher.Name, teacher.LastName, teacher.Email, teacher.Password)
	if err != nil {
		return models.Teacher{}, err
	}
	return teacher, nil
}

func PutTeacher(id int, teacher models.Teacher) (models.Teacher, error) {
	db = db_conections.NewPostgres01()
	getTheTeacherQuery := `Select * from teachers where id = $1`
	rows, err := db.Query(getTheTeacherQuery, id)
	if err != nil {
		return models.Teacher{}, err
	}
	var teacherToUpdate models.Teacher
	for rows.Next() {

		err := rows.Scan(&id, &teacherToUpdate.Name, &teacherToUpdate.LastName, &teacherToUpdate.Email, &teacherToUpdate.Password)
		if err != nil {
			return models.Teacher{}, err
		}
	}

	UpdateTeacherInformation(&teacherToUpdate, teacher)

	fmt.Printf("%+v", teacherToUpdate)

	putTeacherQuery := "update teachers set name = $1, last_name = $2, email = $3,password = $4 where id = $5"

	_, err = db.Exec(putTeacherQuery, teacherToUpdate.Name, teacherToUpdate.LastName, teacherToUpdate.Email, teacherToUpdate.Password, id)
	if err != nil {
		return models.Teacher{}, err
	}

	fmt.Printf("%+v", teacherToUpdate)

	return teacherToUpdate, err
}

func UpdateTeacherInformation(oldTeacher *models.Teacher, newTeacher models.Teacher) {
	if newTeacher.Name != "" {
		oldTeacher.Name = newTeacher.Name
	}
	if newTeacher.LastName != "" {
		oldTeacher.LastName = newTeacher.LastName
	}
	if newTeacher.Email != "" {
		oldTeacher.Email = newTeacher.Email
	}
	if newTeacher.Password != "" {
		oldTeacher.Password = newTeacher.Password
	}
}

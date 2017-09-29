package persistent

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"task2/student-server/model"
	"log"
)

type StudentStorage struct {
	db sql.DB
}

func NewStudentStorage(db *sql.DB) *StudentStorage {
	return &StudentStorage{*db}
}

func (s *StudentStorage) GetStudent(id string) (model.Student, bool) {
	getByIdPreparation, err := s.db.Prepare("SELECT * FROM students where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer getByIdPreparation.Close()

	student := model.Student{}
	err = getByIdPreparation.QueryRow(id).Scan(&student.Id, &student.Name, &student.Score)

	if err != nil {
		if err == sql.ErrNoRows {
			return student, false
		} else {
			log.Fatal(err)
		}
	}

	return student, true
}

func (s *StudentStorage) GetStudents() (students []model.Student) {
	students = make([]model.Student, 0) // prevent body {null} in response. It causes exception on android client

	getAllPreparation, err := s.db.Prepare("SELECT * FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer getAllPreparation.Close()

	rows, err := getAllPreparation.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		student := model.Student{}
		rows.Scan(&student.Id, &student.Name, &student.Score)

		if err != nil {
			log.Fatal(err)
		}

		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func (s *StudentStorage) SaveStudent(student model.Student) {
	insertPreparation, err := s.db.Prepare("INSERT INTO students VALUES( ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer insertPreparation.Close()

	_, err = insertPreparation.Exec(student.Id, student.Name, student.Score)
	if err != nil {
		panic(err.Error())
	}
}

func (s *StudentStorage) UpdateStudent(student model.Student) {
	updatePreparation, err := s.db.Prepare("UPDATE students SET id=?, name=?, score=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer updatePreparation.Close()

	_, err = updatePreparation.Exec(student.Id, student.Name, student.Score, student.Id)
	if err != nil {
		panic(err.Error())
	}
}

func (s *StudentStorage) DeleteStudent(id string) {
	deletePreparation, err := s.db.Prepare("DELETE FROM students WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer deletePreparation.Close()

	_, err = deletePreparation.Exec(id)
	if err != nil {
		panic(err.Error())
	}
}

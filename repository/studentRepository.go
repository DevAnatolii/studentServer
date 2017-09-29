package repository

import "task2/student-server/model"

type StudentRepository interface {
	GetStudent(id string) (model.Student, bool)
	GetStudents() []model.Student
	SaveStudent(s model.Student)
	UpdateStudent(s model.Student)
	DeleteStudent(id string)
}

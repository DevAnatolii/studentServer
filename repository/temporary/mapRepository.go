package temporary

import "task2/student-server/model"

type StudentStorage struct {
	storage map[string]model.Student
}

func NewStudentStorage() *StudentStorage {
	return &StudentStorage{
		storage: make(map[string]model.Student),
	}
}

func (s *StudentStorage) GetStudent(id string) (student model.Student, ok bool) {
	student, ok = s.storage[id]
	return
}

func (s *StudentStorage) GetStudents() (students []model.Student) {
	students = make([]model.Student, 0) // prevent body {null} in response. It causes exception on android client
	for _, val := range s.storage {
		students = append(students, val)
	}
	return
}

func (s *StudentStorage) SaveStudent(student model.Student) {
	s.storage[student.Id] = student
}

func (s *StudentStorage) UpdateStudent(student model.Student) {
	s.storage[student.Id] = student
}

func (s *StudentStorage) DeleteStudent(id string) {
	delete(s.storage, id)
}

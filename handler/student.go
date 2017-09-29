package handler

import (
	"net/http"
	"task2/student-server/repository"
	"encoding/json"
	"task2/student-server/model"
	"task2/student-server/generator"
	"fmt"
)

const HandlePath = "/students/"

type studentHandler struct {
	storage repository.StudentRepository
}

func NewStudentHandler(repository repository.StudentRepository) *studentHandler {
	return &studentHandler{repository}
}

func (sh *studentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.Method {
	case http.MethodGet:
		sh.getStudents(w, r)
	case http.MethodPost:
		sh.createStudent(w, r)
	case http.MethodPut:
		sh.updateStudent(w, r)
	case http.MethodDelete:
		sh.deleteStudent(w, r)
	}
}

func (sh *studentHandler) getStudents(w http.ResponseWriter, r *http.Request) {
	writeHeaders(w)
	encoder := json.NewEncoder(w)
	students := sh.storage.GetStudents()
	fmt.Print(students)
	err := encoder.Encode(students)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (sh *studentHandler) createStudent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var student model.Student
	err := decoder.Decode(&student)

	fmt.Print(student)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	student.Id = generator.RandomId()
	sh.storage.SaveStudent(student)
	writeHeaders(w)
	encoder := json.NewEncoder(w)
	encoder.Encode(student)
	w.WriteHeader(http.StatusCreated)
}

func (sh *studentHandler) updateStudent(w http.ResponseWriter, r *http.Request) {
	id := obtainId(r)

	if len(id) == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var student model.Student
	err := decoder.Decode(&student)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	if _, ok := sh.storage.GetStudent(id); !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	student.Id = id
	sh.storage.UpdateStudent(student)
	writeHeaders(w)
	encoder := json.NewEncoder(w)
	encoder.Encode(student)
	w.WriteHeader(http.StatusOK)
}

func (sh *studentHandler) deleteStudent(w http.ResponseWriter, r *http.Request) {
	id := obtainId(r)

	if len(id) == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	sh.storage.DeleteStudent(id)
	writeHeaders(w)
	w.WriteHeader(http.StatusNoContent)
}

func obtainId(r *http.Request) string {
	//because endpoint ends with "/" symbol
	return r.URL.Path[len(HandlePath): len(r.URL.Path)-1]
}

func writeHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

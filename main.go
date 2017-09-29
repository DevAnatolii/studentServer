package main

import (
	"net/http"
	"task2/student-server/handler"
	"task2/student-server/repository/temporary"
	"task2/student-server/repository"
	"database/sql"
	"log"
	"task2/student-server/repository/persistent"
	"flag"
	"strings"
)

const (
	storageFlag       = "storage"
	storagePersistent = "persistent"
	storageTemporary  = "temporary"
	serverAddress     = ":8080"
)

func main() {
	serveMux := http.NewServeMux()

	studentsHandler := handler.NewStudentHandler(createStorage())
	serveMux.Handle(handler.HandlePath, studentsHandler)

	http.ListenAndServe(serverAddress, serveMux)
}

func createStorage() (storage repository.StudentRepository) {
	if isPersistentStorage() {
		db, err := sql.Open("mysql",
			"root:0000@tcp(localhost:3306)/university")
		if err != nil {
			log.Fatal(err)
		}

		storage = persistent.NewStudentStorage(db)
	} else {
		storage = temporary.NewStudentStorage()
	}

	return
}

func isPersistentStorage() (ok bool) {
	storage := flag.String(storageFlag, "", "determines type of storage")
	flag.Parse()
	if strings.EqualFold(strings.ToLower(*storage), storagePersistent) {
		ok = true
	} else if strings.EqualFold(strings.ToLower(*storage), storageTemporary) {
		ok = false
	} else {
		log.Println("Type of storage is not specified. Persistent one is used by default")
		ok = true
	}

	return
}

package main

import (
	"database/sql"
	"enCodeTest/src/Repositories"
	"enCodeTest/src/handlers"
	"enCodeTest/src/services"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	connStr := "postgres://postgres:364678x@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	repo := Repositories.NewPostgresRepository(db)
	service := services.NewPersonService(repo)
	handler := handlers.NewPersonHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/person", handler.GetPersons).Methods("GET")
	router.HandleFunc("/person/{id}", handler.GetPerson).Methods("GET")
	router.HandleFunc("/person", handler.CreatePerson).Methods("POST")
	router.HandleFunc("/person/{id}", handler.UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{id}", handler.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

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

//Project Евгений Козлов

func main() {

	/*Реализация с GORM
	dsn := "host=localhost user=postgres password=364678x dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Person{})
	repo := Repositories.NewGormPersonRepository(db)*/

	connStr := "postgres://postgres:364678x@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Автоматическая миграция (создание таблиц)

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

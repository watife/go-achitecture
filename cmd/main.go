package main

import (
	"database/sql"
	"fakorede-bolu/go-ach/cmd/database/psql"
	"fakorede-bolu/go-ach/cmd/routes"
	"fakorede-bolu/go-ach/cmd/todo"
	"fakorede-bolu/go-ach/pkg/logs"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	var todoRepo todo.Repository

	todoRepo = psql.NewPostgresTodoRepository(openDB("postgresql://postgres@localhost/test?sslmode=disable"))

	// Services
	service := todo.NewService(todoRepo)

	// Routes
	r := routes.NewRouter(service)

	srv := &http.Server{
		Addr:           ":4000",
		ErrorLog:       logs.ErrorLog,
		Handler:        r,
		IdleTimeout:    time.Minute,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 524288,
	}

	logs.InfoLog.Printf("Starting server on %s", ":4000")

	err := srv.ListenAndServe()

	logs.ErrorLog.Fatal(err)
}

func openDB(database string) *sql.DB {
	fmt.Println("Connecting to PostgreSQL DB")
	db, err := sql.Open("postgres", database)
	if err != nil {
		fmt.Println("failed")
		log.Fatalf("%s", err)
		panic(err)
	}
	return db
}

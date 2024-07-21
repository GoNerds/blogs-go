package main

import (
	"log"
	"net/http"

	"blogs-go/internal/db"
	"blogs-go/internal/routes"
)

func main() {
	//Initialize the database

	db.Connect("mongodb://localhost:27017", "blogs-go")
	defer db.Close()

	r := routes.InitializeRoutes(db.GetDatabase())

	http.Handle("/", r)
	log.Println("Server started on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
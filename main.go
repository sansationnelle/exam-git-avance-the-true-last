package main

import (
	"log"
	"main/handlers"
	"net/http"
)

func main() {
	log.Println("Serveur Go en démarrage...")

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/book", handlers.BookHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

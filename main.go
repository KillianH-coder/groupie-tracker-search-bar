package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)

	http.HandleFunc("/artist/", handlers.ArtistHandler)

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Started Server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server Start Error:", err)
	}
}

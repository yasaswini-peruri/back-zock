package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"productmanagement/handlers"
)

func main() {
	router := mux.NewRouter()

	// API Endpoints
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET")
	router.HandleFunc("/products", handlers.GetProductsByUserID).Methods("GET")

	// Start Server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

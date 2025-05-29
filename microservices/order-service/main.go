package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"order-service/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", handlers.GetOrder).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5002"
	}

	log.Printf("Order service is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

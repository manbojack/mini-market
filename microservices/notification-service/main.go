package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "notification-service/handlers"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/notify", handlers.SendNotification).Methods("POST")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Notification Service running on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}

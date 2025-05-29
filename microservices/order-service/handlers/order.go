package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"fmt"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Order struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	ProductID int   `json:"product_id"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"`
}

func connectDB() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URI")
	return sql.Open("postgres", dbURL)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = db.QueryRow(
	    "INSERT INTO orders (user_id, product_id, quantity, status) VALUES ($1, $2, $3, $4) RETURNING id",
	    order.UserID, order.ProductID, order.Quantity, "created",
	).Scan(&order.ID)
	if err != nil {
	    http.Error(w, fmt.Sprintf("DB insert error: %v", err), http.StatusInternalServerError)
	    return
	}

	json.NewEncoder(w).Encode(order)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var order Order
	err = db.QueryRow("SELECT id, user_id, product_id, quantity, status FROM orders WHERE id=$1", id).
		Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Status)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

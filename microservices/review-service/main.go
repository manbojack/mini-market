package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"

    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

type Review struct {
    ID        int    `json:"id"`
    ProductID int    `json:"product_id"`
    UserID    int    `json:"user_id"`
    Rating    int    `json:"rating"`
    Comment   string `json:"comment"`
    CreatedAt string `json:"created_at,omitempty"`
    UpdatedAt string `json:"updated_at,omitempty"`
}

var db *sql.DB

func main() {
    var err error
    connStr := "host=postgres user=user password=password dbname=mini_market sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    router := mux.NewRouter()
    router.HandleFunc("/reviews", createReview).Methods("POST")
    router.HandleFunc("/reviews/{product_id}", getReviewsByProduct).Methods("GET")
    router.HandleFunc("/reviews/{id}", updateReview).Methods("PUT")
    router.HandleFunc("/reviews/{id}", deleteReview).Methods("DELETE")

    fmt.Println("Starting review-service on :8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func createReview(w http.ResponseWriter, r *http.Request) {
    var rev Review
    err := json.NewDecoder(r.Body).Decode(&rev)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if rev.Rating < 1 || rev.Rating > 5 {
        http.Error(w, "Rating must be between 1 and 5", http.StatusBadRequest)
        return
    }

    query := `INSERT INTO reviews (product_id, user_id, rating, comment) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
    err = db.QueryRow(query, rev.ProductID, rev.UserID, rev.Rating, rev.Comment).Scan(&rev.ID, &rev.CreatedAt, &rev.UpdatedAt)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(rev)
}

func getReviewsByProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    productID, err := strconv.Atoi(vars["product_id"])
    if err != nil {
        http.Error(w, "Invalid product_id", http.StatusBadRequest)
        return
    }

    rows, err := db.Query("SELECT id, product_id, user_id, rating, comment, created_at, updated_at FROM reviews WHERE product_id=$1", productID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    reviews := []Review{}
    for rows.Next() {
        var rev Review
        err := rows.Scan(&rev.ID, &rev.ProductID, &rev.UserID, &rev.Rating, &rev.Comment, &rev.CreatedAt, &rev.UpdatedAt)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        reviews = append(reviews, rev)
    }
    json.NewEncoder(w).Encode(reviews)
}

func updateReview(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid review id", http.StatusBadRequest)
        return
    }

    var rev Review
    err = json.NewDecoder(r.Body).Decode(&rev)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if rev.Rating < 1 || rev.Rating > 5 {
        http.Error(w, "Rating must be between 1 and 5", http.StatusBadRequest)
        return
    }

    query := `UPDATE reviews SET rating=$1, comment=$2, updated_at=NOW() WHERE id=$3`
    res, err := db.Exec(query, rev.Rating, rev.Comment, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    count, _ := res.RowsAffected()
    if count == 0 {
        http.Error(w, "Review not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Review updated")
}

func deleteReview(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid review id", http.StatusBadRequest)
        return
    }

    res, err := db.Exec("DELETE FROM reviews WHERE id=$1", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    count, _ := res.RowsAffected()
    if count == 0 {
        http.Error(w, "Review not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Review deleted")
}

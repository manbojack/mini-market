package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"

    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}

var db *sql.DB

func main() {
    var err error
    dbURI := os.Getenv("DATABASE_URI")
    if dbURI == "" {
        log.Fatal("DATABASE_URI env var is required")
    }

    db, err = sql.Open("postgres", dbURI)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = createTable()
    if err != nil {
        log.Fatal(err)
    }

    r := mux.NewRouter()

    r.HandleFunc("/products", getProducts).Methods("GET")
    r.HandleFunc("/products/{id}", getProduct).Methods("GET")
    r.HandleFunc("/products", createProduct).Methods("POST")
    r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
    r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

    fmt.Println("Product service listening on :8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}

func createTable() error {
    query := `
    CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT,
        price NUMERIC(10,2) NOT NULL
    )`
    _, err := db.Exec(query)
    return err
}

func getProducts(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, description, price FROM products")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    products := []Product{}
    for rows.Next() {
        var p Product
        err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        products = append(products, p)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    var p Product
    err = db.QueryRow("SELECT id, name, description, price FROM products WHERE id=$1", id).
        Scan(&p.ID, &p.Name, &p.Description, &p.Price)
    if err == sql.ErrNoRows {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(p)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
    var p Product
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if p.Name == "" || p.Price <= 0 {
        http.Error(w, "Name and price are required", http.StatusBadRequest)
        return
    }

    err = db.QueryRow(
        "INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id",
        p.Name, p.Description, p.Price,
    ).Scan(&p.ID)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(p)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    var p Product
    err = json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if p.Name == "" || p.Price <= 0 {
        http.Error(w, "Name and price are required", http.StatusBadRequest)
        return
    }

    res, err := db.Exec(
        "UPDATE products SET name=$1, description=$2, price=$3 WHERE id=$4",
        p.Name, p.Description, p.Price, id,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    rowsAffected, _ := res.RowsAffected()
    if rowsAffected == 0 {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }

    p.ID = id
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(p)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    res, err := db.Exec("DELETE FROM products WHERE id=$1", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    rowsAffected, _ := res.RowsAffected()
    if rowsAffected == 0 {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

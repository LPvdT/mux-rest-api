package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name      string
	Price     float64
	Avalaible bool
}

func main() {
	const DB = "./test.db"

	// Connect
	connStr := DB

	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping db: %v", err)
	}

	// Create Table
	createTable(db)

	// Insert record
	product := Product{"Book2", 15.55, true}
	pk := insertProduct(db, product)

	log.Printf("Inserted data for ID: %d\n", pk)

	// Query single-row data
	var (
		name      string
		available bool
		price     float64
	)

	query := `
		SELECT
			name,
			available,
			price
		FROM
			product
		WHERE
			id = $1
	`

	err = db.QueryRow(query, pk).Scan(&name, &available, &price)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No rows found for ID: %d", pk)
		}
		log.Fatalf("Error querying: %v", err)
	}

	log.Printf("Data for ID: %d\n", pk)
	log.Printf("\tName: %s\n", name)
	log.Printf("\tAvailable: %t\n", available)
	log.Printf("\tPrice: %f\n", price)

	// Query multiple-row data
	data := []Product{}

	rows, err := db.Query(`SELECT name, available, price FROM product`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name, &available, &price)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Product{name, price, available})
	}

	log.Println(data)
}

func createTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS product (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			price NUMERIC(6, 2) NOT NULL,
			available BOOLEAN,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
}

func insertProduct(db *sql.DB, product Product) int {
	query := `
		INSERT INTO product (name, price, available)
			VALUES ($1, $2, $3) RETURNING id
	`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Avalaible).Scan(&pk)
	if err != nil {
		log.Fatalf("Error inserting to db: %v", err)
	}

	return pk
}

package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	"lpvdt/api/entity"

	_ "github.com/mattn/go-sqlite3"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

// Creates post table in sqlite DB
func init() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
		return
	}

	task := `
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			data JSON
		)
	`

	_, err = db.Exec(task)
	if err != nil {
		log.Fatalf("Error creating posts table: %v", err)
		return
	}

	defer db.Close()
}

// ...
func NewRepository() PostRepository {
	return &repo{}
}

// ...
func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
		return nil, err
	}
	defer db.Close()

	var id int
	query := `
		INSERT INTO posts (data)
			VALUES ('{
				"Title": "$1",
				"Text": "$2"
			}')
	`
	err = db.QueryRow(query, post.Title, post.Text).Scan(&id)
	if err != nil {
		log.Fatalf("Error inserting into db: %v", err)
		return nil, err
	}

	return post, nil
}

// ...
func (*repo) FindAll() ([]entity.Post, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
		return nil, err
	}
	defer db.Close()

	// Prepare container vars
	var (
		result string
		title  string
		text   string
		post   entity.Post
	)

	posts := []entity.Post{}

	// Send query
	query := `SELECT id, data FROM posts`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	// Process results
	for rows.Next() {
		err := rows.Scan(&post.Id, &result)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		// TODO: Serialize result string field to JSON
		encoded, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		// TODO: Unserialize and extract text and title - find out how
		// ...

		// Push to array
		posts = append(posts, post)
	}

	return posts, nil
}

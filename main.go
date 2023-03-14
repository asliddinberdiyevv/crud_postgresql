package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Book struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Author    string `json:"author,omitempty"`
	Publisher string `json:"publisher,omitempty"`
}

var books []Book

func main() {
	godotenv.Load(".env")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to database!")

	router := mux.NewRouter()

	router.HandleFunc("/", getBooks(db)).Methods("GET")

	// server
	var addr = ":" + os.Getenv("PORT")
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}
	// start the server
	log.Fatal(server.ListenAndServe())
}

func getBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Execute the SELECT query
		rows, err := db.Query("SELECT * FROM posts")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Process the results
		for rows.Next() {
			var book Book
			err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher)
			if err != nil {
				log.Fatal(err)
			}
			books = append(books, book)
		}

		// Write the response
		json.NewEncoder(w).Encode(books)
	}
}

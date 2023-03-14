package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	v1 "posts/pkgs/api/v1"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

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
	router.HandleFunc("/", v1.Version).Methods("GET")

	// server
	var addr = ":" + os.Getenv("PORT")
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}
	// start the server
	log.Fatal(server.ListenAndServe())
}

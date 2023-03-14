package main

import (
	"log"
	"net/http"
	"os"
	"posts/pkgs/api"
	"posts/pkgs/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Load env error: %v\n", err)
	}

	/* ---------- DATABASE ---------- */
	db, err := database.New()
	if err != nil {
		log.Printf("Error verifying database: %v\n", err)
		return
	} else {
		log.Println("Database is ready to use.")
	}

	/* ---------- ROUTER ---------- */
	router, err := api.NewRouter(db)
	if err != nil {
		log.Printf("Router: %v\n", err)
	}

	/* ---------- SERVER ---------- */
	var addr = ":" + os.Getenv("PORT")
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}
	log.Printf("Server run port: %s\n", os.Getenv("PORT"))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("Server: %v\n", err)
	}
}

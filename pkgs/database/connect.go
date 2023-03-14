package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/namsral/flag"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

var (
	databaseTimeout = flag.Int64("database-timeout-ms", 5000, "")
)

func Connect() (*sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("connect.go -> Connect(), Load env: %v", err)
	}

	conn, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, errors.Wrap(err, "Could not connect to database")
	}

	conn.SetMaxOpenConns(32)

	if err := waitForDB(conn.DB); err != nil {
		return nil, err
	}

	// Migrate database schema
	if err := migrateDB(conn.DB); err != nil {
		return nil, errors.Wrap(err, "could not migrate database.")
	} else {
		log.Println("Connected and migrated database")
	}

	return conn, nil
}

func New() (Database, error) {
	conn, err := Connect()
	if err != nil {
		return nil, err
	}

	d := &database{
		conn: conn,
	}

	return d, nil
}

func waitForDB(conn *sql.DB) error {
	ready := make(chan struct{})
	go func() {
		for {
			if err := conn.Ping(); err == nil {
				close(ready)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {
	case <-ready:
		return nil
	case <-time.After(time.Duration(*databaseTimeout) * time.Millisecond):
		return errors.New("Database not ready")
	}
}

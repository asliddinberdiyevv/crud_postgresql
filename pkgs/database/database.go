package database

import (
	"io"

	"github.com/jmoiron/sqlx"
)

const UniqueViolation = "unique_violation"

type Database interface {
	PostDB

	io.Closer
}

type database struct {
	conn *sqlx.DB
}

func (d *database) Close() error {
	return d.conn.Close()
}

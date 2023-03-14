package models

import (
	"errors"
	"time"
)

type Post struct {
	ID        string     `json:"id,omitempty" db:"post_id"`
	Name      *string    `json:"name,omitempty" db:"post_name"`
	Like      *bool      `json:"like,omitempty" db:"post_like"`
	Star      *bool      `json:"star,omitempty" db:"post_star"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

func (c *Post) Verify() error {
	if c.Name == nil || len(*c.Name) == 0 {
		return errors.New("name is requered")
	}
	return nil
}

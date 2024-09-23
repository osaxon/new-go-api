// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type Task struct {
	ID          int64          `json:"id"`
	Title       string         `json:"title"`
	UserID      int64          `json:"user_id"`
	DueDate     sql.NullTime   `json:"due_date"`
	Status      string         `json:"status"`
	Description sql.NullString `json:"description"`
}

type User struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Password string         `json:"-"`
	Bio      sql.NullString `json:"bio"`
}
// Filename: internal/data/entries.go
package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"quiz.3.driane.perez.net/internal/validator"
)

type Todo_list struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Task_Name   string    `json:"task_name"`
	Description string    `json:"desription"`
	Notes       string    `json:"notes"`
	Category    string    `json:"category"`
	Priority    string    `json:"priority"`
	Status      []string  `json:"status"`
	Version     int32     `json:"version"`
}
func ValidateEntires(v *validator.Validator, entries *Todo_list)  {
	//use the check method to execute our validation checks
	v.Check(entries.Task_Name != "", "task_name", "must be provided")
	v.Check(len(entries.Task_Name) <= 200, "task_name", "must not be more than 200 bytes long")

	v.Check(entries.Description != "", "description", "must be provided")
	v.Check(len(entries.Description) <= 200, "description", "must not be more than 200 bytes long")

	v.Check(entries.Notes != "", "notes", "must be provided")
	v.Check(len(entries.Notes) <= 200, "notes", "must not be more than 200 bytes long")

	v.Check(entries.Category != "", "category", "must be provided")
	v.Check(len(entries.Category) <= 200, "category", "must not be more than 200 bytes long")

	v.Check(entries.Priority != "", "priority", "must be provided")
	v.Check(len(entries.Priority) <= 200, "priority", "must not be more than 200 bytes long")

	v.Check(entries.Status != nil, "status", "must be provided")
	v.Check(len(entries.Status) >= 1, "status", "must contain one Status")
	v.Check(len(entries.Status) <= 5, "status", "must contain at least five Status")
	v.Check(validator.Unique(entries.Status),"status", "must not contain duplicate Status")
	
}
//define a todo_list model which wraps a sql.DB connection pool
type Todo_listModel struct {
	DB *sql.DB
}

//Insert() allows us to create a new todo_list
func (m Todo_listModel) Insert(Todo_list *Todo_list) error {
	query := `
	INSERT INTO todo_list (task_name, description, notes, category, priority, status)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at, version
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Collect the data fields into a slice
	args := []interface{}{Todo_list.Task_Name, Todo_list.Description, Todo_list.Notes,
		Todo_list.Category, Todo_list.Priority, pq.Array(Todo_list.Status),
	}
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&Todo_list.ID, &Todo_list.CreatedAt, &Todo_list.Version)
}
//GET () allow us to retrieve a specific todo_list
func (m Todo_listModel) Get(id int64) (*Todo_list, error) {
	return nil, nil
}
//Update() allows us to edit/alter a specific Todolist
func (m Todo_listModel) Update(Todo_list *Todo_list) error {
	return nil
}
//deletes() removes a specific todolist
func (m Todo_listModel) Delete(id int64) error {
	return nil
}
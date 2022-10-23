// Filename: internal/data/entries.go
package data

import (
	"context"
	"database/sql"
	"errors"
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

func ValidateEntires(v *validator.Validator, entries *Todo_list) {
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
	v.Check(validator.Unique(entries.Status), "status", "must not contain duplicate Status")

}

// define a todo_list model which wraps a sql.DB connection pool
type Todo_listModel struct {
	DB *sql.DB
}

// Insert() allows us to create a new todo_list
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
	args := []interface{}{
		Todo_list.Task_Name,
		Todo_list.Description,
		Todo_list.Notes,
		Todo_list.Category,
		Todo_list.Priority,
		pq.Array(Todo_list.Status),
	}
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&Todo_list.ID, &Todo_list.CreatedAt, &Todo_list.Version)
}

// GET () allow us to retrieve a specific todo_list
func (m Todo_listModel) Get(id int64) (*Todo_list, error) {
	//ensure that there is a valid id
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Create query
	query := `
		SELECT id, created_at, task_name, description, notes, category, priority, status, version
		FROM todo_list
		WHERE id = $1
	`
	// Declare a Todo_list variable to hold the return data
	var todo_list Todo_list

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Execute Query using the QueryRowContext()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&todo_list.ID,
		&todo_list.CreatedAt,
		&todo_list.Task_Name,
		&todo_list.Description,
		&todo_list.Notes,
		&todo_list.Category,
		&todo_list.Priority,
		pq.Array(&todo_list.Status),
		&todo_list.Version,
	)
	// Handle any errors
	if err != nil {
		// Check the type of error
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	// Success
	return &todo_list, nil
}

// Update() allows us to edit/alter a specific Todolist
func (m Todo_listModel) Update(Todo_list *Todo_list) error {
	query := `
	UPDATE todo_list 
	set task_name = $1,
	description = $2, 
	notes = $3,
	category = $4, 
	priority = $5,
	status = $6, 
	version = version + 1
	WHERE id = $7
	AND version = $8
	RETURNING version
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()

	args := []interface{}{
		Todo_list.Task_Name,
		Todo_list.Description,
		Todo_list.Notes,
		Todo_list.Category,
		Todo_list.Priority,
		pq.Array(Todo_list.Status),
		Todo_list.ID,
		Todo_list.Version,
	}
	// Check for edit conflicts
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&Todo_list.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

// deletes() removes a specific todolist
func (m Todo_listModel) Delete(id int64) error {
	// Ensure that there is a valid id
	if id < 1 {
		return ErrRecordNotFound
	}
	// Create the delete query
	query := `
		DELETE FROM todo_list
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Execute the query
	results, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	// Check how many rows were affected by the delete operations. We
	// call the RowsAffected() method on the result variable
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}
	// Check if no rows were affected
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

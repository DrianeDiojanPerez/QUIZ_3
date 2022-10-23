// Filename: internal/data/entries.go
package data

import (
	"time"

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
	
	// Name string `json:"name"`
	// String string `json:"string"`
	// Translate string `json:"translate"`
	// Phone string `json:"phone"`
	// Email string `json:"email"`
	// Website string `json:"website"`
	// Mode []string `json:"mode"`
	
	// Status string `json:"status,omitempty"`
	// Enviornment string `json:"enviornment,omitempty"`
	// Version string `json:"version,omitempty"`
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

	v.Check(entries.Version != 0, "priority", "must be provided")
	v.Check(entries.Version <= 200, "priority", "must not be more than 200 bytes long")
	
}
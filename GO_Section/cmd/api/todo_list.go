// Filename: cmd/api/entries.go
package main

import (
	"errors"
	"fmt"
	"net/http"

	"quiz.3.driane.perez.net/internal/data"
	"quiz.3.driane.perez.net/internal/validator"
)

// create entires hander for the POST /v1/entries endpoint
func (app *application) createtodo_listHandler(w http.ResponseWriter, r *http.Request) {
	//our target decode destination
	var todo_listDATA struct {
		Task_Name   string   `json:"task_name"`
		Description string   `json:"desription"`
		Notes       string   `json:"notes"`
		Category    string   `json:"category"`
		Priority    string   `json:"priority"`
		Status      []string `json:"status"`
	}
	err := app.readJSON(w, r, &todo_listDATA)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//copyung the values
	entries := &data.Todo_list{
		Task_Name:   todo_listDATA.Task_Name,
		Description: todo_listDATA.Description,
		Notes:       todo_listDATA.Notes,
		Category:    todo_listDATA.Category,
		Priority:    todo_listDATA.Priority,
		Status:      todo_listDATA.Status,
	}

	//initialize a new validator instance
	v := validator.New()

	//check the map to determine if there were any validation errors
	if data.ValidateEntires(v, entries); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	//create a todo_list
	err = app.models.Todo_list.Insert(entries)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	//creates a location header for newly created resource/todo_list
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/todo_list/%d", entries.ID))
	//write the JSON response with 201 - created status code with a the body
	//being the school data and the header being the headers map
	err = app.writeJSON(w, http.StatusCreated, envelope{"todo_list": entries}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// create showentires hander for the GET /v1/entries/:id endpoint
func (app *application) showtodo_listHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//fetch the specific todolist
	data_todolist, err := app.models.Todo_list.Get(id)
	//handle errors
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w,r)
		default:
			app.serverErrorResponse(w,r,err)
		}
		return
	}
	//write the data returned by Get()
	err = app.writeJSON(w, http.StatusOK, envelope{"todo_list": data_todolist}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

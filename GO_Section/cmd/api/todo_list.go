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
	var todo_listtodolistdata struct {
		Task_Name   string   `json:"task_name"`
		Description string   `json:"desription"`
		Notes       string   `json:"notes"`
		Category    string   `json:"category"`
		Priority    string   `json:"priority"`
		Status      []string `json:"status"`
	}
	err := app.readJSON(w, r, &todo_listtodolistdata)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//copyung the values
	entries := &data.Todo_list{
		Task_Name:   todo_listtodolistdata.Task_Name,
		Description: todo_listtodolistdata.Description,
		Notes:       todo_listtodolistdata.Notes,
		Category:    todo_listtodolistdata.Category,
		Priority:    todo_listtodolistdata.Priority,
		Status:      todo_listtodolistdata.Status,
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
	//being the school todolistdata and the header being the headers map
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
	todolistdata_todolist, err := app.models.Todo_list.Get(id)
	//handle errors
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	//write the todolistdata returned by Get()
	err = app.writeJSON(w, http.StatusOK, envelope{"todo_list": todolistdata_todolist}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) updateTodo_listHandler(w http.ResponseWriter, r *http.Request) {
	// This method does a partial replacement
	// Get the id for the todo_list item that needs updating
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Fetch the original record from the todolistdatabase
	todolist, err := app.models.Todo_list.Get(id)
	// hadles error
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Create an input struct to hold todolistdata read in from the client
	// We update the input struct to use pointers because pointers have a
	// default value of nil false
	// if a field remains nil then we know that the client did not update it
	//create an input struct to hold the todolistdata
	var todolistdata struct {
		Task_Name   *string  `json:"task_name"`
		Description *string  `json:"description"`
		Notes       *string  `json:"notes"`
		Category    *string  `json:"category"`
		Priority    *string  `json:"priority"`
		Status      []string `json:"status"`
	}

	//Initalize a new json.Decoder instance
	err = app.readJSON(w, r, &todolistdata)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Check for updates
	if todolistdata.Task_Name != nil {
		todolist.Task_Name = *todolistdata.Task_Name
	}
	if todolistdata.Description != nil {
		todolist.Description = *todolistdata.Description
	}
	if todolistdata.Notes != nil {
		todolist.Notes = *todolistdata.Notes
	}
	if todolistdata.Category != nil {
		todolist.Category = *todolistdata.Category
	}
	if todolistdata.Priority != nil {
		todolist.Priority = *todolistdata.Priority
	}
	if todolistdata.Status != nil {
		todolist.Status = todolistdata.Status
	}

	// Perform Validation on the updated Todo_list item. If validation fails then
	// we send a 422 - unprocessable entity response to the client
	// initialize a new Validator instance
	v := validator.New()

	//Check the map to determine if there were any validation errors
	if data.ValidateEntires(v, todolist); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Pass the update todo record to the Update() method
	err = app.models.Todo_list.Update(todolist)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusCreated, envelope{"todo_list": todolist}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
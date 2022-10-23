// Filename: cmd/api/entries.go
package main

import (
	"fmt"
	"net/http"

	"quiz.3.driane.perez.net/internal/data"
	"quiz.3.driane.perez.net/internal/validator"
)

//create entires hander for the POST /v1/entries endpoint
func (app *application) createInformationHandler(w http.ResponseWriter, r *http.Request){
	//our target decode destination
	var todo_list struct{
		Task_Name   string    `json:"task_name"`
		Description string    `json:"desription"`
		Notes       string    `json:"notes"`
		Category    string    `json:"category"`
		Priority    string    `json:"priority"`
		Status      []string  `json:"status"`	
	}
	err := app.readJSON(w, r, &todo_list)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//copyung the values
	entries := &data.Todo_list{
		Task_Name: todo_list.Task_Name,
		Description: todo_list.Description,
		Notes: todo_list.Notes,
		Category: todo_list.Category,
		Priority: todo_list.Priority,
		Status: todo_list.Status,
	}

	//initialize a new validator instance
	v := validator.New()
	
	//check the map to determine if there were any validation errors
	if data.ValidateEntires(v,entries); !v.Valid(){
		app.failedValidationResponse(w,r,v.Errors)
		return
	}
	//Display the request
	fmt.Fprintf(w, "%+v\n", todo_list)

	//fmt.Fprintln(w, "Create a New Entry")
}
//create showentires hander for the GET /v1/entries/:id endpoint
func (app *application) showRandomHandler(w http.ResponseWriter, r *http.Request){
	
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	int1:=int(id)
	tools := &data.Tools{
		Int: int1,
	}
	v := validator.New()
	if data.ValidateInt(v,tools); !v.Valid(){
		app.failedValidationResponse(w,r,v.Errors)
		return
	}
	strw:=tools.GenerateRandomString(int1)
	data := envelope{
		"id": int1,
		"random_string": strw,
		}
	err = app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}


	
}
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
	var input struct{
		Name string `json:"name"`
		String string `json:"string"`
		Translate string `json:"translate"`
		Phone string `json:"phone"`
		Email string `json:"email"`
		Website string `json:"website"`
		Mode []string `json:"mode"`
		
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//copyung the values
	entries := &data.EntriesData{
		Name: input.Name,
		String: input.String,
		Translate: input.Translate,
		Phone: input.Phone,
		Email: input.Email,
		Website: input.Website,
		Mode: input.Mode,
	}

	//initialize a new validator instance
	v := validator.New()
	
	//check the map to determine if there were any validation errors
	if data.ValidateEntires(v,entries); !v.Valid(){
		app.failedValidationResponse(w,r,v.Errors)
		return
	}
	//Display the request
	fmt.Fprintf(w, "%+v\n", input)

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
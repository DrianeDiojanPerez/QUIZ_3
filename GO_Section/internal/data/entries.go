// Filename: internal/data/entries.go
package data

import (
	"time"

	"quiz.3.driane.perez.net/internal/validator"
)

type EntriesData struct {
	ID int64 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name string `json:"name"`
	String string `json:"string"`
	Translate string `json:"translate"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Website string `json:"website"`
	Mode []string `json:"mode"`
	Version int32 `json:"version"`
	// Status string `json:"status,omitempty"`
	// Enviornment string `json:"enviornment,omitempty"`
	// Version string `json:"version,omitempty"`
}
func ValidateEntires(v *validator.Validator, entries *EntriesData)  {
	//use the check method to execute our validation checks
	v.Check(entries.Name != "", "name", "must be provided")
	v.Check(len(entries.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(entries.String != "", "string", "must be provided")
	v.Check(len(entries.String) <= 200, "string", "must not be more than 200 bytes long")

	v.Check(entries.Translate != "", "translate", "must be provided")
	v.Check(len(entries.Translate) <= 200, "translate", "must not be more than 200 bytes long")

	v.Check(entries.Phone != "", "phone", "must be provided")
	v.Check(validator.Matches(entries.Phone,validator.PhoneRegex), "phone", "must provide a valid phone number")
	v.Check(len(entries.Phone) <= 300, "phone", "must not be more than 300 bytes long")

	v.Check(entries.Email != "", "email", "must be provided")
	v.Check(len(entries.Email) <= 300, "email", "must not be more than 300 bytes long")
	v.Check(validator.Matches(entries.Email,validator.EmailRegex), "email", "must provide a valid email")

	v.Check(entries.Website != "", "website", "must be provided")
	v.Check(len(entries.Website) <= 200, "website", "must not be more than 200 bytes long")
	v.Check(validator.ValidWebsite(entries.Website), "website", "must be a valid URL")

	v.Check(entries.Mode != nil, "mode", "must be provided")
	v.Check(len(entries.Mode) >= 1, "mode", "must contain one mode")
	v.Check(len(entries.Mode) <= 5, "mode", "must contain at least five mode")
	v.Check(validator.Unique(entries.Mode),"mode", "must not contain duplicate entries")






			
		
		//Mode []string `json:"mode"`
	
}
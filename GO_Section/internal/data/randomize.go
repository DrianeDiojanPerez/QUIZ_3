// Filename: internal/data/randomize.go
package data

import (
	"crypto/rand"

	"quiz.3.driane.perez.net/internal/validator"
)

type Tools struct{
	Int int `json:"int"`
}

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"

func (t *Tools) GenerateRandomString(length int) string {
	s := make([]rune, length)
	r := []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		y := uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)

}
func ValidateInt(v *validator.Validator, Tool *Tools)  {
	//use the check method to execute our validation checks
	v.Check(Tool.Int <= 10000 , "int", "must provide a value less than or equal to 10,000")

}
package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	firstName   = "first_name"
	lastName    = "last_name"
	email       = "email"
	dateOfBirth = "date_of_birth"
)

type MinimalUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}

func EmailSignup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.ParseForm() != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}
	user := MinimalUser{}

	for k := range r.PostForm {
		switch k {
		case firstName:
			user.FirstName = ps.ByName(firstName)
		case lastName:
			user.LastName = ps.ByName(lastName)
		case email:
			user.Email = ps.ByName(email)
		case dateOfBirth:
			user.DateOfBirth = ps.ByName(dateOfBirth)

		}

	}
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.DateOfBirth == "" {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}
}

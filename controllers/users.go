package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

// Users controller meta-object
type Users struct {
	NewView *views.View
}

// SignupForm for use with gorilla/schema
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// NewUsers creates a users controller
//
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// New is used to render the new user signup form
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a user
// tries to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is ", form.Email)
	fmt.Fprintln(w, "Password is ", form.Password)
}

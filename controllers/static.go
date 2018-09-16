package controllers

import "lenslocked.com/views"

// NewStatic returns a Static struct for all semi-static views
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
		FAQ:     views.NewView("bootstrap", "views/static/faq.gohtml"),
	}
}

// Static struct to hold semi-static views
type Static struct {
	Home    *views.View
	Contact *views.View
	FAQ     *views.View
}

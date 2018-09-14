package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}
func contact(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "To get in touch, please send an email "+
    "to <a href=\"mailto:fake@address.com\">"+
    "fake@address.com</a>")
}
func faq(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Frequently Asked Questions:</h1></p> "+
    "1. Yes? No</p>")
}
func notfound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprint(w, "<h1>Unable to find the page you requested.</h1>")
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", home)
  r.HandleFunc("/contact", contact)
  r.HandleFunc("/faq", faq)
  var h http.Handler = http.HandlerFunc(notfound)
  r.NotFoundHandler = h
  http.ListenAndServe(":3000", r)
}

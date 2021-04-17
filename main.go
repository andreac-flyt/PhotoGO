package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("/", "home")
	fmt.Fprint(w, "<h1>Welcome to my Awesome Page!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("/contact", "text/http")
	fmt.Fprint(w, "<h1>To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a></h1>.")

}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("/faq", "text/http")
	fmt.Fprint(w, "<h1>Frequently Asked Questions...</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/http")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>The page you are looking for does not exist.</h1>")
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}

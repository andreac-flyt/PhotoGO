package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/andreac-flyt/PhotoGO/internal/views"

	"github.com/gorilla/mux"
)

var (
		contactTemplate *template.Template
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("/", "home")
	homeTemplate := views.NewView("internal/views/home.gohtml")
	if err := homeTemplate.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("/faq", "text/http")
	fmt.Fprint(w, "<h1>Frequently Asked Questions...</h1>")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}

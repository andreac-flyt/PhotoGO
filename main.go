package main

import (
	"fmt"
	"net/http"

	"github.com/andreac-flyt/PhotoGO/internal/views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
<<<<<<< HEAD
	err := homeView.Template.Execute(w, nil)
=======
	err := homeView.Template.Exexcute(w, nil)
>>>>>>> de94e36d3a4db04bfcb8bbbc2f9515306644483c
	if err != nil {
		panic(err)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("/faq", "text/http")
	fmt.Fprint(w, "<h1>Frequently Asked Questions...</h1>")
}

func main() {
<<<<<<< HEAD
	homeView = views.NewView("internal/views/home.gohtml")
	contactView = views.NewView("internal/views/contact.gohtml")
=======
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
>>>>>>> de94e36d3a4db04bfcb8bbbc2f9515306644483c

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}

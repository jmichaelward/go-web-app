package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type LayoutData struct {
	PageTitle string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html")

	if err != nil {
		fmt.Fprintf(w, "Something has gone horribly wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, LayoutData{
		PageTitle: "J. Michael Ward, Software Engineer",
	})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html")

	if err != nil {
		fmt.Fprintf(w, "Something has gone horribly wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, LayoutData{
		PageTitle: "J. Michael Ward, Software Engineer – About",
	})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/about", AboutHandler)

	http.Handle("/", router)

	http.ListenAndServe(":8000", router)
}

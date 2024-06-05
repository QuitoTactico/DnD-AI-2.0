package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Print("Server Running\n")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/home.html"))
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8000", nil)

	fmt.Print("Server Close")
}

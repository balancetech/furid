package main

import (
	"log"
	"net/http"
	"html/template"
)

var rootTemplate = template.Must(template.ParseFiles("templates/index.html"))

func root(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

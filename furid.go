package main

import (
	"log"
	"net/http"
	"html/template"
	"database/sql"

	_ "github.com/lib/pq"
)

var rootTemplate = template.Must(template.ParseFiles("templates/index.html"))
var createAccountTemplate = template.Must(template.ParseFiles("templates/createaccount.html"))

func root(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, nil)
}

func createaccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		createAccountTemplate.Execute(w, nil)
	} else if r.Method == "POST" {
		n := r.FormValue("username")
		log.Print(n)
		w.Write([]byte("Go fuck yourself"))
		db, err := sql.Open("postgres", "user=wesley dbname=balance sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		if _, err := db.Exec("INSERT INTO users (username) VALUES ($1)", n); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/createaccount", createaccount)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

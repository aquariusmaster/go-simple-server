package server

import (
	"fmt"
	"net/http"
	"html/template"
)

func InitServer() {
	fmt.Println("started")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err!= nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

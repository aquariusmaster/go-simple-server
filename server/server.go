package server

import (
	"fmt"
	"net/http"
)

func InitServer() {
	fmt.Println("started")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world!</h1>")
}

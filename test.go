package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("view/index.html")
		t.Execute(w, nil)
	}
}

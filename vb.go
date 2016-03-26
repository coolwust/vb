package vb

import (
	"net/http"
	"html/template"
)

func init() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("view/index.html")).Execute(w, nil)
}

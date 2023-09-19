package main

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed post-method.html
var postMethodHtml string

func indexHandler(w http.ResponseWriter, req *http.Request) {
	tmpl := template.New("")
	t, err := tmpl.Parse(postMethodHtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

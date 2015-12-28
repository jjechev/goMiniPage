package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}


func handler(w http.ResponseWriter, r *http.Request) {

	url := r.URL
	tpl := url.Path[1:]

	if tpl == "" {
		tpl = "homePage"
	}

	file := "template/" + tpl + ".html"
	t, err := template.ParseFiles(file)
	if err == nil {
		t.Execute(w, "")
	} else {
		fmt.Fprintf(w, "404 not found")
	}
}

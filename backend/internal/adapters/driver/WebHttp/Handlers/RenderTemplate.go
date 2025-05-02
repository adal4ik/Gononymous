package handlers

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) *Page {
	filename := "web/templates/" + title
	body, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &Page{Title: title, Body: body}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	p := LoadPage(tmpl)
	t, err := template.ParseFiles("web/templates/" + tmpl)
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Execute(w, p)
}

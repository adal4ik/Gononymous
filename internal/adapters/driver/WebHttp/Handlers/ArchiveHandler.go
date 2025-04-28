package handlers

import (
	"html/template"
	"net/http"
)

func ArchiveHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/archive.html"))
	tmpl.Execute(w, nil)
}

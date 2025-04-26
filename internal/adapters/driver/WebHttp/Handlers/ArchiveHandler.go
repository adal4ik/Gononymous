package handlers

import "net/http"

func ArchiveHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "archive.html")
}

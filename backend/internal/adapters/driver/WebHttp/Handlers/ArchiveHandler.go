package handlers

import (
	"bytes"
	"html/template"
	"net/http"

	driverports "backend/internal/core/ports/driver_ports"
)

type ArchiveHandler struct {
	service driverports.PostDriverPortInterface
	BaseHandler
}

func NewArchiveHandler(service driverports.PostDriverPortInterface, baseHandler BaseHandler) *ArchiveHandler {
	return &ArchiveHandler{
		service:     service,
		BaseHandler: baseHandler,
	}
}

func (h *ArchiveHandler) GetArchivePage(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetAll()
	if err != nil {
		h.handleError(w, r, http.StatusInternalServerError, "Failed to get posts", err)
		return
	}

	tmpl, err := template.ParseFiles("templates/archive.html")
	if err != nil {
		h.handleError(w, r, http.StatusInternalServerError, "Failed to parse template", err)
		return
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, posts)
	if err != nil {
		h.handleError(w, r, http.StatusInternalServerError, "Failed to execute template", err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(buf.Bytes())
	if err != nil {
		h.handleError(w, r, http.StatusInternalServerError, "Failed to write response", err)
		return
	}
	h.logger.Info("Archive page rendered successfully", "url", r.URL.Path)
}

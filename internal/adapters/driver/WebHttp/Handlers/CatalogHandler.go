package handlers

import (
	"net/http"

	driverports "Gononymous/internal/core/ports/driver_ports"
)

type CatalogHandler struct {
	service driverports.PostDriverPortInterface
}

func NewCatalogHandler() *CatalogHandler {
	return &CatalogHandler{}
}

func (c *CatalogHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "catalog.html")
}

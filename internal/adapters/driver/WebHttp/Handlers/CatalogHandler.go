package handlers

import (
	"Gononymous/internal/adapters/driver/WebHttp/middleware"
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
	middleware.CreateCookie(w, r)
	renderTemplate(w, "catalog.html")
}

package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"backend/internal/core/domains/dto"
	driverports "backend/internal/core/ports/driver_ports"
)

type CatalogHandler struct {
	service driverports.PostDriverPortInterface
	BaseHandler
}

func NewCatalogHandler(service driverports.PostDriverPortInterface, baseHandler BaseHandler) *CatalogHandler {
	return &CatalogHandler{service: service, BaseHandler: baseHandler}
}

func (c *CatalogHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	posts, err := c.service.GetActive()
	if err != nil {
		c.handleError(w, r, 500, "failed to get", err)
		c.RenderError(w, 500, "asd")
		return

	}
	tmpl := template.Must(template.ParseFiles("web/templates/catalog.html"))
	HTMLXposts := map[string][]dto.PostDto{
		"Posts": posts,
	}
	fmt.Println(HTMLXposts)
	err = tmpl.Execute(w, HTMLXposts)
	if err != nil {
		c.handleError(w, r, 500, "failed to get", err)
		c.RenderError(w, 500, "asd")
		return
	}
}

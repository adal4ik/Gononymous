package handlers

import (
	"Gononymous/internal/core/domains/dto"
	"fmt"
	"html/template"
	"log"
	"net/http"

	driverports "Gononymous/internal/core/ports/driver_ports"
)

type CatalogHandler struct {
	service driverports.PostDriverPortInterface
}

func NewCatalogHandler(service driverports.PostDriverPortInterface) *CatalogHandler {
	return &CatalogHandler{service: service}
}

func (c *CatalogHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	posts, err := c.service.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/catalog.html"))
	HTMLXposts := map[string][]dto.PostDto{
		"Posts": posts,
	}
	fmt.Println(HTMLXposts)
	err = tmpl.Execute(w, HTMLXposts)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	renderTemplate(w, "catalog.html")
}

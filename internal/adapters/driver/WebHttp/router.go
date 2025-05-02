package WebHttp

import (
	"Gononymous/internal/adapters/driver/WebHttp/middleware"
	driverports "Gononymous/internal/core/ports/driver_ports"
	"net/http"

	handlers "Gononymous/internal/adapters/driver/WebHttp/Handlers"
)

func Router(handlers *handlers.Handler, sessionservice driverports.SessionServiceDriverInterface) http.Handler {
	mux := http.NewServeMux()
	// Initializing Handlers

	// CATALOG RELATED STUFF
	mux.HandleFunc("/", handlers.CatalogHandler.MainPage)

	// POST RELATED STAFF
	mux.HandleFunc("/create-post", handlers.PostHandler.MainPage)
	mux.HandleFunc("POST /submit-post", handlers.PostHandler.SubmitPostHandler)

	var handler http.Handler = mux
	handler = middleware.SessionHandler(mux, sessionservice)
	// ARHIEVE RELATED STAFF
	// mux.HandleFunc("/archive", handlers.ArchiveHandler)
	return handler
}

package WebHttp

import (
	"net/http"

	handlers "backend/internal/adapters/driver/WebHttp/Handlers"
	"backend/internal/adapters/driver/WebHttp/middleware"
	driverports "backend/internal/core/ports/driver_ports"
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
	mux.HandleFunc("/archive", handlers.ArchiveHandler.GetArchivePage)
	return handler
}

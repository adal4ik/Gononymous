package WebHttp

import (
	"net/http"

	"backend/internal/adapters/driver/WebHttp/middleware"

	handlers "backend/internal/adapters/driver/WebHttp/Handlers"
)

func Rounter(handlers *handlers.Handler) http.Handler {
	mux := http.NewServeMux()
	// Initializing Handlers

	// CATALOG RELATED STUFF
	mux.HandleFunc("/", handlers.CatalogHandler.MainPage)

	// POST RELATED STAFF
	mux.HandleFunc("/create-post", handlers.PostHandler.MainPage)
	mux.HandleFunc("POST /submit-post", handlers.PostHandler.SubmitPostHandler)

	var handler http.Handler = mux
	handler = middleware.SessionHandler(mux)
	// ARHIEVE RELATED STAFF
	// mux.HandleFunc("/archive", handlers.ArchiveHandler)
	return handler
}

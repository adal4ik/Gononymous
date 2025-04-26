package WebHttp

import (
	"net/http"

	handlers "Gononymous/internal/adapters/driver/WebHttp/Handlers"
)

func Rounter(handlers *handlers.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	// Initializing Handlers

	// CATALOG RELATED STUFF
	mux.HandleFunc("/", handlers.CatalogHandler.MainPage)

	// POST RELATED STAFF
	mux.HandleFunc("/create-post", handlers.PostHandler.MainPage)
	mux.HandleFunc("POST /submit-post", handlers.PostHandler.SubmitPostHandler)

	// ARHIEVE RELATED STAFF
	// mux.HandleFunc("/archive", handlers.ArchiveHandler)
	return mux
}

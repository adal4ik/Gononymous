package WebHttp

import (
	"backend/internal/adapters/driver/WebHttp/middleware"
	"net/http"

	handlers "backend/internal/adapters/driver/WebHttp/Handlers"

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
	mux.HandleFunc("/post/{id}", handlers.PostHandler.PostPage)

	// COMMENT RELATED STUFF
	mux.HandleFunc("POST /submit-comment", handlers.CommentHandler.SubmitComment)

	var handler http.Handler = mux
	handler = middleware.SessionHandler(mux, sessionservice)
	// ARHIEVE RELATED STAFF
	mux.HandleFunc("/archive", handlers.ArchiveHandler.GetArchivePage)
	mux.HandleFunc("/archive-post/{id}", handlers.ArchiveHandler.GetArchivePost)
	return handler
}

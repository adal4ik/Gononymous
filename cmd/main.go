package main

import (
	"fmt"
	"net/http"

	db "Gononymous/internal/adapters/driven/database"
	"Gononymous/internal/adapters/driver/WebHttp"
	handlers "Gononymous/internal/adapters/driver/WebHttp/Handlers"
	"Gononymous/internal/core/services"
)

var port = ":8080"

func main() {
	database := db.ConnectDB()
	defer database.Close()

	repositories := db.New(database)
	services := services.New(repositories)
	handlers := handlers.New(services)

	mux := WebHttp.Rounter(handlers)
	fmt.Println("Server is running on port: http//localhost" + port)
	http.ListenAndServe(port, mux)
}

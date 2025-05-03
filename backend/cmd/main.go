package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	db "backend/internal/adapters/driven/database"
	"backend/internal/adapters/driver/WebHttp"
	handlers "backend/internal/adapters/driver/WebHttp/Handlers"
	"backend/internal/core/services"
	"backend/utils"
)

var port = ":8080"

func main() {
	ctx := context.Background()

	database := db.ConnectDB()
	defer database.Close()
	logger, logFile := utils.Logger()
	defer logFile.Close()
	baseHandler := handlers.NewBaseHandler(*logger)
	repositories := db.New(database)
	services := services.New(repositories)
	handlers := handlers.New(services, *baseHandler)

	mux := WebHttp.Router(handlers, services.SessionService)

	httpServer := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	go func() {
		fmt.Println("Server is running on port: http//localhost" + port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
}

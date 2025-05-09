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
	"backend/internal/adapters/driver/cli"
	"backend/internal/core/services"
	"backend/utils"
)

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

	// ⬇️ Создаем context для фоновых задач
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// ⬇️ Запускаем фоновую задачу архивирования постов
	services.PostsService.StartPostArchiver(ctx, time.Minute)

	mux := WebHttp.Router(handlers, services.SessionService)

	httpServer := &http.Server{
		Addr:    cli.Port,
		Handler: mux,
	}

	// Запускаем HTTP сервер в отдельной горутине
	go func() {
		fmt.Println("Server is running on port: http://localhost" + cli.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	// Ждем завершения (Ctrl+C)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
}

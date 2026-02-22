package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"x-service/internal/core/config"
	"x-service/internal/core/pkg"
	"x-service/internal/users/adapters/postgres"
	"x-service/internal/users/handlers"
	"x-service/internal/users/usecases"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file found: %v", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := pkg.NewPostgresDB(ctx, cfg.Postgres)
	if err != nil {
		log.Fatalf("init database: %v", err)
	}

	defer func() {
		if err := pkg.Close(db); err != nil {
			log.Printf("close db: %v", err)
		}
	}()

	repository := postgres.NewStorage(db)
	service := usecases.NewUserService(repository)
	httpHandlers := handlers.NewUserHandlers(service)
	server := handlers.NewHTTPServer(cfg.Server.Port, httpHandlers)

	go func() {
		log.Printf("server starting on port %s", cfg.Server.Port)
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutdown signal received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("server stopped gracefully")
}

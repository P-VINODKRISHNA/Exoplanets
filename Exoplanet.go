package main

import (
	"Exoplanet/config"
	"Exoplanet/routers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Println("Error initializing configuration:", err)
		return
	}

	r := routers.Init()

	// Start server
	go func() {
		if err := r.Run(":" + viper.GetString("HTTP_PORT")); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	srv := &http.Server{
		Addr:        ":" + viper.GetString("HTTP_PORT"),
		Handler:     r,
		ReadTimeout: 5 * time.Second, // Set the ReadHeaderTimeout here
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}

package main

import (
	"context"
	"go-auth/internal/app"
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()

	a, err := app.New()
	if err != nil {
		log.Fatalf("Start up failed: %v", err)
	}
	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Printf("Shutdown warning: %v", err)
		}
	}()

	router := httpserver.NewRouter(a)

	srv := &http.Server{
		Addr:              ":5000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Api is running on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server closed")
			return
		}
		log.Fatalf("Server error: %v", err)
	}

}

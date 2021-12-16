package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func serveHTTP() {
	ctx := context.Background()

	router := mux.NewRouter().StrictSlash(false)

	srv := &http.Server{
		Addr:    port(),
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("starting server: %s", err)
			os.Exit(3)
		}
	}()
	log.Printf("starting server: %s", port())
	<-done

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("stopping server: %s", err)
		os.Exit(4)
	}
}

func port() string {
	port := "8080"
	return fmt.Sprintf(":%s", port)
}

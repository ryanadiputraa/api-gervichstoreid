package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	dbUtil "github.com/ryanadiputraa/api-gervichstore.id/pkg/database"
	"github.com/spf13/viper"

	_productHandler "github.com/ryanadiputraa/api-gervichstore.id/app/product/delivery/http"
	_productRepository "github.com/ryanadiputraa/api-gervichstore.id/app/product/repository/psql"
	_productUsecase "github.com/ryanadiputraa/api-gervichstore.id/app/product/usecase"
)

func serveHTTP() {
	ctx := context.Background()
	readConnection, writeConnection := initDatabase(ctx)

	sessionRead := readConnection.NewSession(nil)
	sessionRead.Timeout = 10 * time.Second
	sessionWrite := writeConnection.NewSession(nil)
	sessionWrite.Timeout = 10 * time.Second

	router := mux.NewRouter().StrictSlash(false)
	router.Use(Recovery)

	v1 := router.PathPrefix("/v1").Subrouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		Debug:            true,
		AllowCredentials: true,
	})
	c.Handler(CorsMiddleware())
	handler := c.Handler(router)

	productRepository := _productRepository.NewProductRepository(sessionRead, sessionWrite)
	productUsecase := _productUsecase.NewProductUseCase(sessionRead, sessionWrite, productRepository)
	_productHandler.NewProductHandler(v1, productUsecase)

	srv := &http.Server{
		Addr:    port(),
		Handler: handler,
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
	port := viper.GetString("PORT")
	if len(port) == 0 {
		port = ":8080"
	}
	return fmt.Sprintf(":%s", port)
}

func initDatabase(ctx context.Context) (read, write *dbr.Connection) {
	read, err := dbUtil.CreateConnection(
		viper.GetString("database.driver"),
		viper.GetString("database.read"),
		viper.GetInt("database.max_conns"),
		viper.GetInt("database.max_idle"),
	)
	if err != nil {
		log.Printf("CREATE database connection READ: %s", err.Error())
		os.Exit(3)
	}

	err = read.PingContext(ctx)
	if err != nil {
		log.Printf("PING database connection READ: %s", err.Error())
		os.Exit(2)
	}

	write, err = dbUtil.CreateConnection(
		viper.GetString("database.driver"),
		viper.GetString("database.read"),
		viper.GetInt("database.max_conns"),
		viper.GetInt("database.max_idle"),
	)
	if err != nil {
		log.Printf("CREATE database connection WRITE: %s", err.Error())
		os.Exit(1)
	}

	err = write.PingContext(ctx)
	if err != nil {
		log.Printf("PING database connection WRITE: %s", err.Error())
		os.Exit(2)
	}

	return
}

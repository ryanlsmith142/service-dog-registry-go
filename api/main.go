package main

import (
	"../models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)
const version = "1.0.0"
func main() {

	newApi := createNewApi()

	srv := &http.Server {
		Addr: fmt.Sprintf(":%d", newApi.config.port),
		Handler: newApi.newRouter(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	db, err := openDB(newApi.config)
	if err != nil {
		newApi.logger.Fatal(err)
	}

	newApi = &api {
		models: models.NewModels(db),
	}

	defer db.Close()

	newApi.logger.Println("Starting server on port", newApi.config.port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}



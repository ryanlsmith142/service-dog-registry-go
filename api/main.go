package main

import (
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

	newApi.logger.Println("Starting server on port", newApi.config.port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}



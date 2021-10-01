package main

import (
	"flag"
	"log"
	"os"

)

type config struct {
	port int
	env  string
	db struct {
		dsn string
	}
}

type ApiStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type api struct {
	config config
	logger *log.Logger
}

func createNewConfig() config {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application Environment (development | production")
	flag.StringVar(&cfg.db.dsn, "dsn", localDSN, "postgres connection string")
	flag.Parse()
	return cfg
}

func createNewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return logger
}

func createNewApi() api {
	a := api {
		config: createNewConfig(),
		logger: createNewLogger(),
	}

	return a
}

package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/johhess40/backendapp/models"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type Config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config Config
	logger *log.Logger
	models models.Models
}

func main() {
	var cfg Config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on ")
	flag.StringVar(&cfg.env, "env", "dev", "Environment to build in")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgresql://localhost:5432/mydb?sslmode=disable", "Environment to build in")
	flag.Parse()

	logger := log.New(os.Stdout, fmt.Sprintf("[%s]", cfg.env), log.Ldate|log.Ltime)

	db, err := openDb(&cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: *models.NewModels(db),
	}

	log.Printf("Running...")
	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: time.Minute,
	}

	logger.Println("Starting server on port", cfg.port)

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func openDb(cfg *Config) (*sql.DB, error) {
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

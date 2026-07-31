package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// api version
const version = "1.0.0"

// application configuration

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}
type application struct {
	cfg    config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment development|staging|production")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "Maximum number of open connections to the database")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 5, "Maximum number of idle connections to the database")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")
	// environement variable for PostgreSQL DSN
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	DSN := os.Getenv("DB_DSN")

	if DSN == "" {
		log.Fatalf("DSN was not set yet")
		os.Exit(1)
	}
	// Connection to PostgreSQL DSN
	flag.StringVar(&cfg.db.dsn, "db-dsn", DSN, "PostgreSQL DSN")
	flag.Parse()

	// initializing logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error("error opening database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		cfg:    cfg,
		logger: logger,
	}

	logger.Info("Connection to PostgreSQL established")
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server on port:", "addr", srv.Addr, "env", cfg.env)
	err = srv.ListenAndServe()

	logger.Error(err.Error())

	os.Exit(1)

}

// Function for connection to PostgreSQL
func openDB(cfg config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close() // close connection with Database
		return nil, err
	}

	return db, nil
}

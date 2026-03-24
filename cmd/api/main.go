package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// config - представляет конфигурации приложения
type config struct {
	port int // Порт подключения
	env  string
}

// application - представляет из собой структуру приложения
type application struct {
	config config
	logger *slog.Logger
}

func main() {

	var cfg config

	// Флаги
	flag.IntVar(&cfg.port, "port", 4000, "API server port")                                        // получение порта
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)") // окружение
	flag.Parse()

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: log,
	}

	mux := http.NewServeMux() // Создание обрабочик
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)

	// Создание сервера
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port), // порт подключения
		Handler:      mux,                          // обработчики
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(log.Handler(), slog.LevelError),
	}

	log.Info("Starting API server on address: ", server.Addr)
	err := server.ListenAndServe() //

	log.Error(err.Error())
	os.Exit(1)

}

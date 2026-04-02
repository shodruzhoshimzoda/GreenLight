package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes устанавливает маршруты для приложения.
func (app *application) routes() http.Handler {

	router := httprouter.New() // Создание обработчика на основе библеотеки httprouter
	// Настройка использования наших кастомных хелперов для обработки ошибок.
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheck)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}

package main

import (
	"fmt"
	"net/http"
)

// logger - метод для логгирования ошибок
func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		url    = r.URL.String()
	)

	app.logger.Error(err.Error(), "method", method, "url", url)
}

// errorResponse - вспомогательный метод для отправки ответов об ошибках
func (app *application) errorResponses(w http.ResponseWriter, r *http.Request, status int, message any) {

	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)

	}

}

// serveErrorResponse используется при возникновении неожиданных проблем при выполнении.
// Он логирует детальную ошибку и отправляет клиенту статус 500.
func (app *application) serveErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request."
	app.errorResponses(w, r, http.StatusInternalServerError, message)

}

// NotFoundResponse отправляет клиенту статус 404 (Not Found) в формате JSON.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponses(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse отправляет статус 405, если вызванный HTTP-метод
// не поддерживается данным эндпоинтом.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the requested method %s is not allowed", r.Method)
	app.errorResponses(w, r, http.StatusMethodNotAllowed, message)
}

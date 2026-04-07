package main

import (
	"fmt"
	"net/http"
)

// recoverPanic — это middleware, которое восстанавливает работу приложения после возникновения паники.
// В случае паники оно закрывает соединение и отправляет клиенту HTTP-ответ 500 Internal Server Error.
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")

				app.serveErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

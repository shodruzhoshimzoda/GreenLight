package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display create movie handler")
}
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParams(r)
	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintln(w, "Display show movie handler: ", id)
}

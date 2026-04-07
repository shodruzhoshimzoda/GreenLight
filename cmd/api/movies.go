package main

import (
	"Greenlight/internal/data"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		app.errorResponses(w, r, http.StatusBadRequest, err.Error())
		return

	}

	fmt.Fprintf(w, "%v", input)

}
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Пример создания структуры для отправки для
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serveErrorResponse(w, r, err)
	}

}

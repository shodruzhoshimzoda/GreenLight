package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// readIDParams - данная вспомогательная функция используется для извлечения ID из тело URL и возвршает его
func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

type envelope map[string]any

// writeJson - вспомогательная функция которая помогает для переобразование объекта в JSON
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, header http.Header) error {

	js, err := json.MarshalIndent(data, " ", "\t") // Для более удобочитаемости используется MarshalIndent()
	if err != nil {
		return err
	}

	for k, v := range header {
		w.Header()[k] = v
	}

	js = append(js, '\n') // Добавления в конце спика байтов символ перехода на новую строку
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	w.Write(js)

	return nil
}

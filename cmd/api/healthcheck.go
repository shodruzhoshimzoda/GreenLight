package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "status: available")
	//fmt.Fprintln(w, "environment: ", app.config.env)
	//fmt.Fprintln(w, "version: "+version)

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// json.Marshal function returns encoded json
	js, err := json.Marshal(data)

	if err != nil {

		// Calling our logger
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your reque", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set header
	w.Write(js)                                        // write data in object body

}

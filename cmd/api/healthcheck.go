package main

import (
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
	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "There was an error processing the request: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

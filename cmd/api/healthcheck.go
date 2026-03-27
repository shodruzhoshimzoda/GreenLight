package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "status: available")
	//fmt.Fprintln(w, "environment: ", app.config.env)
	//fmt.Fprintln(w, "version: "+version)

	json := `{"status": "available", "environment":%q, "version":%q}`

	jsonTag := fmt.Sprintf(json, app.config.env, version)

	// Change header
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(jsonTag))

}

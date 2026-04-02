package main

import (
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "status: available")
	//fmt.Fprintln(w, "environment: ", app.config.env)
	//fmt.Fprintln(w, "version: "+version)

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serveErrorResponse(w, r, err)
	}

}

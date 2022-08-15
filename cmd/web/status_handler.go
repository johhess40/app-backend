package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentstatus := AppStatus{
		Status:      "available",
		Environment: app.config.env,
		Version:     version,
	}
	js, err := json.MarshalIndent(currentstatus, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (app *application) getStateFile(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errJson(w, err)
		return
	}
	app.logger.Println("Id is ", id)

	resource, err := app.models.DB.GetResource(id)

	err = app.writeJson(w, http.StatusOK, resource, "state")
}

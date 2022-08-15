package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/state/:id", app.getStateFile)
	router.HandlerFunc(http.MethodGet, "/v1/resources", app.getAllResources)
	return router
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) writeJson(w http.ResponseWriter, statusCode int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper["wrapper"] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
	return nil
}

func (app *application) errJson(w http.ResponseWriter, err error) {
	type jsonErr struct {
		Message string `json:"message"`
	}

	theError := jsonErr{Message: fmt.Sprintf("Error getting JSON with bad id: %s", err.Error())}

	app.writeJson(w, http.StatusBadRequest, theError, "error")
}

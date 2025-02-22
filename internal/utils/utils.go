package utils

import (
	"encoding/json"
	"net/http"
)

func HandleResponse(w http.ResponseWriter, resp interface{}) {
	jsonOut, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonOut)
}

func HandleErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	jsonOut, err := json.Marshal(err.Error())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(jsonOut)

}

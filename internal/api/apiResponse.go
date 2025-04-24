package api

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

type APIResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func JsonResponse(w http.ResponseWriter, data any, err error, successCode int) {
	w.Header().Set("Content-Type", "application/json")

	response := APIResponse{
		Data:    data,
		Code:    successCode,
		Message: "success",
	}

	if err != nil {
		response.Message = err.Error()
	}

	switch {
	case reflect.ValueOf(data).IsNil():
		response.Message = "no data found"
		response.Code = http.StatusNotFound

	case err != nil:
		response.Message = err.Error()
		// Don't override status code if it was already set to something specific
		if response.Code == http.StatusOK {
			log.Printf("developer setting status code ok when error present")
			response.Code = http.StatusInternalServerError
		}

	}

	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func JsonResponseBadRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	response := APIResponse{
		Data:    nil,
		Code:    http.StatusBadRequest,
		Message: "request does not match schema",
	}

	if err != nil {
		response.Message = err.Error()
	}

	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

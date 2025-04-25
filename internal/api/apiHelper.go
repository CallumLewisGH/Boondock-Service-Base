package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func EncodeJsonResponse(w http.ResponseWriter, data any, err error, responseCode int) {
	w.Header().Set("Content-Type", "application/json")

	response := APIResponse{
		Data:    data,
		Code:    responseCode,
		Message: "success",
	}

	if err != nil {
		response.Message = err.Error()
	}

	if responseCode == http.StatusOK && err != nil {
		log.Printf("developer setting status code ok when error present")
		response.Code = http.StatusInternalServerError
	}

	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func EncodeJsonResponseBadRequest(w http.ResponseWriter, err error) {
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

func DecodeJsonRequest[T any](r *http.Request, w http.ResponseWriter, model T) *T {
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		EncodeJsonResponseBadRequest(w, err)
		return nil
	}
	defer r.Body.Close()

	return &model
}

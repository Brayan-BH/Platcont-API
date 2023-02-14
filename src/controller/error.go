package controller

import (
	"encoding/json"
	"net/http"
)

func ErrorUser(w http.ResponseWriter, err error, statusCode int) {
	response := NewResponseManager()
	response.Msg = err.Error()
	response.Status = "Error de Usuario"
	response.StatusCode = statusCode
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func ErrorServer(w http.ResponseWriter, err error) {
	response := NewResponseManager()
	response.Msg = err.Error()
	response.Status = "Error de Servidor"
	response.StatusCode = 500
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func ErrorsError(w http.ResponseWriter, err error) {
	response := NewResponseManager()
	response.Msg = err.Error()
	response.Status = "Error"
	response.StatusCode = 400
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func ErrorsWaning(w http.ResponseWriter, err error) {
	response := NewResponseManager()
	response.Msg = err.Error()
	response.Status = "Warning"
	response.StatusCode = 300
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func ErrorsSuccess(w http.ResponseWriter, err error) {
	response := NewResponseManager()
	response.Msg = err.Error()
	response.Status = "Success"
	response.StatusCode = 200
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func ErrorsInfo(w http.ResponseWriter, err error) {
	response := NewResponseManager()
	response.Msg = err.Error()
	response.Status = "Info"
	response.StatusCode = 201
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

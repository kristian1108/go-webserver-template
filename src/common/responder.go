package common

import (
	"encoding/json"
	"net/http"
)

type JSONError struct {
	Data interface{} `json:"data"`
	Err  string      `json:"error"`
}

type coder interface {
	Code() int
}

func ErrorResponse(rw http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if e, ok := err.(coder); ok && e.Code() != 0 {
		status = e.Code()
	}

	jsonError := JSONError{Err: err.Error()}
	body, err := json.Marshal(jsonError)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	_, _ = rw.Write(body)
}

func JsonResponse(rw http.ResponseWriter, data interface{}, status int) {
	body, err := json.Marshal(data)
	if err != nil {
		ErrorResponse(rw, err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	_, _ = rw.Write(body)
}

func OK(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusOK)
}

package handlers

import (
	"go-template/src/common"
	"net/http"
)

type TestResponse struct {
	Hello string `json:"hello"`
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	response := TestResponse{
		Hello: "hello",
	}

	common.JsonResponse(w, response, 200)
}

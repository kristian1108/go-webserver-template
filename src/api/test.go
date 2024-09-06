package api

import (
	"go-template/src/common"
	"net/http"
)

type TestResponse struct {
	Hello string `json:"hello"`
}

func (a *Api) TestHandler(w http.ResponseWriter, r *http.Request) {
	response := TestResponse{
		Hello: "hello",
	}

	common.JsonResponse(w, response, 200)
}

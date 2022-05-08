package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newResponse(Message string, data interface{}) *response {
	return &response{Message: Message, Data: data}
}

func ResponseWriter(res http.ResponseWriter, statusCode int, message string, data interface{}) error {
	res.WriteHeader(statusCode)
	httpResponse := newResponse(message, data)
	err := json.NewEncoder(res).Encode(httpResponse)
	return err
}

package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(res http.ResponseWriter, statusCode int, data interface{}) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(data)
}

package utils

import(
	"net/http"
	"encoding/json"
)

func ResponseJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, err RestErr) {
	ResponseJson(w, err.Status, err)
}


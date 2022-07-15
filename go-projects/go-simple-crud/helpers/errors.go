package helpers

import (
	"encoding/json"
	"net/http"
)

// errors handling for response
func ResponseWithErrs(w http.ResponseWriter, status int, msg string) {
	ResponseWithJSON(w, status, map[string]string{"error": msg})
}

func ResponseWithJSON(w http.ResponseWriter, status int, data interface{}) {
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(result)
}

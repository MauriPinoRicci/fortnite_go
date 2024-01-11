package utils

import (
	"encoding/json"
	"net/http"
)

func RenderResponse(w http.ResponseWriter, r *http.Request, status int, obj interface{}) {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(objBytes)
}

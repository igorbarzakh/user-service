package httpx

import (
	"encoding/json"
	"net/http"
)


func WriteError(w http.ResponseWriter,  message string, status int,) {                                                    
	w.Header().Set("Content-Type", "application/json")                                                                  
	w.WriteHeader(status)                                                                                               
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
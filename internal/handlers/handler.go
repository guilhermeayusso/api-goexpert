package handlers

import (
    "net/http"
    "encoding/json"
)

// HandleGet processes GET requests
func HandleGet(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // Example response
    response := map[string]string{"message": "GET request successful"}
    json.NewEncoder(w).Encode(response)
}

// HandlePost processes POST requests
func HandlePost(w http.ResponseWriter, r *http.Request) {
    var requestData map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    response := map[string]string{"message": "POST request successful"}
    json.NewEncoder(w).Encode(response)
}
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// GetUpdatesHandler handles GET requests to fetch updates
func GetUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve updates from the database
	updates, err := GetUpdates()
	if err != nil {
		http.Error(w, "Failed to retrieve updates", http.StatusInternalServerError)
		return
	}

	// Encode updates as JSON and send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updates); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// SubmitUpdateHandler handles POST requests to submit new updates
func SubmitUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var update StandupUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set current time as the creation time of the update
	update.CreatedAt = time.Now()

	// Insert the update into the database
	if err := InsertUpdate(update); err != nil {
		http.Error(w, "Failed to submit update", http.StatusInternalServerError)
		return
	}

	// Return the updated list of updates
	GetUpdatesHandler(w, r)
}

package standup

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// StandupUpdate represents a team member's daily standup update
type StandupUpdate struct {
	ID        int       `json:"id"`
	UserID    string    `json:"userId"`
	Update    string    `json:"update"`
	Blockers  string    `json:"blockers"`
	CreatedAt time.Time `json:"createdAt"`
}

var updates []StandupUpdate
var updatesLock sync.Mutex

func GetUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	updatesLock.Lock()
	defer updatesLock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updates); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func SubmitUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var update StandupUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	update.CreatedAt = time.Now()

	updatesLock.Lock()
	defer updatesLock.Unlock()

	// Assign a simple ID (replace this with a more robust ID generation mechanism)
	update.ID = len(updates) + 1

	// Add the update to the in-memory slice
	updates = append(updates, update)

	GetUpdatesHandler(w, r) // Return the updated list of updates
}

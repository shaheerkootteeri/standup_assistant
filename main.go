package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

//go:embed html/*
var content embed.FS

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

func getUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	updatesLock.Lock()
	defer updatesLock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updates); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func submitUpdateHandler(w http.ResponseWriter, r *http.Request) {
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

	getUpdatesHandler(w, r) // Return the updated list of updates
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Read and serve the embedded HTML file
	file, err := content.Open("html/index.html")
	if err != nil {
		http.Error(w, "Failed to read HTML file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	htmlContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read HTML file content", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/getUpdates", getUpdatesHandler).Methods("GET")
	r.HandleFunc("/submitUpdate", submitUpdateHandler).Methods("POST")

	// Serve static files (CSS, JS, etc.)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(content))))

	// Start the server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

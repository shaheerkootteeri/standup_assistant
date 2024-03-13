package standup

import "github.com/gorilla/mux"

// NewRouter creates a new router and defines routes for API endpoints
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Define routes for GET and POST requests
	router.HandleFunc("/updates", GetUpdatesHandler).Methods("GET")
	router.HandleFunc("/submit", SubmitUpdateHandler).Methods("POST")

	return router
}

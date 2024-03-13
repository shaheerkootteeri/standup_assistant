package main

import (
	"net/http"

	"github.com/shaheerkootteeri/standup_assistant/standup" // Importing the standup package
)

func main() {
	router := standup.NewRouter()
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandlerReady is the function executed when a HTTP GET request
// is made in the server
func HandlerReady(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CheckDeploy())
}

// StartServer starts a new server connection in localhost
// and port 8080. It also adds a handler to a GET request for
// the endpoint "/ready"
func StartServer() {
	http.HandleFunc("/ready", HandlerReady)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"net/http"
	
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// This handler is used to check if the server is ready to accept requests.
	// It can be used for health checks in production environments.
	respondwithJSON(w, 200, struct{}{})
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("OK")) 
}
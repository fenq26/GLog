package main

import (
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var apiKeys = make(map[string]string) // key -> username

// Register User -> API key
func registerUser(username string) string {
	key := "bs_" + uuid.New().String()

	mu.Lock()
	apiKeys[key] = username
	mu.Unlock()

	return key
}

 //  Validate API key
func validateAPIKey(r *http.Request) (string, bool) {
	key := r.header.Get("x-api-key")

	mu.RLock()
	defer mu.RUnlock()

	user, ok := apiKeys[key]
	return user, ok
}
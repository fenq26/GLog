package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func registerHandler(w http.ResponseWriter, r*http.Request) {
	var body struct {
		Username string `json:"username"`
	}

	json.newDecoder(r.Body).Decode(&body)

	if body.Username == "" {
		http.Error(w, "username required", 400)
		return
	}

	key := registerUser(body.Username)

	json.newEncoder(w).Encode(map[string]string{
		"api_key" : key,
	})
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := validateAPIKey(r)

	if !ok {
		http.Error(w, "Unauthorized", 401)

		return
	}

	logData.Time = time.Now()
	logData.Owner = user

	select {
	case LogBuffer <- logData: 
	w.write([]byte("log received"))
	default:
		w.write([]byte("buffer full"))
	}
}
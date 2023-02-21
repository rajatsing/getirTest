package inmemory

import (
	"encoding/json"
	"fmt"
	"getir/models"
	"log"
	"net/http"
)

type LocalMemory struct {
	Data map[string]string
}

// GetInMemoryHandler function will be responsible for handling the in memory requests

func (local *LocalMemory) GetInMemoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var payload models.Payload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		local.Data[payload.Key] = payload.Value
		response, err := json.Marshal(payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", response)
	} else if r.Method == "GET" {
		key := r.URL.Query().Get("key")

		value, ok := local.Data[key]

		if !ok {
			log.Println("Key not found: ", key)
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}

		var payload models.Payload

		payload.Key = key
		payload.Value = value

		response, err := json.Marshal(payload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Response: ", string(response))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", response)
	}
}

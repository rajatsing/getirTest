package writeresponse

import (
	"encoding/json"
	"net/http"
)

// WriteResponse function will be responsible for writing the response to the client

func WriteResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

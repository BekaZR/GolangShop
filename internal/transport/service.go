package transport

import (
	"encoding/json"
	"log"
	"net/http"
)

// writeJSON provides function to format output response in JSON
func WriteJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error Parsing JSON")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func DecodePost(r *http.Request, structure interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Println("Error parsing post data")
	}
}

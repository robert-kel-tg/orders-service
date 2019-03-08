package response

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON ...
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWith200 ...
func RespondWith200(w http.ResponseWriter, payload interface{}) {
	RespondWithJSON(w, http.StatusOK, payload)
}

// RespondWith201 ...
func RespondWith201(w http.ResponseWriter, payload interface{}) {
	RespondWithJSON(w, http.StatusCreated, payload)
}

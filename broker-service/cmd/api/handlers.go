package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

<<<<<<< HEAD
	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
=======
	_ = app.writeJSON(w, http.StatusOK, payload)
>>>>>>> 3554a1a81cdaf14d021ad355dc82054a4a0b0e56
}

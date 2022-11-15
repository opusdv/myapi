package internal

import (
	"encoding/json"
	"net/http"
)

type DriveMessage struct {
	Message string
}

type DriveHendler struct {
	Message string
}

func (dh DriveHendler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := DriveMessage{
		Message: "test",
	}
	jsonData, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "ivalid json marshal", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonData)
}

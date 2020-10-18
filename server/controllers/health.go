package controllers

import (
	"encoding/json"
	"net/http"
)

func GetHealth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

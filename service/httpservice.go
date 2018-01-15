package service

import (
	"encoding/json"
	"log"
	"net/http"
)

// Top5Stats returns a http.HandlerFunc which always responds with the last
// known top5 words and letters.
func Top5Stats(top5 *Top5WordLetterStat) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(top5.Stat()); err != nil {
			log.Printf("http error: failed to write top5 stat for request from %+q: %+s", r.RemoteAddr, err)
			return
		}
	}
}

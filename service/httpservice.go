package service

import (
	"encoding/json"
	"log"
	"net/http"
)

// Top5Stats returns a http.HandlerFunc which always responds with the last
// known top5 words and letters.
func Top5Stats(top5 *Top5WordLetterStat) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(writer).Encode(top5.Stat()); err != nil {
			log.Printf("http error: failed to write top5 stat for request from %+q: %+s", request.RemoteAddr, err)
		}
	}
}

package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var (
	bufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}
)

// Top5Stats returns a http.HandlerFunc which always responds with the last
// known top5 words and letters.
func Top5Stats(top5 *Top5WordLetterStat) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buffer := bufferPool.Get().(*bytes.Buffer)
		defer bufferPool.Put(buffer)

		if err := json.NewEncoder(buffer).Encode(top5.Stat()); err != nil {
			log.Printf("http error: failed to write top5 stat for request from %+q: %+s", r.RemoteAddr, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		buffer.WriteTo(w)
	}
}

package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/influx6/faux/tests"
)

var (
	basicSentence = "Miss. Greg left with me and we met at the cafe and went home. After which I got me some coffee; but Miss. Greg appeared and kissed me before I knew we were married"
	basicWords    = lexSentence(basicSentence)
)

func TestTop5StatsHTTPHandler(t *testing.T) {
	letters := NewLetterCounter()
	words := NewWordCounter()

	for _, word := range basicWords {
		letters.Compute(word)
		words.Compute(word)
	}

	ltstat, lttotal := letters.Stat()
	wdstat, wdtotal := words.Stat()

	top5 := new(Top5WordLetterStat)
	top5.Update(FreshStat{
		Words:        wdstat,
		Letters:      ltstat,
		TotalWords:   wdtotal,
		TotalLetters: lttotal,
	})

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		tests.FailedWithError(err, "Should have created request")
	}
	tests.Passed("Should have created request")

	router := mux.NewRouter()
	router.Path("/stats").HandlerFunc(Top5Stats(top5)).Methods("GET")

	router.ServeHTTP(res, req)

	validateStandardResponse(res)

	var stat Top5Stat
	if err := json.NewDecoder(res.Body).Decode(&stat); err != nil {
		tests.FailedWithError(err, "Should have recieved valid json response")
	}
	tests.Passed("Should have recieved valid json response")

	if stat.Count != top5Expected.Count {
		tests.Info("Received: %d", stat.Count)
		tests.Info("Expected: %d", top5Expected.Count)
		tests.Failed("Should have gotten exact count stats from http service")
	}
	tests.Passed("Should have gotten exact count stats from http service")

	if !compareSlices(stat.Letters, top5Expected.Letters) {
		tests.Info("Received: %+q", stat.Letters)
		tests.Info("Expected: %+q", top5Expected.Letters)
		tests.Failed("Should have gotten exact letters stats from http service")
	}
	tests.Passed("Should have gotten exact letters stats from http service")

	if !compareSlices(stat.Words, top5Expected.Words) {
		tests.Info("Received: %+q", stat.Words)
		tests.Info("Expected: %+q", top5Expected.Words)
		tests.Failed("Should have gotten exact words stats from http service")
	}
	tests.Passed("Should have gotten exact words stats from http service")

}

func validateStandardResponse(res *httptest.ResponseRecorder) {
	if res.Code != http.StatusOK {
		tests.Failed("Should have recieved http.StatusOk from service")
	}
	tests.Passed("Should have recieved http.StatusOk from service")

	if !strings.Contains(res.HeaderMap.Get("Content-Type"), "application/json") {
		tests.Failed("Should have received content type of 'application/json'")
	}
	tests.Passed("Should have received content type of 'application/json'")

	if res.Body == nil {
		tests.Failed("Should have received respond body")
	}
	tests.Passed("Should have received respond body")

	if res.Body.Len() == 0 {
		tests.Failed("Should have data pending in response body writer")
	}
	tests.Passed("Should have data pending in response body writer")
}

package service_test

import (
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/influx6/wordies/service"
)

func TestLetterCount(t *testing.T) {
	expected := map[int][]string{2: []string{"B", "K", "P"}, 3: []string{"C"}, 26: []string{"E"}, 6: []string{"D", "F"}, 7: []string{"S", "W"}, 9: []string{"A", "M", "T"}, 8: []string{"I", "R"}, 1: []string{"L", "U"}, 5: []string{"G", "H", "N", "O"}, 0: []string{"J", "Q", "V", "X", "Y", "Z"}}

	counter := service.NewLetterCounter()
	for _, word := range basicWords {
		counter.Compute(word)
	}

	stat, total := counter.Stat()
	if total != 26 {
		tests.Info("Expected: %d", 21)
		tests.Info("Received: %d", total)
		tests.Failed("Should have counted 21 words")
	}
	tests.Passed("Should have counted 21 words")

	for count, letters := range stat {
		if !compareSlices(letters, expected[count]) {
			tests.Info("Expected: %+q", expected[count])
			tests.Info("Received: %+q", letters)
			tests.Failed("Should have matched expected frequencies for %d", count)
		}
	}

	tests.Passed("Should have matched all letter occurrence")
}

func TestWordCounter(t *testing.T) {
	expected := map[int][]string{1: []string{"after", "appeared", "at", "before", "but", "cafe", "coffee", "got", "home", "kissed", "knew", "left", "married", "met", "some", "the", "went", "were", "which", "with"}, 2: []string{"greg", "i", "miss.", "we"}, 3: []string{"and", "me"}}

	counter := service.NewWordCounter()
	for _, word := range basicWords {
		counter.Compute(word)
	}

	stat, total := counter.Stat()
	if total != 26 {
		tests.Info("Expected: %d", 26)
		tests.Info("Received: %d", total)
		tests.Failed("Should have counted 26 words")
	}
	tests.Passed("Should have counted 26 words")

	for count, letters := range stat {
		if !compareSlices(letters, expected[count]) {
			tests.Info("Expected: %+q", expected[count])
			tests.Info("Received: %+q", letters)
			tests.Failed("Should have matched expected frequencies for %d", count)
		}
	}

	tests.Passed("Should have matched all word occurrence")
}

func compareSlices(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}

	for _, item := range first {
		var found bool
		for _, target := range second {
			if target == item {
				found = true
			}
		}
		if !found {
			return false
		}
	}

	return true
}

package internal_test

import (
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/influx6/wordies/internal"
)

func TestLexSentence(t *testing.T) {
	testLexSentenceWithPeriods(t)
	testLexSentenceWithMultiSentences(t)
	testLexSentenceWithAbbreviationsAndClosedPeriod(t)
}

func testLexSentenceWithPeriods(t *testing.T) {
	sentence := "This, that, and the other.We've decided to run-before the marathon Dr..."
	expected := []string{"This", "that", "and", "the", "other", "We've", "decided", "to", "run-before", "the", "marathon", "Dr."}

	words := internal.LexSentence(sentence)
	for index, word := range words {
		if word != expected[index] {
			tests.Failed("Should have matched word %q at index %d", word, index)
		}
	}
	tests.Passed("Should have successfully matched all lexed words")
}

func testLexSentenceWithMultiSentences(t *testing.T) {
	sentence := "This became the 20th day of the year. We went to a club house!"
	expected := []string{"This", "became", "the", "20th", "day", "of", "the", "year", "We", "went", "to", "a", "club", "house"}

	words := internal.LexSentence(sentence)
	for index, word := range words {
		if word != expected[index] {
			tests.Failed("Should have matched word %q at index %d", word, index)
		}
	}
	tests.Passed("Should have successfully matched all lexed words")
}

func testLexSentenceWithAbbreviationsAndClosedPeriod(t *testing.T) {
	sentence := "Miss. Greg left.Dir went home bef. we met."
	expected := []string{"Miss.", "Greg", "left", "Dir", "went", "home", "bef.", "we", "met"}

	words := internal.LexSentence(sentence)
	for index, word := range words {
		if word != expected[index] {
			tests.Failed("Should have matched word %q at index %d", word, index)
		}
	}
	tests.Passed("Should have successfully matched all lexed words")
}

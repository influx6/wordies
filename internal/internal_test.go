package internal_test

import (
	"sync"
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/influx6/wordies/internal"
)

func TestLexSentenceWithSplicer(t *testing.T) {
	testLexSentenceSplicerWithPeriods(t)
	testLexSentenceWithSplicerWithMultiSentences(t)
	testLexSentenceWithSplicerWithAbbreviationsAndClosedPeriod(t)
}

func BenchmarkWordSplicer(b *testing.B) {
	b.ReportAllocs()

	var wg sync.WaitGroup
	sentence := "This, that, and the other.We've decided to run-before the marathon Dr..."

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := make(chan string, 0)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range res {
			}
		}()
		internal.WordSplicer(sentence, res)
	}

	b.StopTimer()
	wg.Wait()
}

func testLexSentenceSplicerWithPeriods(t *testing.T) {
	sentence := "This, that, and the other.We've decided to run-before the marathon Dr..."
	expected := []string{"This", "that", "and", "the", "other", "We've", "decided", "to", "run-before", "the", "marathon"}

	words := lexSentenceWithSplicer(sentence)
	for index, word := range expected {
		if word != words[index] {
			tests.Failed("Should have matched word %q at index %d", word, index)
		}
	}
	tests.Passed("Should have successfully matched all lexed words")
}

func testLexSentenceWithSplicerWithMultiSentences(t *testing.T) {
	sentence := "This became the 20th day of the year. We went to a club house!"
	expected := []string{"This", "became", "the", "20th", "day", "of", "the", "year", "We", "went", "to", "a", "club", "house"}

	words := lexSentenceWithSplicer(sentence)
	for index, word := range expected {
		if word != words[index] {
			tests.Failed("Should have matched word %q at index %d", word, index)
		}
	}
	tests.Passed("Should have successfully matched all lexed words")
}

func testLexSentenceWithSplicerWithAbbreviationsAndClosedPeriod(t *testing.T) {
	sentence := "Miss. Greg left.Dir went home bef. we met.!"
	expected := []string{"Miss.", "Greg", "left", "Dir", "went", "home", "bef.", "we", "met"}

	words := lexSentenceWithSplicer(sentence)
	for index, word := range expected {
		if word != words[index] {
			tests.Failed("Should have matched word %q at index %d", word, index)
		}
	}
	tests.Passed("Should have successfully matched all lexed words")
}

// lexSentenceWithSplicer transforms giving sentences into individual
// words using WordSplicer undeneath.
func lexSentenceWithSplicer(sentences string) []string {
	var words []string
	res := make(chan string, 0)
	go internal.WordSplicer(sentences, res)
	for word := range res {
		words = append(words, word)
	}
	return words
}

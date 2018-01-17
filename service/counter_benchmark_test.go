package service_test

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/influx6/wordies/internal"
	"github.com/influx6/wordies/service"
)

var (
	sample300Words   = lexSentenceWithSplicer(fake.SentencesN(300))
	sample3000Words  = lexSentenceWithSplicer(fake.SentencesN(3000))
	sample10000Words = lexSentenceWithSplicer(fake.SentencesN(10000))
	sample50000Words = lexSentenceWithSplicer(fake.SentencesN(50000))
)

func benchmarkLetterFrequency(b *testing.B, samples []string) {
	b.StopTimer()
	b.ReportAllocs()

	counter := service.NewLetterCounter()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range samples {
			counter.Compute(word)
		}
	}
	b.StopTimer()
}

func benchmarkCountFrequency(b *testing.B, samples []string) {
	b.StopTimer()
	b.ReportAllocs()

	wordCounter := service.NewWordCounter()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range samples {
			wordCounter.Compute(word)
		}
	}
	b.StopTimer()
}

func BenchmarkLetterFrequency300(b *testing.B) {
	benchmarkLetterFrequency(b, sample300Words)
}

func BenchmarkLetterFrequency3000(b *testing.B) {
	benchmarkLetterFrequency(b, sample3000Words)
}

func BenchmarkLetterFrequency10000(b *testing.B) {
	benchmarkLetterFrequency(b, sample10000Words)
}

func BenchmarkLetterFrequency50000(b *testing.B) {
	benchmarkLetterFrequency(b, sample50000Words)
}

func BenchmarkCountFrequency300(b *testing.B) {
	benchmarkCountFrequency(b, sample300Words)
}

func BenchmarkCountFrequency3000(b *testing.B) {
	benchmarkCountFrequency(b, sample3000Words)
}

func BenchmarkCountFrequency10000(b *testing.B) {
	benchmarkCountFrequency(b, sample10000Words)
}

func BenchmarkCountFrequency50000(b *testing.B) {
	benchmarkCountFrequency(b, sample50000Words)
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

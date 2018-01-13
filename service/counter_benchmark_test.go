package service

import (
	"testing"

	"github.com/icrowley/fake"
)

var (
	sample300Words   = lexSentence(fake.SentencesN(300))
	sample3000Words  = lexSentence(fake.SentencesN(3000))
	sample10000Words = lexSentence(fake.SentencesN(10000))
	sample50000Words = lexSentence(fake.SentencesN(50000))
)

func benchmarkLetterFrequency(b *testing.B, samples []string) {
	b.StopTimer()
	b.ReportAllocs()

	counter := NewLetterCounter()

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

	wordCounter := NewWordCounter()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range samples {
			wordCounter.Compute(word)
		}
	}
	b.StopTimer()
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

package service_test

import (
	"log"
	"os"
	"testing"

	"github.com/dgraph-io/badger"
	"github.com/icrowley/fake"
	"github.com/influx6/wordies/internal"
	"github.com/influx6/wordies/service"
)

var (
	sample300Words   = internal.LexSentence(fake.SentencesN(300))
	sample3000Words  = internal.LexSentence(fake.SentencesN(3000))
	sample10000Words = internal.LexSentence(fake.SentencesN(10000))
	sample50000Words = internal.LexSentence(fake.SentencesN(50000))
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

func benchmarkBadgerFrequency(b *testing.B, samples []string) {
	b.SkipNow()
	b.StopTimer()
	b.ReportAllocs()

	dbPath := "./artifacts"
	defer os.RemoveAll(dbPath)

	op := badger.DefaultOptions
	op.Dir = dbPath
	op.ValueDir = dbPath
	db, err := badger.Open(op)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	wordCounter := service.NewBadgerWordCounter(db)

	b.StartTimer()
	for _, word := range samples {
		wordCounter.Compute(word)
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

func BenchmarkBadgerFrequency300(b *testing.B) {
	benchmarkBadgerFrequency(b, sample300Words)
}

func BenchmarkBadgerFrequency3000(b *testing.B) {
	benchmarkBadgerFrequency(b, sample3000Words)
}

func BenchmarkBadgerFrequency10000(b *testing.B) {
	benchmarkBadgerFrequency(b, sample10000Words)
}

func BenchmarkBadgerFrequency50000(b *testing.B) {
	benchmarkBadgerFrequency(b, sample50000Words)
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

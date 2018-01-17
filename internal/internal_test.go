package internal_test

import (
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/influx6/wordies/internal"
)

func BenchmarkLexSentence(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	sentence := "This, that, and the other.We've decided to run-before the marathon Dr..."
	for i := 0; i < b.N; i++ {
		internal.LexSentence(sentence)
	}

	b.StopTimer()
}

func TestLexSentence(t *testing.T) {
	testLexSentenceWithPeriods(t)
	testLexSentenceWithMultiSentences(t)
	testLexSentenceWithAbbreviationsAndClosedPeriod(t)
}

func testLexSentenceWithPeriods(t *testing.T) {
	sentence := "This, that, and the other.We've decided to run-before the marathon Dr..."
	expected := []string{"This", "that", "and", "the", "other", "decided", "to", "run-before", "the", "marathon", "Dr.", "We've"}

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
	expected := []string{"Miss.", "Greg", "left", "went", "home", "bef.", "we", "met", "Dir"}

	words := internal.LexSentence(sentence)
	for index, word := range words {
		if word != expected[index] {
			tests.Failed("Should have matched word %q at index %d", word, index)
		}
	}
	tests.Passed("Should have successfully matched all lexed words")
}

//func TestLexSentenceWithSplicer(t *testing.T) {
//	testLexSentenceSplicerWithPeriods(t)
//	testLexSentenceWithSplicerWithMultiSentences(t)
//	testLexSentenceWithSplicerWithAbbreviationsAndClosedPeriod(t)
//}

//func BenchmarkLexSentenceWithSplicer(b *testing.B) {
//	b.ResetTimer()
//	b.ReportAllocs()
//
//	sentence := "This, that, and the other.We've decided to run-before the marathon Dr..."
//	for i := 0; i < b.N; i++ {
//		internal.LexSentenceWithSplicer(sentence)
//	}
//
//	b.StopTimer()
//}
//
//func BenchmarkWordSplicer(b *testing.B) {
//	b.ReportAllocs()
//
//	var wg sync.WaitGroup
//	sentence := "This, that, and the other.We've decided to run-before the marathon Dr..."
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		res := make(chan string, 0)
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for range res {
//			}
//		}()
//		internal.WordSplicer(sentence, res)
//	}
//
//	b.StopTimer()
//	wg.Wait()
//}
//
//func testLexSentenceSplicerWithPeriods(t *testing.T) {
//	sentence := "This, that, and the other.We've decided to run-before the marathon Dr..."
//	expected := []string{"This", "that", "and", "the", "other", "decided", "to", "run-before", "the", "marathon", "Dr.", "We've"}
//
//	words := internal.LexSentenceWithSplicer(sentence)
//	fmt.Printf("%#v\n", words)
//	for index, word := range words {
//		if word != expected[index] {
//			tests.Failed("Should have matched word %q at index %d", word, index)
//		}
//	}
//	tests.Passed("Should have successfully matched all lexed words")
//}
//
//func testLexSentenceWithSplicerWithMultiSentences(t *testing.T) {
//	sentence := "This became the 20th day of the year. We went to a club house!"
//	expected := []string{"This", "became", "the", "20th", "day", "of", "the", "year", "We", "went", "to", "a", "club", "house"}
//
//	words := internal.LexSentenceWithSplicer(sentence)
//	fmt.Printf("%#v\n", words)
//	for index, word := range words {
//		if word != expected[index] {
//			tests.Failed("Should have matched word %q at index %d", word, index)
//		}
//	}
//	tests.Passed("Should have successfully matched all lexed words")
//}
//
//func testLexSentenceWithSplicerWithAbbreviationsAndClosedPeriod(t *testing.T) {
//	sentence := "Miss. Greg left.Dir went home bef. we met."
//	expected := []string{"Miss.", "Greg", "left", "went", "home", "bef.", "we", "met", "Dir"}
//
//	words := internal.LexSentenceWithSplicer(sentence)
//	fmt.Printf("%#v\n", words)
//	for index, word := range words {
//		if word != expected[index] {
//			tests.Failed("Should have matched word %q at index %d", word, index)
//		}
//	}
//	tests.Passed("Should have successfully matched all lexed words")
//}

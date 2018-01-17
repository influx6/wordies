package service

import (
	"sort"
	"strings"

	"sync"

	"unicode"
)

var (
	chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// sortFrequencies returns a slice of ints which contains the sorted
// frequencies counts from the frequency map.
func sortFrequencies(items map[int][]string) []int {
	freqs := make([]int, len(items))
	for counter := range items {
		freqs = append(freqs, counter)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freqs)))
	return freqs
}

// WordCounter implements a simple and slow word counter,
// which internally keeps tracks of all words seen from
// its instantiation.
type WordCounter struct {
	ml    sync.RWMutex
	words map[string]int
}

// NewWordCounter returns a new instance of a word counter.
func NewWordCounter() *WordCounter {
	return &WordCounter{
		words: make(map[string]int),
	}
}

// Stat returns stats of all letters seen from
// computed words since instantiation/creation.
func (wc *WordCounter) Stat() (map[int][]string, int) {
	wc.ml.RLock()
	defer wc.ml.RUnlock()

	freq := make(map[int][]string)
	for word, count := range wc.words {
		freq[count] = append(freq[count], word)
	}

	for _, items := range freq {
		sort.Strings(items)
	}

	return freq, len(wc.words)
}

// Compute generates frequency counts by updating it's
// count of provided word.
func (wc *WordCounter) Compute(word string) error {
	wc.ml.Lock()
	defer wc.ml.Unlock()
	word = strings.ToLower(word)
	wc.words[word]++
	return nil
}

// LetterCounter implements a letter counting mechanism,
// where provided words are tagged against the occurence of
// the standard english alphabets. It returns a map of
// letter with occurrence of all words that have being seen
// since it's initialization.
// LetterCounter is safe for concurrent use.
type LetterCounter struct {
	ml      sync.RWMutex
	letters map[rune]int
}

// NewLetterCounter returns a new instance of the LetterCounter.
func NewLetterCounter() *LetterCounter {
	counter := new(LetterCounter)
	counter.letters = mapLetters()
	return counter
}

// Stat returns stats of all letters seen from
// computed words since instantiation/creation.
func (lc *LetterCounter) Stat() (map[int][]string, int) {
	lc.ml.RLock()
	defer lc.ml.RUnlock()

	freq := make(map[int][]string)
	for letter, counter := range lc.letters {
		freq[counter] = append(freq[counter], string(letter))
	}

	for _, items := range freq {
		sort.Strings(items)
	}
	return freq, len(lc.letters)
}

// Compute generates frequency counts by updating it's
// count of provided word.
func (lc *LetterCounter) Compute(word string) error {
	lc.ml.Lock()
	defer lc.ml.Unlock()

	for _, rune := range word {
		if !unicode.IsLetter(rune) {
			continue
		}

		if !unicode.IsUpper(rune) {
			lc.letters[unicode.ToUpper(rune)]++
			continue
		}

		lc.letters[rune]++
	}
	return nil
}

// mapLetters returns a map of all english alphabets with
// associated int64 counter.
func mapLetters() map[rune]int {
	mapp := map[rune]int{}
	for _, char := range chars {
		mapp[char] = 0
	}
	return mapp
}

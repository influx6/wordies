package service

import (
	"regexp"
	"sort"
	"strings"

	"sync"

	"github.com/influx6/wordies/abbrvs"
)

var (
	chars        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsSet     = strings.Split(chars, "")
	punctuations = ";:,!?"
	mixedDot     = regexp.MustCompile(`\w+\.[\w\d]+`)
	multidot     = regexp.MustCompile(`\.+`)
)

// sortFrequencies returns a slice of ints which contains the sorted
// frequencies keys from the frequency map.
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

// Stat returns a map of frequency to string slicemap generated from
// the internal word map of the word counter.
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

// Compute generates from provided sentences words which it uses to update
// it's words hashmap.
func (wc *WordCounter) Compute(word string) {
	wc.ml.Lock()
	defer wc.ml.Unlock()
	wc.words[word]++
}

// LetterCounter implements a letter counting mechanism,
// where provided words are tagged against the occurence of
// the standard english alphabets. It returns a map of
// letter with occurrence of all words that have being seen
// since it's initialization.
// LetterCounter is safe for concurrent use.
type LetterCounter struct {
	ml      sync.RWMutex
	letters map[string]int
}

// NewLetterCounter returns a new instance of the LetterCounter.
func NewLetterCounter() *LetterCounter {
	counter := new(LetterCounter)
	counter.letters = mapLetters()
	return counter
}

// Stat returns stats of all letters seen from
// computed words since creation.
func (lc *LetterCounter) Stat() (map[int][]string, int) {
	lc.ml.RLock()
	defer lc.ml.RUnlock()

	freq := make(map[int][]string)
	for letter, counter := range lc.letters {
		freq[counter] = append(freq[counter], letter)
	}

	for _, items := range freq {
		sort.Strings(items)
	}
	return freq, len(lc.letters)
}

// Compute handles the update of internal letter stat
// with provided words.
func (lc *LetterCounter) Compute(word string) error {
	lc.ml.Lock()
	defer lc.ml.Unlock()

	word = strings.ToUpper(word)
	for letter, count := range lc.letters {
		newInc := strings.Count(word, letter)
		lc.letters[letter] = count + newInc
	}
	return nil
}

// mapLetters returns a map of all english alphabets with
// associated int64 counter.
func mapLetters() map[string]int {
	mapp := map[string]int{}
	for _, char := range charsSet {
		mapp[char] = 0
	}
	return mapp
}

// lexSentences breaks down provided sentence into it's individual words.
// This provides a simplistic, non complicated version using the golang
// internal package, more sophisticated approaches could be research, or
// consider looking into ML with Textbox (https://machinebox.io/docs/textbox).
func lexSentence(sentence string) []string {
	var words []string

	initials := strings.Fields(strings.Map(replacePunctuations, sentence))
	total := len(initials)
	for index := 0; index < total; index++ {
		word := initials[index]
		if mixedDot.MatchString(word) {
			words = append(words, strings.Split(word, ".")...)
			continue
		}

		if multidot.MatchString(word) {
			word = multidot.ReplaceAllString(word, ".")
			if !abbrvs.IsAbbreviated(word) {
				words = append(words, strings.TrimSuffix(word, "."))
				continue
			}

			if suffixes, ok := abbrvs.GetAllSuffix(word); ok {
				if index+1 >= total && !abbrvs.IsStrictlyAbbreviated(word) {
					words = append(words, strings.TrimSuffix(word, "."))
					continue
				}

				// find a possible suffix match for giving abbreviation.
				// if we hit one, then add next word together with last, else
				// skip op.
				var found bool
				nextWord := initials[index+1]
				for _, suffix := range suffixes {
					if suffix != nextWord {
						continue
					}

					words = append(words, word+" "+nextWord)
					found = true
					index++
					break
				}

				if !found {
					words = append(words, strings.TrimSuffix(word, "."))
				}

				continue
			}
		}

		words = append(words, word)
	}

	return words
}

// replacePunctuations is used to replace all existing punctuations
// except for periods(.), due to special cases for abbreviations.
func replacePunctuations(r rune) rune {
	if strings.ContainsRune(punctuations, r) {
		return -1
	}
	return r
}

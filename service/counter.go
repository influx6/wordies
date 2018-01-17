package service

import (
	"sort"
	"strings"

	"encoding/binary"
	"sync"

	"unicode"

	"github.com/dgraph-io/badger"
)

var (
	chars    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsSet = strings.Split(chars, "")
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

// BadgerWordCounter implements a word counter ontop of badger
// key-value store.
// TODO: Currently benchmark shows power performance here,
// very different from expectation, might be approach of benchmarks
// but badger should provide higher value when dealing with larger
// word sets, will need to look more into this.
type BadgerWordCounter struct {
	db *badger.DB
}

// NewBadgerWordCounter returns a new BadgerWordCounter instance.
func NewBadgerWordCounter(db *badger.DB) *BadgerWordCounter {
	return &BadgerWordCounter{
		db: db,
	}
}

// Stat returns stats of all letters seen from computed words since
// instantiation/creation. It retrieves from the underline badger db.
func (bwc *BadgerWordCounter) Stat() (map[int][]string, int, error) {
	freq := make(map[int][]string)

	var total int
	if err := bwc.db.View(func(tx *badger.Txn) error {
		itr := tx.NewIterator(badger.IteratorOptions{
			PrefetchSize:   120,
			PrefetchValues: true,
		})

		defer itr.Close()

		for itr.Rewind(); itr.Valid(); itr.Next() {
			item := itr.Item()
			value, err := item.Value()
			if err != nil {
				return err
			}

			count := toInt(value)
			freq[count] = append(freq[count], string(item.Key()))
			total++
		}

		return nil
	}); err != nil {
		return nil, 0, err
	}

	for _, items := range freq {
		sort.Strings(items)
	}

	return freq, total, nil
}

// Compute generates frequency counts by updating it's
// count of provided word.
func (bwc *BadgerWordCounter) Compute(word string) error {
	word = strings.ToLower(word)
	wordBytes := []byte(word)
	return bwc.db.Update(func(tx *badger.Txn) error {
		val, err := tx.Get(wordBytes)
		if err != nil {
			if err != badger.ErrKeyNotFound {
				return err
			}

			return tx.Set(wordBytes, fromInt(1))
		}

		lastCountByte, err := val.Value()
		if err != nil {
			return err
		}

		lastCount := toInt(lastCountByte)
		return tx.Set(wordBytes, fromInt(lastCount+1))
	})
}

func fromInt(d int) []byte {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, uint16(d))
	return data
}

func toInt(b []byte) int {
	return int(binary.BigEndian.Uint16(b))
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

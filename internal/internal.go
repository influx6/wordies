package internal

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/influx6/wordies/internal/abbrvs"
)

var (
	dot          = rune('.')
	punctuations = ";:,!?"
	mixedDot     = regexp.MustCompile(`\w+\.[\w\d]+`)
	multidot     = regexp.MustCompile(`\.+`)
)

// WordSplicer breaks down provided sentence into it's individual words.
// This provides a simple approach, but we recommend to use more sophisticated
// approaches. Consider looking into ML built properly  for text processing with
// all rules related to gramma. Such has with Textbox (https://machinebox.io/docs/textbox).
func WordSplicer(sentences string, res chan string) {
	defer close(res)

	var word string
	var lastIndex int
	for index, rune := range sentences {
		if unicode.IsSpace(rune) {
			word = cleanWord(sentences[lastIndex:index])
			lastIndex = index + 1

			if dotIndex := strings.IndexRune(word, dot); dotIndex != -1 && dotIndex+1 < len(word) {
				dotSplicer(word, res)
				continue
			}

			// Probably we dealing with many dot type or single dot, so
			// just treat has having single dot, check abbreviation if
			// its one, if one, get abbreviation else add without dot.
			if strings.HasSuffix(word, ".") {
				if abbrvs.IsAbbreviated(word) {
					res <- word
					continue
				}

				res <- strings.TrimSuffix(word, ".")
				continue
			}

			res <- word
			continue
		}
	}

	if word = sentences[lastIndex:]; len(word) != 0 {
		word = cleanWord(word)

		if abbrvs.IsAbbreviated(word) {
			res <- word
			return
		}

		if strings.HasSuffix(word, ".") {
			res <- strings.TrimSuffix(word, ".")
			return
		}

		res <- word
	}
}

// replacePunctuations is used to replace all existing punctuations
// except for periods(.), due to special cases for abbreviations.
func replacePunctuations(r rune) rune {
	if strings.ContainsRune(punctuations, r) {
		return -1
	}
	return r
}

func dotSplicer(input string, res chan string) {
	var word string
	var lastIndex int
	for index, rune := range input {
		if rune == dot {
			word = input[lastIndex:index]
			lastIndex = index + 1
			res <- cleanWord(word)
			continue
		}
	}

	if word = input[lastIndex:]; len(word) != 0 {
		res <- cleanWord(input[lastIndex:])
	}
}

func cleanWord(input string) string {
	return strings.TrimSpace(strings.Map(replacePunctuations, input))
}

package internal

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/influx6/wordies/internal/abbrvs"
)

var (
	dot          = rune('.')
	empty        = rune(' ')
	hiphen       = rune('\'')
	dash         = rune('-')
	punctuations = ";:,!?"
	mixedDot     = regexp.MustCompile(`\w+\.[\w\d]+`)
	multidot     = regexp.MustCompile(`\.+`)
)

// LexSentences breaks down provided sentence into it's individual words.
// This provides a simplistic, non to optimized version, a more sophisticated
// approaches could be research, or consider looking into ML built properly
// for text processing with all rules related to gramma. Such has with
// Textbox (https://machinebox.io/docs/textbox).
func LexSentence(sentence string) []string {
	initials := strings.Fields(strings.Map(replacePunctuations, sentence))
	total := len(initials)

	var word string
	for index := 0; index < total; index++ {
		word = initials[index]
		if mixedDot.MatchString(word) {
			if parts := strings.Split(word, "."); len(parts) > 0 {
				initials = append(initials, parts[1:]...)
				initials[index] = parts[0]
			}
			continue
		}

		if multidot.MatchString(word) {
			word = multidot.ReplaceAllString(word, ".")
			if abbrvs.IsAbbreviated(word) {
				initials[index] = word
				continue
			}

			initials[index] = strings.TrimSuffix(word, ".")
		}
	}

	return initials
}

// replacePunctuations is used to replace all existing punctuations
// except for periods(.), due to special cases for abbreviations.
func replacePunctuations(r rune) rune {
	if strings.ContainsRune(punctuations, r) {
		return -1
	}
	return r
}

// WordSplicer runs through all unicode runes of giving string
// returning each word seperated by space it finds.
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
	return strings.TrimSpace(strings.Map(func(r rune) rune {
		if strings.ContainsRune(punctuations, r) {
			return -1
		}
		return r
	}, input))
}

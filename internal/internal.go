package internal

import (
	"regexp"
	"strings"

	"github.com/influx6/wordies/internal/abbrvs"
)

var (
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

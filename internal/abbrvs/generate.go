// +build ignore

package abbrvs

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/influx6/faux/tmplutil"
)

var (
	template = `package abbrvs

//go:generate go run generate.go

var (
	abbreviations = map[string][]string{
		{{ range $key, $value := .Abbrvs }}
		{{quote $key}}: []string{ {{range $value}}{{quote .}},{{end}} },
		{{end}}
	}

	suffixes = map[string][]string{
		{{ range $key, $value := .Suffixes }}
		{{quote $key}}: []string{ {{range $value}}{{quote .}},{{end}} },
		{{end}}
	}
)

// IsStrictlyAbbreviated returns true/false if giving text is an abbreviation.
// This function does not provide optimized truthy value based on prefix. It
// strictly validates if this is in itself a abbreviation.
func IsStrictlyAbbreviated(item string) bool {
	_, ok := abbreviations[item]
	return ok
}

// IsAbbreviated returns true/false if giving text is an abbreviation.
// IsAbbreviated provides a optimistic positive return value of true
// if giving abbreviation has other suffix like Song of Sol. etc.
func IsAbbreviated(item string) bool {
	// if we have direct hit, then return true.
	if _, ok := abbreviations[item]; ok {
		return ok
	}


	// We need to attempt indirect hit with combo map.
	// and return optimistic truthy value if any is found.
	if _, ok := suffixes[item]; ok {
		return ok
	}

	
	return false
}

// GetAllSuffix returns a slice of strings that represent
// suffixed types of giving abbreviation word if any exists.
// That is if you provide 'Song.', then we return ['Sol.', 'of Sol.'].
func GetAllSuffix(item string) ([]string, bool) {
	suffixes, ok := suffixes[item]
	return suffixes, ok
}

// Mutations returns a slice of strings that represent
// all possible mutation of giving abbreviation prefix if any.
// That is if you provide 'Song.', then we return ['Song. Sol.', 'Song of Sol.'].
func Mutations(item string) ([]string, bool) {
	if suffixes, ok := suffixes[item]; ok {
		var parts []string
		for _, suffix := range suffixes {
			parts = append(parts, item+" "+suffix)
		}
		return parts, true
	}

	return nil, false
}

// GetMeanings returns all available meanings of given abbreviation if found.
// Bool argument is used to indicate if abbreviation was found.
func GetMeanings(abbr string) ([]string, bool) {
	meanings, ok := abbreviations[abbr]
	return meanings, ok
}

`
)

func main() {
	page, err := goquery.NewDocument("http://public.oed.com/how-to-use-the-oed/abbreviations/")
	if err != nil {
		log.Fatalf("failed to read url: %+s", err)
		return
	}

	abbrevs := make(map[string][]string)
	abbrevs["Dr."] = []string{"Doctor"}
	abbrevs["Mr."] = []string{"Mister"}
	abbrevs["Miss."] = []string{"Miss"}
	abbrevs["Mrs."] = []string{"Misses"}

	combos := make(map[string][]string)

	page.Find("div#content div.col-8 div.page-content table").Each(func(i int, selection *goquery.Selection) {
		selection.Find("tbody").Find("tr").Each(func(trIndex int, trSel *goquery.Selection) {

			// if we have a h2 within td then skip, that is probably header.
			if trSel.Find("td h2").Length() != 0 {
				return
			}

			tds := trSel.Find("td")
			if tds.Length() < 2 {
				return
			}

			abbr := tds.Get(0)
			meaning := tds.Get(1)

			if abbr.FirstChild == nil || meaning.FirstChild == nil {
				return
			}

			abbrVal := abbr.FirstChild.Data
			if parts := strings.Fields(abbrVal); len(parts) > 1 {
				key := parts[0]

				if rest, ok := combos[key]; ok {
					rest = append(rest, strings.Join(parts[1:], " "))
					combos[key] = rest
				} else {
					combos[key] = []string{strings.Join(parts[1:], " ")}
				}
			}

			abbrevs[abbrVal] = strings.Split(meaning.FirstChild.Data, ",")
			fmt.Printf(".")
		})
	})

	fmt.Printf("\n")

	abbrFile, err := os.Create("./abbrvs.go")
	if err != nil {
		log.Fatalf("failed to create file: %+s", err)
		return
	}

	defer abbrFile.Close()

	if err := tmplutil.MustFrom("abbr", template).Execute(abbrFile, struct {
		Abbrvs   map[string][]string
		Suffixes map[string][]string
		//AbbrvsLower   map[string]string
		//SuffixesLower map[string]string
	}{
		Abbrvs:   abbrevs,
		Suffixes: combos,
		//AbbrvsLower:   abbrevsLower,
		//SuffixesLower: comboLower,
	}); err != nil {
		log.Fatalf("ParseError: template parsing failed: %+s", err)
		return
	}
}

package abbrvs_test

import (
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/influx6/wordies/internal/abbrvs"
)

func TestGetMeanings(t *testing.T) {
	meanings, ok := abbrvs.GetMeanings("Att.")
	if !ok {
		tests.Failed("Should have succesffully gotten abbr meanings")
	}
	tests.Passed("Should have succesffully gotten abbr meanings")

	if len(meanings) != 1 {
		tests.Failed("Should have atleast one meaning for abbreviation")
	}
	tests.Passed("Should have atleast one meaning for abbreviation")
}

func TestIsAbbreviated(t *testing.T) {
	if !abbrvs.IsAbbreviated("Att.") {
		tests.Failed("Should have validated abbreviation to be true")
	}
	tests.Passed("Should have validated abbreviation to be true")

	if abbrvs.IsAbbreviated("att.") {
		tests.Failed("Should have failed to validated abbreviation of lowercase to be false")
	}
	tests.Passed("Should have failed to validated abbreviation of lowercase to be false")
}

func TestGetAllSuffix(t *testing.T) {
	suffix, ok := abbrvs.GetAllSuffix("Song")
	if !ok {
		tests.Failed("Should have found suffix values for abbreviation")
	}
	tests.Passed("Should have found suffix values for abbreviation")

	if len(suffix) != 2 {
		tests.Failed("Should have received 2 suffix but got %d", len(suffix))
	}
	tests.Passed("Should have received 2 suffix")

	for ind, val := range []string{"of Sol.", "Sol."} {
		if val != suffix[ind] {
			tests.Failed("Should match suffix at index %d with %q not %q", ind, val, suffix[ind])
		}
	}
}

func TestMutations(t *testing.T) {
	mutations, ok := abbrvs.Mutations("pa.")
	if !ok {
		tests.Failed("Should have found mutations for abbreviation")
	}
	tests.Passed("Should have found mutations for abbreviation")

	if mutations[0] != "pa. pple." {
		tests.Failed("Expected 'pa. pple.' but got %+q", mutations[0])
	}

	if mutations[1] != "pa. t." {
		tests.Failed("Expected 'pa. t.' but got %+q", mutations[1])
	}
}

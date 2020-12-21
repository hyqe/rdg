package rdg

import (
	"testing"
)

func TestNameGenerators(t *testing.T) {
	var nameFunctions = []func() string {
		Prefix,
		Suffix,
		FirstName,
		LastName,
		PrefixMale,
		PrefixFemale,
		SuffixMale,
		SuffixFemale,
		FirstNameMale,
		FirstNameFemale,
	}
	var funcNames = []string {
		"Prefix",
		"Suffix",
		"FirstName",
		"LastName",
		"PrefixMale",
		"PrefixFemale",
		"SuffixMale",
		"SuffixFemale",
		"FirstNameMale",
		"FirstNameFemale",
	}
	var inc int = 0
	for i:=0;i<1000;i++ {
		if inc == len(nameFunctions) {
			inc = 0
		}
		if nameFunctions[inc]() == "" {
			t.Fatalf("function (%v()) failed!\n", funcNames[inc])
		}
		inc = inc + 1
	}
}


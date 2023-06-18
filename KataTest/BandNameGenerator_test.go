package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"testing"
)

func Test_BandNameGenerator(t *testing.T) {
	t.Log("start test", t.Name())

	var testCases = []struct {
		intput   string
		expected string
	}{
		{"knife", "The Knife"},
		{"tart", "Tartart"},
		{"sandles", "Sandlesandles"},
		{"bed", "The Bed"},
	}
	for _, tc := range testCases {
		result := BandNameGenerator(tc.intput)
		if result != tc.expected {
			t.Fatalf("Expected %+v but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

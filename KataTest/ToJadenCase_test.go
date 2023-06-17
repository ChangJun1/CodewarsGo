package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"testing"
)

func Test_ToJadenCase(t *testing.T) {
	t.Log("start test", t.Name())

	var testCases = []struct {
		intput   string
		expected string
	}{
		{"most trees are blue", "Most Trees Are Blue"},
		{"All the rules in this world were made by someone no smarter than you. So make your own.", "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."},
		{"When I die. then you will realize", "When I Die. Then You Will Realize"},
		{"Jonah Hill is a genius", "Jonah Hill Is A Genius"},
		{"Dying is mainstream", "Dying Is Mainstream"},
	}
	for _, tc := range testCases {
		result := ToJadenCase(tc.intput)
		if result != tc.expected {
			t.Fatalf("Expected %+v but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"reflect"
	"testing"
)

func Test_LongestVowelChain(t *testing.T) {
	t.Log("start test", t.Name())

	testCases := []struct {
		input    string
		expected int
	}{
		{"codewarriors", 2},
		{"suoidea", 3},
		{"ultrarevolutionariees", 3},
		{"strengthlessnesses", 1},
		{"cuboideonavicuare", 2},
		{"chrononhotonthuooaos", 5},
		{"iiihoovaeaaaoougjyaw", 8},
	}

	for _, tc := range testCases {
		result := LongestVowelChain(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Fatalf("Expected %d, but got %d", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"testing"
)

func Test_Spacify(t *testing.T) {
	t.Log("start test", t.Name())

	testCases := []struct {
		input    string
		expected string
	}{
		{"hello world", "h e l l o   w o r l d"},
		{"12345", "1 2 3 4 5"},
		{"a", "a"},
	}

	for _, tc := range testCases {
		result := Spacify(tc.input)
		if result != tc.expected {
			t.Fatalf("Expected %s, but got %s", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

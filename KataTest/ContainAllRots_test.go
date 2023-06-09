package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"testing"
)

func Test_ContainAllRots(t *testing.T) {
	t.Log("start test", t.Name())

	var testCases = []struct {
		inputStr string
		inputArr []string
		expected bool
	}{
		{"bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}, true},
		{"XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}, false},
		{"", []string{}, true},
	}
	for _, tc := range testCases {
		result := ContainAllRots(tc.inputStr, tc.inputArr)
		if result != tc.expected {
			t.Fatalf("Expected %+v but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

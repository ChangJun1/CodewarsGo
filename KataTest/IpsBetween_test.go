package KataTest

import (
	. "CodewarsTraining/RankUp"
	"testing"
)

func Test_IpsBetween(t *testing.T) {
	t.Log("start test", t.Name())

	var testCases = []struct {
		start    string
		end      string
		expected int
	}{
		{"10.0.0.0", "10.0.0.50", 50},
		{"10.0.0.0", "10.0.1.0", 256},
		{"20.0.0.10", "20.0.1.0", 246},
	}
	for _, tc := range testCases {
		result := IpsBetween(tc.start, tc.end)
		if result != tc.expected {
			t.Fatalf("Expected %+v, but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

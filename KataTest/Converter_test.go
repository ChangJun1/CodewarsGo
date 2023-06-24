package KataTest

import (
	. "CodewarsTraining/RankUp"
	"math"
	"testing"
)

func Test_Converter(t *testing.T) {
	t.Log("start test", t.Name())

	testCases := []struct {
		n        float64
		decimals int
		base     float64
		expected string
	}{

		{13, 0, math.Pi, "103"},
		{10, 0, math.Pi, "100"},
		{math.Pi, 0, math.Pi, "10"},

		{13, 3, math.Pi, "103.010"},
		{10, 2, math.Pi, "100.01"},
		{math.Pi, 5, math.Pi, "10.00000"},

		{13, 0, 8, "15"},
		{10, 0, 16, "A"},
		{math.Pi, 0, 2, "11"},

		{13, 0, 8, "15"},
		{7, 0, 19, "7"},
		{1, 0, 2, "1"},

		{13.5, 4, 16, "D.8000"},
		{10.5, 0, 16, "A"},
		{1, 2, 2, "1.00"},

		{-10, 0, 23, "-A"},
		{0, 4, 26, "0.0000"},
		{-15.5, 2, 23, "-F.BB"},

		{13, 0, 10, "13"},
		{10, 0, 10, "10"},
		{5.5, 1, 10, "5.5"},
	}

	for _, tc := range testCases {
		result := Converter(tc.n, tc.decimals, tc.base)
		if result != tc.expected {
			t.Fatalf("Expected %s, but got %s", tc.expected, result)
		}
	}
	t.Log("passed all the tests")

}

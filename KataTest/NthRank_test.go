package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"testing"
)

func Test_NthRank(t *testing.T) {
	t.Log("start test", t.Name())

	var testCases = []struct {
		st       string
		we       []int
		n        int
		expected string
	}{
		{
			st:       "Addison,Jayden,Sofia,Michael,Andrew,Lily,Benjamin",
			we:       []int{4, 2, 1, 4, 3, 1, 2},
			n:        4,
			expected: "Benjamin",
		},
		{
			st:       "Elijah,Chloe,Elizabeth,Matthew,Natalie,Jayden",
			we:       []int{1, 3, 5, 5, 3, 6},
			n:        2,
			expected: "Matthew",
		},
	}
	for _, tc := range testCases {
		result := NthRank(tc.st, tc.we, tc.n)
		if result != tc.expected {
			t.Fatalf("Expected %+v but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

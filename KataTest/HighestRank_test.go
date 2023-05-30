package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"testing"
)

func Test_HighestRank(t *testing.T) {
	t.Log("start test", t.Name())

	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{12, 10, 8, 12, 7, 6, 4, 10, 12}, 12},
		{[]int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10}, 12},
		{[]int{12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10}, 3},
	}

	for _, tc := range testCases {
		result := HighestRank(tc.input)
		if result != tc.expected {
			t.Fatalf("Expected %d, but got %d", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

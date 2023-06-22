package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"testing"
)

func Test_Determinant(t *testing.T) {
	t.Log("start test", t.Name())

	testCases := []struct {
		input    [][]int
		expected int
	}{
		{[][]int{{1}}, 1},
		{[][]int{{1, 3}, {2, 5}}, -1},
		{[][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}}, -20},
	}

	for _, tc := range testCases {
		result := Determinant(tc.input)
		if result != tc.expected {
			t.Fatalf("Expected %d, but got %d", tc.expected, result)
		}
	}
	t.Log("passed all the tests")

}

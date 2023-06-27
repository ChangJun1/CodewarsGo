package KataTest

import (
	. "CodewarsTraining/RankUp"
	"reflect"
	"testing"
)

func Test_Spiralize(t *testing.T) {
	t.Log("start test", t.Name())

	result1 := [][]int{{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1},
		{1, 1, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1}}
	result2 := [][]int{{1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1}}

	testCases := []struct {
		intput   int
		expected [][]int
	}{

		{5, result1},
		{8, result2},
	}

	for _, tc := range testCases {
		result := Spiralize(tc.intput)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Fatalf("Expected %+v, but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")

}

package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"reflect"
	"testing"
)

type Input struct {
	arr1 []int
	arr2 []int
}

func Test_MergeArrays(t *testing.T) {
	t.Log("start test", t.Name())

	var testCases = []struct {
		input    Input
		expected []int
	}{
		{Input{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{Input{[]int{1, 3, 5, 7, 9}, []int{10, 8, 6, 4, 2}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{Input{[]int{1, 3, 5, 7, 9, 11, 12}, []int{1, 2, 3, 4, 5, 10, 12}}, []int{1, 2, 3, 4, 5, 7, 9, 10, 11, 12}},
	}
	for _, tc := range testCases {
		result := MergeArrays(tc.input.arr1, tc.input.arr2)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Fatalf("Expected %+v but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

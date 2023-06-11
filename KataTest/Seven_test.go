package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"reflect"
	"testing"
)

func Test_Seven(t *testing.T) {
	t.Log("start test", t.Name())

	var testCases = []struct {
		input    int64
		expected []int
	}{
		{477557101, []int{28, 7}},
		{477557102, []int{47, 7}},
	}
	for _, tc := range testCases {
		result := Seven(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Fatalf("Expected %+v but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

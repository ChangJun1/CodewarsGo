package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"reflect"
	"testing"
)

func Test_NameValue(t *testing.T) {
	t.Log("start test", t.Name())

	testCases := []struct {
		input    []string
		expected []int
	}{
		{[]string{"abc", "abc", "abc", "abc"}, []int{6, 12, 18, 24}},
		{[]string{"codewars", "abc", "xyz"}, []int{88, 12, 225}},
	}

	for _, tc := range testCases {
		result := NameValue(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Fatalf("Expected %+v, but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

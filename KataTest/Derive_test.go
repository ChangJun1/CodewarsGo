package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"reflect"
	"testing"
)

func Test_Derive(t *testing.T) {
	t.Log("start test", t.Name())

	testCases := []struct {
		name     string
		input    []int
		expected string
	}{
		{"Test Case 1", []int{7, 8}, "56x^7"},
		{"Test Case 2", []int{5, 9}, "45x^8"},
	}

	for _, tc := range testCases {
		result := Derive(tc.input[0], tc.input[1])
		if !reflect.DeepEqual(result, tc.expected) {
			t.Fatalf("Expected %s, but got %s", tc.expected, result)
		}
	}
	t.Log("passed all the tests")

}

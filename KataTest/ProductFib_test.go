package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"reflect"
	"testing"
)

func Test_ProductFib(t *testing.T) {
	t.Log("start test", t.Name())

	var testCases = []struct {
		intput   uint64
		expected [3]uint64
	}{
		{714, [3]uint64{21, 34, 1}},
		{4895, [3]uint64{55, 89, 1}},
		{5895, [3]uint64{89, 144, 0}},
		{74049690, [3]uint64{6765, 10946, 1}},
		{84049690, [3]uint64{10946, 17711, 0}},
	}
	for _, tc := range testCases {
		result := ProductFib(tc.intput)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Fatalf("Expected %+v but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

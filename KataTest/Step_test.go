package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"reflect"
	"testing"
)

type input struct {
	g int
	m int
	n int
}

func Test_Step(t *testing.T) {
	t.Log("start test", t.Name())

	var testCases = []struct {
		intput   input
		expected []int
	}{
		{input{2, 100, 110}, []int{101, 103}},
		{input{4, 100, 110}, []int{103, 107}},
		{input{6, 100, 110}, []int{101, 107}},
		{input{8, 300, 400}, []int{359, 367}},
		{input{10, 300, 400}, []int{307, 317}},
		{input{11, 30000, 100000}, nil},
	}
	for _, tc := range testCases {
		result := Step(tc.intput.g, tc.intput.m, tc.intput.n)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Fatalf("Expected %+v but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")
}

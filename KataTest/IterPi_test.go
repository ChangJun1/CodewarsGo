package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"testing"
)

type expected struct {
	iteration    int
	appoximation string
}

func Test_IterPi(t *testing.T) {
	t.Log("start test", t.Name())

	var testCases = []struct {
		epsilon float64
		res     expected
	}{
		{0.1, expected{10, "3.0418396189"}},
		{0.01, expected{100, "3.1315929036"}},
		{0.001, expected{1000, "3.1405926538"}},
		{0.0001, expected{10000, "3.1414926536"}},
		{0.00001, expected{100001, "3.1416026535"}},
		{0.000001, expected{1000001, "3.1415936536"}},
		{0.05, expected{20, "3.0916238067"}},
	}
	for _, tc := range testCases {
		iteration, approx := IterPi(tc.epsilon)
		if iteration != tc.res.iteration || approx != tc.res.appoximation {
			t.Fatalf("Expected %+v, %+v but got %+v, %+v", tc.res.iteration, tc.res.appoximation, iteration, approx)
		}
	}
	t.Log("passed all the tests")
}

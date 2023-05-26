package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"reflect"
	"testing"
)

func Test_FindMultiples(t *testing.T) {
	t.Log("start test", t.Name())
	actual := FindMultiples(5, 25)
	expected := []int{5, 10, 15, 20, 25}

	if reflect.DeepEqual(actual, expected) {
		t.Log("passed all the tests")
	} else {
		t.Log("failed the tests")
	}

}

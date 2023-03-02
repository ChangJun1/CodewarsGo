package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"testing"
)

func Test_Disemvowel(t *testing.T) {
	t.Log("start test", t.Name())
	actual := Disemvowel("This website is for losers LOL!")
	expected := "Ths wbst s fr lsrs LL!"

	if actual != expected {
		t.Errorf("failed to pass the test, which actual is \"%s\", and expected is \"%s\" ", actual, expected)
	} else {
		t.Log("passed all the tests")
	}

}

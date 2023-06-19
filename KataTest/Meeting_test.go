package KataTest

import (
	. "CodewarsTraining/Fundamental"
	"testing"
)

func Test_Meeting(t *testing.T) {
	t.Log("start test", t.Name())

	testCases := []struct {
		input    string
		expected string
	}{
		{"Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn",
			"(ARNO, ANN)(BELL, JOHN)(CORNWELL, ALEX)(DORNY, ABBA)(KERN, LEWIS)(KORN, ALEX)(META, GRACE)(SCHWARZ, VICTORIA)(STAN, MADISON)(STAN, MEGAN)(WAHL, ALEXIS)"},
		{"John:Gates;Michael:Wahl;Megan:Bell;Paul:Dorries;James:Dorny;Lewis:Steve;Alex:Meta;Elizabeth:Russel;Anna:Korn;Ann:Kern;Amber:Cornwell",
			"(BELL, MEGAN)(CORNWELL, AMBER)(DORNY, JAMES)(DORRIES, PAUL)(GATES, JOHN)(KERN, ANN)(KORN, ANNA)(META, ALEX)(RUSSEL, ELIZABETH)(STEVE, LEWIS)(WAHL, MICHAEL)"},
	}

	for _, tc := range testCases {
		result := Meeting(tc.input)
		if result != tc.expected {
			t.Fatalf("Expected %s, but got %s", tc.expected, result)
		}
	}
	t.Log("passed all the tests")

}

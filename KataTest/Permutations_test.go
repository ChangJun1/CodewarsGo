package KataTest

import (
	"CodewarsTraining/RankUp"
	"reflect"
	"testing"
)

func Test_Permutations(t *testing.T) {
	t.Log("start test", t.Name())

	abcd := []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
		"bacd", "badc", "bcad", "bcda", "bdac", "bdca",
		"cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
		"dabc", "dacb", "dbac", "dbca", "dcab", "dcba"}
	testCases := []struct {
		intput   string
		expected []string
	}{

		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},

		{"abcd", abcd},
		{"bcda", abcd},
		{"dcba", abcd},

		{"aa", []string{"aa"}},
		{"aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
		{"aaaab", []string{"aaaab", "aaaba", "aabaa", "abaaa", "baaaa"}},
		{"aaaaab", []string{"aaaaab", "aaaaba", "aaabaa", "aabaaa", "abaaaa", "baaaaa"}},
	}

	for _, tc := range testCases {
		result := RankUp.Permutations(tc.intput)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Fatalf("Expected %+v, but got %+v", tc.expected, result)
		}
	}
	t.Log("passed all the tests")

}

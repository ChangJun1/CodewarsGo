package Fundamental

import (
	"strings"
)

func BandNameGenerator(word string) string {
	if len(word) == 0 {
		return ""
	}
	if word[0] == word[len(word)-1] {
		return strings.Title(word) + word[1:]
	}
	return "The " + strings.Title(word)
}

/*func bandNameGenerator(word string) string {
	first := word[:1]
	last := word[len(word)-1:]
	if first != last {
		return "The " + strings.Title(word)
	}
	return strings.Title(word) + word[1:]
}*/

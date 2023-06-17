package Fundamental

import "strings"

/*Your task is to convert strings to how they would be written by Jaden Smith. The strings are actual quotes from Jaden Smith, but they are not capitalized in the same way he originally typed them.

Example:

Not Jaden-Cased: "How can mirrors be real if our eyes aren't real"
Jaden-Cased:     "How Can Mirrors Be Real If Our Eyes Aren't Real"
*/

func ToJadenCase(str string) string {
	return strings.Title(str)

	// strings.Title不稳定，https://pkg.go.dev/strings#Title
	/*words := strings.Split(str, " ")

	for i, word := range words {
		if len(word) == 0 {
			continue
		}

		w := strings.ToUpper(string(word[0]))
		words[i] = w
		if len(word) == 1 {
			continue
		}

		words[i] = w + word[1:]
	}

	return strings.Join(words, " ")*/
}

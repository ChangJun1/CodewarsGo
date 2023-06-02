package Fundamental

import (
	"fmt"
	"strconv"
	"strings"
)

/*Description:
Encrypt this!

You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:

Your message is a string containing space separated words.
You need to encrypt each word in the message using the following rules:
The first letter must be converted to its ASCII code.
The second letter must be switched with the last letter
Keepin' it simple: There are no special characters in the input.
Examples:
EncryptThis("Hello") == "72olle"
EncryptThis("good") == "103doo"
EncryptThis("hello world") == "104olle 119drlo"*/

func EncryptThis(text string) string {
	if len(text) == 0 {
		return ""
	}
	strSlice := strings.Split(text, " ")
	var outStrSlice []string
	for _, v := range strSlice {
		newStr := strconv.Itoa(int(v[0]))
		if len(v) == 2 {
			newStr += string(v[1])
		} else if len(v) >= 3 {
			newStr = fmt.Sprintf("%s%s%s%s", newStr, string(v[len(v)-1]), v[2:len(v)-1], string(v[1]))
		}
		outStrSlice = append(outStrSlice, newStr)
	}
	return strings.Join(outStrSlice, " ")
}

/*func EncryptThis(text string) string {
	// Implement me! :)
	res := ""
	for _, word := range strings.Split(text, " ") {
		for i, letter := range word {
			switch i {
			case 0:
				res += string(strconv.Itoa(int(letter)))
			case 1:
				res += string(word[len(word)-1])
			case len(word) - 1:
				res += string(word[1])
			default:
				res += string(letter)
			}
		}
		res += " "
	}
	return strings.TrimSpace(res)
}
*/

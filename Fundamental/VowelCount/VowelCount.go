package main

import "fmt"

/*
Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.

*/

func GetCount(str string) (count int) {
	// Enter solution here
	for _, v := range str {
		switch v {
		case 'a':
			fallthrough
		case 'e':
			fallthrough
		case 'i':
			fallthrough
		case 'o':
			fallthrough
		case 'u':
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(GetCount("abracadabra"))
}

package Fundamental

import (
	"fmt"
)

/*Modify the spacify function so that it returns the given string with spaces inserted between each character.

spacify=("hello world") -> # returns "h e l l o   w o r l d*/

func Spacify(s string) string {
	res := ""
	for i, v := range s {
		if i != len(s)-1 {
			res += fmt.Sprintf("%c ", v)
		} else {
			res += fmt.Sprintf("%c", v)
		}
	}
	return res
}

/*func Spacify(s string) string {
	return strings.Join(strings.Split(s, ""), " ")
}

func Spacify(s string) string {
	d := ""
	var sb strings.Builder
	for _, c := range s {
		sb.WriteString(d)
		sb.WriteRune(c)
		d = " "
	}
	return sb.String()
}*/

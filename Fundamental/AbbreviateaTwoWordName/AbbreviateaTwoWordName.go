package main

import (
	"fmt"
	"strings"
)

/*
Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them. //两位首字母大写，以点号连接

It should look like this:

Sam Harris => S.H

patrick feeney => P.F
*/

func AbbrevName(name string) string {
	//your code here
	a := strings.Split(name, " ")
	if len(a) < 2 {
		return ""
	}
	return fmt.Sprintf("%s.%s", strings.ToUpper(string(a[0][0])), strings.ToUpper(string(a[1][0])))
}

func main() {
	fmt.Println(AbbrevName("Sam Harris"))
	fmt.Println(AbbrevName("patrick feeneys"))
}

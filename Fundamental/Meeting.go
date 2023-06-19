package Fundamental

import (
	"fmt"
	"sort"
	"strings"
)

/*John has invited some friends. His list is:

s = "Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill";
Could you make a program that

makes this string uppercase
gives it sorted in alphabetical order by last name.
When the last names are the same, sort them by first name. Last name and first name of a guest come in the result between parentheses separated by a comma.

So the result of function meeting(s) will be:

"(CORWILL, ALFRED)(CORWILL, FRED)(CORWILL, RAPHAEL)(CORWILL, WILFRED)(TORNBULL, BARNEY)(TORNBULL, BETTY)(TORNBULL, BJON)"
It can happen that in two distinct families with the same family name two people have the same first name too.*/

func Meeting(s string) string {
	s = strings.ToUpper(s)
	strSli := strings.Split(s, ";")
	sort.Slice(strSli, func(i, j int) bool {
		s1Sli := strings.Split(strSli[i], ":")
		s2Sli := strings.Split(strSli[j], ":")
		if s1Sli[1] != s2Sli[1] {
			return s1Sli[1] < s2Sli[1]
		}
		return s1Sli[0] < s2Sli[0]
	})
	res := ""
	for _, v := range strSli {
		vSli := strings.Split(v, ":")
		res += fmt.Sprintf("(%s, %s)", vSli[1], vSli[0])
	}
	return res
}

/*func Meeting(s string) string {
	sw := strings.Split(strings.ToUpper(s), ";")
	res := []string{}
	for _, pn := range sw {
		a := strings.Split(pn, ":")
		s := "(" + a[1] + ", " + a[0] + ")"
		res = append(res, s)
	}
	sort.Strings(res)
	return strings.Join(res, "")
}*/

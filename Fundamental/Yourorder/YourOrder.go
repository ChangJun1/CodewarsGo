package main

/*
Your task is to sort a given string. Each word in the string will contain a single number. This number is the position the word should have in the result.

Note: Numbers can be from 1 to 9. So 1 will be the first word (not 0).

If the input string is empty, return an empty string. The words in the input String will only contain valid consecutive numbers.

Examples
"is2 Thi1s T4est 3a"  -->  "Thi1s is2 3a T4est"
"4of Fo1r pe6ople g3ood th5e the2"  -->  "Fo1r the2 g3ood 4of th5e pe6ople"
""  -->  ""
*/

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

//func Order(sentence string) string {
//	if len(sentence) == 0 {
//		return ""
//	}
//
//	arr := strings.Split(sentence, " ")
//	if len(arr) == 1 {
//		i := findNumber(arr[0])
//		if i == -1 {
//			return ""
//		}
//		return sentence
//	}
//
//	arrNum := make([]int, len(arr))
//	for i, v := range arr {
//		arrNum[i] = findNumber(v)
//	}
//
//	for i := 0; i < len(arr); i++ {
//		for j := i + 1; j < len(arr); j++ {
//			if arrNum[i] > arrNum[j] {
//				arrNum[i], arrNum[j] = arrNum[j], arrNum[i]
//				arr[i], arr[j] = arr[j], arr[i]
//			}
//		}
//	}
//
//	s := strings.Join(arr, " ")
//	return s
//}

func findNumber(s string) int {
	for _, v := range s {
		if v-'0' > 0 && v-'0' < 10 {
			return int(v)
		}
	}
	return -1
}

func Order(sentence string) string {
	if strings.TrimSpace(sentence) == "" {
		return sentence
	}

	re := regexp.MustCompile("[1-9]")
	pieces := strings.Split(sentence, " ")
	sort.SliceStable(pieces, func(i, j int) bool {
		num_i := re.FindString(pieces[i])
		num_j := re.FindString(pieces[j])
		return num_i < num_j
	})
	return strings.Join(pieces, " ")
}

/*func Order(s string) string {
	a := strings.Split(s," ")
	r := make([]string,len(a))
	for _,st := range a {
		for _,c := range st {
			if c>='0' && c<='9' {
				r[int(c-'1')]=st
				break
			}
		}
	}
	return strings.Join(r," ")
}*/

func main() {
	fmt.Println(Order("is2 Thi1s T4est 3a") == "Thi1s is2 3a T4est")
}

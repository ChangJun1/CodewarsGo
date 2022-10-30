package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"
*/

func HighAndLow(in string) string {
	// Code here
	a := strings.Split(in, " ")
	var min, max int
	if len(a) > 0 {
		min, _ = strconv.Atoi(a[0])
		max = min
	}
	for _, v := range a {
		t, _ := strconv.Atoi(v)
		if t < min {
			min = t
		}
		if t > max {
			max = t
		}
	}
	return fmt.Sprintf("%d %d", max, min)
}

func main() {
	fmt.Println(HighAndLow("1 2 3"))
	fmt.Println(HighAndLow("1 2 3 4 5"))
	fmt.Println(HighAndLow("1 2 -3 4 5"))
	fmt.Println(HighAndLow("1 9 3 4 -5"))
}

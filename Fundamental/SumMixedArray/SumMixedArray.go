package main

import (
	"fmt"
	"strconv"
)

/*Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.

Return your answer as a number.*/

func SumMix(arr []any) int {
	var res int
	for _, v := range arr {
		if t, ok := v.(int); ok {
			res += t
		} else if t, ok := v.(string); ok {
			r, _ := strconv.Atoi(t)
			res += r
		} else {
			continue
		}
	}
	return res
}

func main() {
	a := []any{9, 3, "7", "3"}
	res := SumMix(a)
	fmt.Println(res)
}

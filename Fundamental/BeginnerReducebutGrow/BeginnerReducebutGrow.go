package main

import "fmt"

//Given a non-empty array of integers, return the result of multiplying the values together in order. Example:
//
//[1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24

func Grow(arr []int) int {
	var res int = 1
	for _, v := range arr {
		res *= v
	}
	return res
}

func main() {
	fmt.Println(Grow([]int{1, 2, 3}))
	fmt.Println(Grow([]int{4, 1, 1, 1, 4}))
}

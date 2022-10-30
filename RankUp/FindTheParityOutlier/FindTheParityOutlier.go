package main

import "fmt"

/*
	You are given an array (which will have a length of at least 3, but could be very large) containing integers.
	The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N.
	Write a method that takes the array as an argument and returns this "outlier" N.
*/

/*Examples
[2, 4, 0, 100, 4, 11, 2602, 36]
Should return: 11 (the only odd number)

[160, 3, 1719, 19, 11, 13, -21]
Should return: 160 (the only even number)*/

func FindOutlier(integers []int) int {
	var res int
	if len(integers) < 3 {
		return -1
	}
	if IsOdd(integers[0]) && IsOdd(integers[1]) {
		for i := 2; i < len(integers); i++ {
			if !IsOdd(integers[i]) {
				res = integers[i]
				break
			}
		}
	}

	if !IsOdd(integers[0]) && !IsOdd(integers[1]) {
		for i := 2; i < len(integers); i++ {
			if IsOdd(integers[i]) {
				res = integers[i]
				break
			}
		}
	}

	if !IsOdd(integers[0]) && IsOdd(integers[1]) {
		if IsOdd(integers[2]) {
			res = integers[0]
		} else {
			res = integers[1]
		}
	}

	if IsOdd(integers[0]) && !IsOdd(integers[1]) {
		if IsOdd(integers[2]) {
			res = integers[1]
		} else {
			res = integers[0]
		}
	}

	return res
}

func IsOdd(n int) bool {
	return n%2 == 0
}

func GetGrade(a, b, c int) rune {
	// Code here
	average := (a + b + c) / 30
	switch average {
	case 10:
		fallthrough
	case 9:
		return 'A'
	case 8:
		return 'B'
	case 7:
		return 'C'
	case 6:
		return 'D'
	default:
		return 'F'
	}
}

func main() {
	fmt.Println(FindOutlier([]int{2, 6, 8, -10, 3}))
}

package main

import "fmt"

/*
Count the number of divisors of a positive integer n.

Random tests go up to n = 500000.

Examples (input --> output)
4 --> 3 (1, 2, 4)
5 --> 2 (1, 5)
12 --> 6 (1, 2, 3, 4, 6, 12)
30 --> 8 (1, 2, 3, 5, 6, 10, 15, 30)
Note you should only return a number, the count of divisors. The numbers between parentheses are shown only for you to see which numbers are counted in each case.
*/

func Divisors(n int) int {
	//Put your code here
	if n <= 0 {
		return -1
	}
	if n == 1 || n == 2 {
		return n
	}
	var res int = 2
	t := n
	for j := 2; j < t; j++ {
		if n%j == 0 {
			t = n / j
			if t == j {
				res++
			} else {
				res += 2
			}
		}
	}
	return res
}

//func Divisors(n int)int{
//	//Put your code here
//	if n <= 0 {
//		return -1
//	}
//	var res int
//	for i := 1; i <= n; i++ {
//		if n%i == 0 {
//			res++
//		}
//	}
//	return res
//}

func main() {
	res := Divisors(54)
	fmt.Println(res)
}

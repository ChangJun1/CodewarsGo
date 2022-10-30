package main

import "fmt"

/*Write a program that will calculate the number of trailing zeros in a factorial of a given number.

N! = 1 * 2 * 3 *  ... * N

Be careful 1000! has 2568 digits...

For more info, see: http://mathworld.wolfram.com/Factorial.html

Examples
zeros(6) = 1
# 6! = 1 * 2 * 3 * 4 * 5 * 6 = 720 --> 1 trailing zero

zeros(12) = 2
# 12! = 479001600 --> 2 trailing zeros*/

func Zeros(n int) int {
	// your code here
	// 10 = 2 * 5
	// so we need to find how many couples of (2,5)
	if n <= 0 {
		return 0
	}
	var a, b int
	for i := 1; i <= n; i++ {
		a += HowMany(i, 2)
		b += HowMany(i, 5)
	}
	if a <= b {
		return a
	}
	return b

}

func HowMany(n, d int) int {
	var res int
	for n >= d {
		if n%d == 0 {
			res++
			n /= d
		} else {
			return res
		}
	}
	return res
}

func main() {
	fmt.Println(Zeros(30))
}

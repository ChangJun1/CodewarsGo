package main

import "fmt"

func Persistence(n int) int {
	// your code
	if n <= 0 {
		// input error
		return -1
	} else if n < 10 {
		return 0
	} else {
		var cnt int
		for n >= 10 {
			var a []int
			for n >= 10 {
				m := n % 10
				n = n / 10
				a = append(a, m)
			}
			for _, v := range a {
				n *= v
			}
			cnt++
		}
		return cnt
	}
	return 0
}

func main() {
	n := 39
	fmt.Println(n)
	n = Persistence(n)
	fmt.Println(n)
}

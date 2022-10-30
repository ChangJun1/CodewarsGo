package main

import "fmt"

func BouncingBall(h, bounce, window float64) int {
	// your code
	if h < 0 || bounce <= 0 || bounce >= 1 || window < 0 || window >= h {
		return -1
	}
	res := 1
	h *= bounce
	for h > window {
		res += 2
		h *= bounce
	}
	return res
}

func main() {
	fmt.Println(BouncingBall(3, 0.66, 1.5) == 3)
}

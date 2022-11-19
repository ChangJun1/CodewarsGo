package main

import "fmt"

/*
Let's pretend your company just hired your friend from college and paid you a referral bonus. Awesome! To celebrate,
you're taking your team out to the terrible dive bar next door and using the referral bonus to buy, and build,
the largest three-dimensional beer can pyramid you can. And then probably drink those beers, because let's pretend it's Friday too.

A beer can pyramid will square the number of cans in each level - 1 can in the top level, 4 in the second, 9 in the next, 16, 25...

Complete the beeramid function to return the number of complete levels of a beer can pyramid you can make, given the parameters of:

your referral bonus, and

the price of a beer can

For example:

beeramid(1500, 2); // should === 12
beeramid(5000, 3); // should === 16
MATHEMATICSALGORITHMS

*/

// 平方和公式
// n(n+1)(2n+1)/6

func Beeramid(bonus int, price float64) int {
	// your code here
	var res int
	for i := 1; price*float64(i*(i+1)*(2*i+1)/6) <= float64(bonus); i++ {
		res = i
	}
	return res
}

func main() {
	fmt.Println(Beeramid(9, 2.0))  // to equal 1
	fmt.Println(Beeramid(21, 1.5)) // To(Equal(3))
}

package main

import (
	"fmt"
)

/*
Your function takes two arguments:

current father's age (years)
current age of his son (years)

Сalculate how many years ago the father was twice as old as his son (or in how many years he will be twice as old).
The answer is always greater or equal to 0, no matter if it was in the past or it is in the future.

FUNDAMENTALS MATHEMATICS
*/

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	// Implement me
	if dadYearsOld < 0 || sonYearsOld < 0 || dadYearsOld <= sonYearsOld {
		return -1
	}
	var res int
	if dadYearsOld == 2*sonYearsOld {
		return 0
	}
	for dadYearsOld > 2*sonYearsOld {
		dadYearsOld--
		res++
		if dadYearsOld == 2*sonYearsOld {
			return res
		}
	}

	for dadYearsOld > sonYearsOld && dadYearsOld < 2*sonYearsOld {
		dadYearsOld++
		res++
		if dadYearsOld == 2*sonYearsOld {
			return res
		}
	}
	return 0
}

//func TwiceAsOld(dadYearsOld int, sonYearsOld int) int {
//	return int(math.Abs(float64(dadYearsOld - (sonYearsOld * 2))))
//}

func main() {
	res := TwiceAsOld(55, 30) // To(Equal(5))
	fmt.Println(res)

}

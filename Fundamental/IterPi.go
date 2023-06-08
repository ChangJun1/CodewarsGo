package Fundamental

import (
	"fmt"
	"math"
)

func IterPi(epsilon float64) (int, string) {
	pi := math.Pi
	sum := 0.0
	i := 0
	for ; math.Abs(4*sum-pi) >= epsilon; i++ {
		sum += math.Pow(-1.0, float64(i)) / float64(2*i+1)
	}
	return i, fmt.Sprintf("%.10f", 4*sum)
}

/*func IterPi(epsilon float64) (int, string) {
	n, v := 1.0, 1.0
	for math.Abs(math.Pi - 4*v) > epsilon {
		v += math.Pow(-1, n)/(2*n + 1)
		n += 1
	}
	return int(n), fmt.Sprintf("%.10f", 4*v)
}*/

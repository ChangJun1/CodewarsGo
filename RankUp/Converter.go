package RankUp

import (
	"fmt"
	"math"
)

/*I started this as a joke among friends, telling that converting numbers to other integer bases is for n00bs, while an actual coder at least converts numbers to more complex bases like pi (or π or however you wish to spell it in your language), so they dared me proving I was better.

And I did it in few hours, discovering that what I started as a joke actually has some math ground and application (particularly the conversion to base pi, it seems).

That said, now I am daring you to do the same, that is to build a function so that it takes a number (any number, you are warned!) and optionally the number of decimals (default: 0) and a base (default: pi), returning the proper conversion as a string:

#Note In Java, Rust and Go there is no easy way with optional parameters so all three parameters will be given; the same in C# because, as of now, the used version is not known.

converter(13.0, 0, math.Pi) return 103.0
converter(13.0, 3, math.Pi) return 103.010
converter(-13.0, 0, 2.0) return -1101.0
I know most of the world uses a comma as a decimal mark, but as English language and culture are de facto the Esperanto of us coders, we will stick to our common glorious traditions and uses, adopting the trivial dot (".") as decimal separator; if the absolute value of the result is <1, you have of course to put one (and only one) leading 0 before the decimal separator.

Finally, you may assume that decimals if provided will always be >= 0 and that no test base will be smaller than 2 (because, you know, converting to base 1 is pretty lame) or greater than 36; as usual, for digits greater than 9 you can use uppercase alphabet letter, so your base of numeration is going to be: '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ'.

That is my first 3-languages-kata, so I count on you all to give me extensive feedback, no matter how harsh it may sound, so to improve myself even further :)*/

func Converter(n float64, decimals int, base float64) string {
	var res string

	if n < 0 {
		res += "-"
		n = -n
	}

	if n == 0 {
		res += "0"
	}

	// 确定整数位数
	b := 0.0
	for n >= math.Pow(base, b) {
		b++
	}

	// base进制下的最大数字
	x := math.Floor(base)
	if x == base {
		x--
	}

	// 整数部分
	for i := b - 1; i >= 0; i-- {
		tmp := math.Pow(base, i)
		if tmp > n {
			res += "0"
			continue
		}

		tmp2 := n / base
		if tmp2 <= x && math.Floor(tmp2) == tmp2 {
			res += i2str(int(tmp2))
			n -= tmp2 * base
			continue
		}

		j := findNumAndSubN(x, tmp, &n)
		res += i2str(j)
	}

	// 没有小数位直接返回
	if decimals == 0 {
		return res
	}

	res += "."

	// 小数部分
	for i := 1; i <= decimals; i++ {
		tmp := math.Pow(base, float64(-i))
		if n == 0 || tmp > n {
			res += "0"
			continue
		}

		j := findNumAndSubN(x, tmp, &n)
		res += i2str(j)
	}

	return res
}

func i2str(i int) string {
	if i < 10 {
		return fmt.Sprintf("%d", i)
	}
	return fmt.Sprintf("%c", 'A'+i-10)
}

func findNumAndSubN(x, y float64, n *float64) int {
	j := x
	for ; j >= 0; j-- {
		if j*y <= *n {
			break
		}
	}
	*n -= j * y
	return int(j)
}

/*const DIGIT = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Converter(n float64, decimals int, base float64) string {
	var r []byte
	if n < 0 {
		r = append(r, '-')
		n = -n
	}
	d := 0
	if n > 1 {
		d = int(math.Log(n) / math.Log(base))
	}
	for i := d; i >= -decimals; i-- {
		z := math.Pow(base, float64(i))
		if i == -1 {
			r = append(r, '.')
		}
		x := math.Floor(n / z)
		r = append(r, DIGIT[int(x)])
		n -= x * z
	}
	return string(r)
}*/

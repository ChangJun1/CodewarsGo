package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*The rgb function is incomplete.

Complete it so that passing in RGB decimal values will result in a hexadecimal representation being returned.

Valid decimal values for RGB are 0 - 255. Any values that fall out of that range must be rounded to the closest valid value.

Note: Your answer should always be 6 characters long, the shorthand with 3 will not work here.

The following are examples of expected output values:

kata.rgb(255, 255, 255) -- returns FFFFFF
kata.rgb(255, 255, 300) -- returns FFFFFF
kata.rgb(0, 0, 0) -- returns 000000
kata.rgb(148, 0, 211) -- returns 9400D3

*/

func RGB(r, g, b int) string {
	// Your code here
	if r < 0 {
		r = 0
	}
	if g < 0 {
		g = 0
	}
	if b < 0 {
		b = 0
	}
	if r > 255 {
		r = 255
	}
	if g > 255 {
		g = 255
	}
	if b > 255 {
		b = 255
	}

	s1 := strconv.FormatInt(int64(r), 16)
	s2 := strconv.FormatInt(int64(g), 16)
	s3 := strconv.FormatInt(int64(b), 16)

	if len(s1) == 1 {
		s1 = "0" + s1
	}
	if len(s2) == 1 {
		s2 = "0" + s2
	}
	if len(s3) == 1 {
		s3 = "0" + s3
	}
	
	return strings.ToUpper(s1 + s2 + s3)
}

func main() {
	fmt.Println(RGB(255, 255, 255))
	fmt.Println(RGB(0, 1, 2))
}

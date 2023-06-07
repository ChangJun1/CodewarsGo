package RankUp

import (
	"strconv"
	"strings"
)

/*Implement a function that receives two IPv4 addresses, and returns the number of addresses between them (including the first one, excluding the last one).

All inputs will be valid IPv4 addresses in the form of strings. The last address will always be greater than the first one.

Examples
* With input "10.0.0.0", "10.0.0.50"  => return   50
* With input "10.0.0.0", "10.0.1.0"   => return  256
* With input "20.0.0.10", "20.0.1.0"  => return  246
*/

func IpsBetween(start, end string) int {
	s := strings.Split(start, ".")
	e := strings.Split(end, ".")
	s0, _ := strconv.Atoi(s[0])
	s1, _ := strconv.Atoi(s[1])
	s2, _ := strconv.Atoi(s[2])
	s3, _ := strconv.Atoi(s[3])
	e0, _ := strconv.Atoi(e[0])
	e1, _ := strconv.Atoi(e[1])
	e2, _ := strconv.Atoi(e[2])
	e3, _ := strconv.Atoi(e[3])
	res := (e0-s0)*256*256*256 + (e1-s1)*256*256 + (e2-s2)*256 + e3 - s3
	return res
}

/*func IpsBetween(start, end string) int {
	p, q := 0, 0
	for _, s := range strings.Split(start, ".") {
		v, _ := strconv.Atoi(s)
		p = p*256 + v
	}
	for _, s := range strings.Split(end, ".") {
		v, _ := strconv.Atoi(s)
		q = q*256 + v
	}
	return q - p
}*/

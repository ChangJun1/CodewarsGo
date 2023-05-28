package Fundamental

func isVowel(b byte) bool {
	switch b {
	case 'a':
		fallthrough
	case 'e':
		fallthrough
	case 'i':
		fallthrough
	case 'o':
		fallthrough
	case 'u':
		return true
	}
	return false
}

func LongestVowelChain(s string) int {
	res := 0
	var j int
	for i := 0; i < len(s); i++ {
		j = i
		for j < len(s) && isVowel(s[j]) {
			j++
		}
		if j == i {
			continue
		}
		length := j - i
		if length > res {
			res = length
		}
		i = j
	}
	return res
}

// most popular
/*func Solve(s string) int {
	n, max := 0, 0
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			if n++; n > max {
				max = n
			}
		default:
			n = 0
		}
	}
	return max
}*/

// 利用正则表达式
/*func Solve(s string) int {
	max := 0
	for _, sub := range MustCompile(`[aeiou]+`).FindAllString(s, -1) {
		if len(sub) > max {
			max = len(sub)
		}
	}
	return max
}*/

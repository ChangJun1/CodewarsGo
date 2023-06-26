package RankUp

import "sort"

/*In this kata, your task is to create all permutations of a non-empty input string and remove duplicates, if present.

Create as many "shufflings" as you can!

Examples:

With input 'a':
Your function should return: ['a']

With input 'ab':
Your function should return ['ab', 'ba']

With input 'abc':
Your function should return ['abc','acb','bac','bca','cab','cba']

With input 'aabb':
Your function should return ['aabb', 'abab', 'abba', 'baab', 'baba', 'bbaa']
Note: The order of the permutations doesn't matter.

Good luck!*/

func Permutations(s string) []string {
	// your code here
	if len(s) == 0 {
		return nil
	}

	if len(s) == 1 {
		return []string{s}
	}

	res := []string{}
	for i, v := range s {
		nextStr := s[:i] + s[i+1:]
		nextSlice := Permutations(nextStr)
		for j, str := range nextSlice {
			newStr := string(v) + str
			nextSlice[j] = newStr
		}

		res = append(res, nextSlice...)
	}

	res = removeMoreThanTwice(res)

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func removeMoreThanTwice(strs []string) []string {
	if len(strs) == 0 {
		return nil
	}
	if len(strs) == 1 {
		return strs
	}
	m := make(map[string]bool)
	result := []string{}
	for _, s := range strs {
		if _, ok := m[s]; !ok {
			m[s] = true
			result = append(result, s)
		}
	}
	return result
}

// the most popular
/*func Permutations(s string) (a []string) {
	if len(s) < 2 {
		return []string{s}
	}
	m := map[string]bool{}
	for _, sub := range Permutations(s[1:]) {
		for i := 0; i <= len(sub); i++ {
			st := sub[0:i] + s[0:1] + sub[i:]
			if m[st] {
				continue
			}
			m[st] = true
			a = append(a, st)
		}
	}
	return
}*/

package Fundamental

import (
	"fmt"
)

/*Input:

a string strng
an array of strings arr
Output of function contain_all_rots(strng, arr) (or containAllRots or contain-all-rots):

a boolean true if all rotations of strng are included in arr
false otherwise
Examples:
contain_all_rots(
"bsjq", ["bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"]) -> true

contain_all_rots(
"Ajylvpy", ["Ajylvpy", "ylvpyAj", "jylvpyA", "lvpyAjy", "pyAjylv", "vpyAjyl", "ipywee"]) -> false)
Note:
Though not correct in a mathematical sense

we will consider that there are no rotations of strng == ""
and for any array arr: contain_all_rots("", arr) --> true*/

func ContainAllRots(strng string, arr []string) bool {
	if strng == "" {
		return true
	}
	m := make(map[string]bool)
	for _, v := range arr {
		m[v] = true
	}
	for i, _ := range strng {
		tmp := fmt.Sprintf("%s%s", strng[i:], strng[:i])
		if _, ok := m[tmp]; ok {
			continue
		} else {
			return false
		}
	}
	return true
}

/*func ContainAllRots(strng string, arr []string) bool {
	counter := 0
	master := []string{}
	for i,_ := range strng{
		master = append(master,strng[i:]+strng[:i])
	}
	for _,j := range master{
		for _,k := range arr{
			if j == k{
				counter ++
				break
			}
		}
	}
	return len(strng) == counter
}*/

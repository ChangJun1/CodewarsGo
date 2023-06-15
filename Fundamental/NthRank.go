package Fundamental

import (
	"sort"
	"strings"
)

/*To participate in a prize draw each one gives his/her firstname.

Each letter of a firstname has a value which is its rank in the English alphabet. A and a have rank 1, B and b rank 2 and so on.

The length of the firstname is added to the sum of these ranks hence a number som.

An array of random weights is linked to the firstnames and each som is multiplied by its corresponding weight to get what they call a winning number.

Example:
names: "COLIN,AMANDBA,AMANDAB,CAROL,PauL,JOSEPH"
weights: [1, 4, 4, 5, 2, 1]

PauL -> som = length of firstname + 16 + 1 + 21 + 12 = 4 + 50 -> 54
The *weight* associated with PauL is 2 so PauL's *winning number* is 54 * 2 = 108.
Now one can sort the firstnames in decreasing order of the winning numbers. When two people have the same winning number sort them alphabetically by their firstnames.

Task:
parameters: st a string of firstnames, we an array of weights, n a rank

return: the firstname of the participant whose rank is n (ranks are numbered from 1)

Example:
names: "COLIN,AMANDBA,AMANDAB,CAROL,PauL,JOSEPH"
weights: [1, 4, 4, 5, 2, 1]
n: 4

The function should return: "PauL"
Notes:
The weight array is at least as long as the number of names, it may be longer.

If st is empty return "No participants".

If n is greater than the number of participants then return "Not enough participants".

See Examples Test Cases for more examples.*/

func NthRank(st string, we []int, n int) string {
	if len(st) == 0 {
		return "No participants"
	}
	stArr := strings.Split(st, ",")
	if n > len(stArr) {
		return "Not enough participants"
	}
	winNums := make([][2]int, len(stArr))
	for i, s := range stArr {
		som := len(s)
		for _, c := range s {
			if c >= 'a' && c <= 'z' {
				som += int(c-'a') + 1
			} else if c >= 'A' && c <= 'Z' {
				som += int(c-'A') + 1
			}
		}
		winNums[i][0] = som * we[i]
		winNums[i][1] = i
	}
	sort.Slice(winNums, func(i, j int) bool {
		if winNums[i][0] != winNums[j][0] {
			return winNums[i][0] > winNums[j][0]
		} else {
			return strings.Compare(stArr[winNums[i][1]], stArr[winNums[j][1]]) == -1
		}
	})
	return stArr[winNums[n-1][1]]
}

/*func NthRank(st string, we []int, n int) string {
	if st == "" {
		return "No participants"
	}

	nombres := strings.Split(st, ",")

	if n > len(nombres) {
		return "Not enough participants"
	}

	var par []participantes

	for i, e := range nombres {
		par = append(par, participantes{e, getWining(strings.ToLower(e), we[i])})
	}

	sort.SliceStable(par, func(i, j int) bool {
		if par[j].numero == par[i].numero {
			return strings.ToLower(par[j].nombre) > strings.ToLower(par[i].nombre)
		}
		return par[j].numero < par[i].numero
	})

	return par[n-1].nombre
}

func getWining(str string, weight int) int {
	sArray := strings.Split(str, "")
	n := 0

	for _, e := range sArray {
		n += int(e[0]) - 96
	}

	return (n + len(sArray)) * weight
}*/

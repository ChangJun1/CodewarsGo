package main

import "fmt"

func main() {
	s := CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	fmt.Println(s)

}

func CreatePhoneNumber(numbers [10]uint) string {
	m := map[uint]string{
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
		0: "0"}
	var s string = "("
	for i, v := range numbers {
		if i == 3 {
			s += ") "
		}
		if i == 6 {
			s += "-"
		}
		s += m[v]
	}
	return s
}

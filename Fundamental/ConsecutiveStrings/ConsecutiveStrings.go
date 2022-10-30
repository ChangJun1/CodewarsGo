package main

import (
	"fmt"
	"strings"
)

/*
You are given an array(list) strarr of strings and an integer k. Your task is to return the first longest string consisting of k consecutive strings taken in the array.

Examples:
strarr = ["tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"], k = 2

Concatenate the consecutive strings of strarr by 2, we get:

treefoling   (length 10)  concatenation of strarr[0] and strarr[1]
folingtrashy ("      12)  concatenation of strarr[1] and strarr[2]
trashyblue   ("      10)  concatenation of strarr[2] and strarr[3]
blueabcdef   ("      10)  concatenation of strarr[3] and strarr[4]
abcdefuvwxyz ("      12)  concatenation of strarr[4] and strarr[5]

Two strings are the longest: "folingtrashy" and "abcdefuvwxyz".
The first that came is "folingtrashy" so
longest_consec(strarr, 2) should return "folingtrashy".

In the same way:
longest_consec(["zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"], 2) --> "abigailtheta"
n being the length of the string array, if n = 0 or k > n or k <= 0 return "" (return Nothing in Elm, "nothing" in Erlang).

Note
consecutive strings : follow one after another without an interruption

*/

func LongestConsec(strarr []string, k int) string {
	// your code
	n := len(strarr)
	if n == 0 || k > n || k <= 0 {
		return ""
	}
	var res string
	var max int
	for i := 0; i <= n-k; i++ {
		var tmp string
		for j := i; j < i+k; j++ {
			tmp += strarr[j]
		}
		l := len(tmp)
		if l > max {
			max = l
			res = tmp
		}
	}
	return res
}

// best solution

func LongestConsec2(strarr []string, k int) string {
	var buffer string
	var largest string

	for i := 0; i <= len(strarr)-k; i++ {
		buffer = strings.Join(strarr[i : i+k][:], "")
		if len(buffer) > len(largest) {
			largest = buffer
		}
	}
	return largest
}

func main() {
	strarr := []string{"zone", "abigail", "theta", "form", "libe", "zas"}
	k := 2
	res := LongestConsec(strarr, k)
	fmt.Println(res == "abigailtheta")

}

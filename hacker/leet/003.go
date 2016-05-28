package main

import (
	"fmt"
	"os"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)

	maxv := 0
	last := 0

	for i, c := range s {

		v, ok := m[c]
		if ok {
			if last < v+1 {
				last = v + 1
			}
		}

		if maxv < i-last+1 {
			maxv = i - last + 1
		}
		m[c] = i
	}
	return maxv
}

func main() {

	fmt.Println(lengthOfLongestSubstring("abcabcab"))
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 3						  //
////////////////////////////////////////////////////

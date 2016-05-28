package main

import (
	"fmt"
	"os"
)

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, c := range nums {
		v, ok := m[target-c]
		if ok {
			return []int{v, i}
		}
		m[c] = i
	}
	return nums
}

func main() {

	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// [0 1]					  //
////////////////////////////////////////////////////

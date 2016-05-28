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

func intMin(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	cnt := len(nums1) + len(nums2)
	if cnt%2 == 1 {
		return float64(findNth(nums1, nums2, cnt/2+1))
	} else {
		return float64((findNth(nums1, nums2, cnt/2) + findNth(nums1, nums2, cnt/2+1))) / 2.0
	}
}

func findNth(nums1 []int, nums2 []int, nth int) int {
	if len(nums1) > len(nums2) {
		return findNth(nums2, nums1, nth)
	}
	if len(nums1) == 0 {
		return nums2[nth-1]
	}
	if len(nums2) == 0 {
		return nums1[nth-1]
	}

	if nth == 1 {
		return intMin(nums1[0], nums2[0])
	}

	pa := intMin(len(nums1), nth/2)
	pb := nth - pa
	if nums1[pa-1] < nums2[pb-1] {
		return findNth(nums1[pa:], nums2, nth-pa)
	} else {
		return findNth(nums1, nums2[pb:], nth-pb)
	}

}

func main() {

	fmt.Printf("%f\n", findMedianSortedArrays([]int{1, 3, 5, 7}, []int{2, 4, 6}))
	fmt.Printf("%f\n", findMedianSortedArrays([]int{1, 3, 5}, []int{2, 4, 6}))
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 3						  //
////////////////////////////////////////////////////

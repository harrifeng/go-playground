package main

import (
	"fmt"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Display() {
	if l == nil {
		fmt.Println("Void ListNode")
	}

	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
	fmt.Println()
}

func createListFromArray(nums []int) *ListNode {
	dummy := &ListNode{
		Val: -1,
	}

	cur := dummy

	for _, c := range nums {
		cur.Next = &ListNode{
			Val: c,
		}
		cur = cur.Next
	}
	return dummy.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -1,
	}
	cur := dummy
	_ = cur
	adv := 0

	for l1 != nil || l2 != nil || adv > 0 {
		cnt := adv
		if l1 != nil {
			cnt += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			cnt += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{
			Val: cnt % 10,
		}
		cur = cur.Next
		adv = cnt / 10
	}
	return dummy.Next
}

func main() {

	l1 := createListFromArray([]int{2, 4, 3})
	l2 := createListFromArray([]int{5, 6, 4})
	l1.Display()
	l2.Display()
	r := addTwoNumbers(l1, l2)
	r.Display()
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 2 4 3 					  //
// 5 6 4 					  //
// 7 0 8 					  //
////////////////////////////////////////////////////

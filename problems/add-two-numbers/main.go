package main

import (
	"fmt"
	"strings"
)

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	current := node
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return node.Next
}

func main() {
	l1 := fromArray([]int{2, 9})
	l2 := fromArray([]int{5, 9, 9, 1})

	ll := addTwoNumbers(l1, l2)
	fmt.Println(format(ll))
}

func fromArray(xs []int) *ListNode {
	node := &ListNode{}
	current := node

	for _, x := range xs {
		current.Next = &ListNode{Val: x}
		current = current.Next
	}
	return node.Next
}

func format(head *ListNode) string {
	var b strings.Builder
	b.WriteString("[")
	for head != nil {
		b.WriteString(fmt.Sprintf("%d", head.Val))
		if head.Next != nil {
			b.WriteString(",")
		}
		head = head.Next
	}
	b.WriteString("]")
	return b.String()
}

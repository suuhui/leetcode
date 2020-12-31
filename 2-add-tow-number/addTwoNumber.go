package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := &ListNode{Val: 1, Next:nil}
	l.Next = &ListNode{Val: 2, Next:nil}
	l.Next.Next = &ListNode{Val: 3, Next:nil}

	r := &ListNode{Val: 3, Next:nil}
	r.Next = &ListNode{Val: 8, Next:nil}
	r.Next.Next = &ListNode{Val: 6, Next:nil}

	t := addTwoNumber(l, r)
	t.print()
}

func addTwoNumber(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	end := result
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		l1Value, l2Value := 0, 0
		if l1 != nil {
			l1Value = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Value = l2.Val
			l2 = l2.Next
		}
		nodeVal := l1Value + l2Value + carry
		carry = 0
		if nodeVal >= 10 {
			nodeVal -= 10
			carry = 1
		}
		node := &ListNode{Val: nodeVal, Next: nil}
		end.Next = node
		end = node
	}
	return result.Next
}

func (ln *ListNode) print() {
	l := ln
	for l != nil {
		fmt.Printf("%d", l.Val)
		l = l.Next
	}
	fmt.Println()
}

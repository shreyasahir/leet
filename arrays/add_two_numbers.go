package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type linklist struct {
	head *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	p := l1
	q := l2
	carry := 0
	curr := head
	var x, y int

	for p != nil || q != nil {
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum := x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return head.Next
}

func (l *linklist) addNode(val int) {
	n := &ListNode{Val: val}
	if l.head == nil {
		l.head = n
	} else {
		curr := l.head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = n
	}
}

func (l *linklist) printList() {
	curr := l.head

	for curr != nil {
		fmt.Println("curr node", curr.Val)
		curr = curr.Next
	}
}

func main() {
	l1 := &linklist{}
	l1.addNode(8)
	l1.addNode(1)
	//l1.addNode(2)
	l1.printList()

	l2 := &linklist{}
	l2.addNode(0)
	// l2.addNode(6)
	// l2.addNode(5)
	l2.printList()
	l3 := addTwoNumbers(l1.head, l2.head)

	for l3 != nil {
		fmt.Println("l3", l3.Val)
		l3 = l3.Next
	}
}

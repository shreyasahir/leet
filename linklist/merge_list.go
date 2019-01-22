package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//ListNode is used for storing node info.
type ListNode struct {
	Val  int
	Next *ListNode
}

type linklist struct {
	head *ListNode
}

func (l *linklist) addNode(v int) {
	n := &ListNode{Val: v}
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

func (l *linklist) printNode() {
	curr := l.head
	for curr != nil {
		fmt.Println("Curr node", curr.Val)
		curr = curr.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2

}

func main() {
	l1 := &linklist{}
	l1.addNode(1)
	l1.addNode(2)
	l1.addNode(4)
	l1.printNode()

	l2 := &linklist{}
	l2.addNode(1)
	l2.addNode(3)
	l2.addNode(4)
	l2.printNode()

}

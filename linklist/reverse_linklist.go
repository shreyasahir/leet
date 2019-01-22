/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import (
	"fmt"
)

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

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next

	}
	return prev
}

func main() {
	l1 := &linklist{}
	l1.addNode(1)
	l1.addNode(2)
	l1.addNode(3)
	l1.addNode(4)
	l1.addNode(5)
	l1.printNode()
}

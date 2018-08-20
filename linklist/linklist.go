package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

type linklist struct {
	head *node
}

func (l *linklist) addNode(v int) {
	n := &node{val: v}
	if l.head == nil {
		l.head = n
	} else {
		curr := l.head

		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}
}

func (l *linklist) printNode() {
	curr := l.head
	for curr != nil {
		fmt.Println("Curr node", curr.val)
		curr = curr.next
	}
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

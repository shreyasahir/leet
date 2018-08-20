package main

import (
	"fmt"
)

//TreeNode is struct to represent  binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Tree pointer to root node
type Tree struct {
	root *TreeNode
}

func main() {

	// [3,9,20,null,null,15,7]
	//[-8,-6,7,6,null,null,null,null,5]
	t := Tree{}
	t.root = &TreeNode{Val: 1}
	t.root.Left = &TreeNode{Val: 2}
	t.root.Right = &TreeNode{Val: 3}

	t1 := Tree{}
	t1.root = &TreeNode{Val: 1}
	t1.root.Left = &TreeNode{Val: 2}
	t1.root.Right = &TreeNode{Val: 3}
	//fmt.Println(maxDepth(t.root))
	fmt.Println(isSameTree(t.root, t1.root))
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	return false

}

//https://leetcode.com/problems/maximum-binary-tree/

package main

import (
	"fmt"
)

// * type TreeNode struct {
// 	*     Val int
// 	*     Left *TreeNode
// 	*     Right *TreeNode
// 	* }

func main() {
	arr := []int{3, 2, 1, 6, 0, 5}
	max, k := sortArray(arr)
	left := arr[:k]
	right := arr[k+1:]
	fmt.Println("max", max)
	fmt.Println("left", left)
	fmt.Println("right", right)
}

func constructMaximumBinaryTree(nums []int) *TreeNode {

	var tnode TreeNode
	var maxk int

	if len(nums) == 0 {
		return nil
	}

	maxk = findMaxIndex(nums)
	tnode.Val = nums[maxk]
	if nums[:maxk] != nil {
		tnode.Left = constructMaximumBinaryTree(nums[:maxk])
	} else {
		tnode.Left = nil
	}

	if nums[maxk+1:] != nil {
		tnode.Right = constructMaximumBinaryTree(nums[maxk+1:])
	} else {
		tnode.Right = nil
	}

	return &tnode
}

func findMaxIndex(slice []int) int {
	index := 0
	for i, j := range slice {
		if j > slice[index] {
			index = i
		}
	}
	return index
}

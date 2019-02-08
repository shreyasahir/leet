// https://leetcode.com/problems/range-sum-of-bst/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
	result []int
)

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	result = nil
	var sum int
	inOrder(root)
	//fmt.Println("result",result)
	if len(result) > 0 {
		for i := 0; i < len(result); i++ {
			if result[i] >= L && result[i] <= R {
				sum += result[i]
			}
		}
	}

	return sum
}

func inOrder(n *TreeNode) {
	if n != nil {
		inOrder(n.Left)
		//f(n)
		result = append(result, n.Val)
		inOrder(n.Right)
	}
}
package leetcode530

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	result, pre := math.MaxInt32, -1

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if pre != -1 && node.Val-pre < result {
			result = node.Val - pre
		}

		pre = node.Val

		dfs(node.Right)
	}

	dfs(root)

	return result
}

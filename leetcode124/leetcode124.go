package leetcode124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum, maxSubSum := search(root)
	return max(maxSum, maxSubSum)
}

func search(root *TreeNode) (maxSum, maxSubSum int) {
	maxSum = math.MinInt32
	maxSubSum = math.MinInt32

	if root == nil {
		return
	}

	leftMaxSum, leftMaxSubSum := search(root.Left)
	maxSum = max(maxSum, leftMaxSum)
	maxSum = max(maxSum, leftMaxSubSum)
	maxSubSum = max(maxSubSum, root.Val+leftMaxSubSum)

	rightMaxSum, rightMaxSubSum := search(root.Right)
	maxSum = max(maxSum, rightMaxSum)
	maxSum = max(maxSum, rightMaxSubSum)
	maxSubSum = max(maxSubSum, root.Val+rightMaxSubSum)

	maxSum = max(maxSum, root.Val+leftMaxSubSum+rightMaxSubSum)
	maxSubSum = max(maxSubSum, root.Val)

	return
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

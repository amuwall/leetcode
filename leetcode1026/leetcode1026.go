package leetcode1026

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	result := searchTree(root, root.Val, root.Val)
	return result
}

func searchTree(root *TreeNode, minVal, maxVal int) (maxDiff int) {
	if root == nil {
		return 0
	}

	maxDiff = max(abs(root.Val-minVal), abs(root.Val-maxVal))
	minVal = min(root.Val, minVal)
	maxVal = max(root.Val, maxVal)
	maxDiff = max(maxDiff, searchTree(root.Left, minVal, maxVal))
	maxDiff = max(maxDiff, searchTree(root.Right, minVal, maxVal))
	return
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

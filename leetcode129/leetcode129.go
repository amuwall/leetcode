package leetcode129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return calc(root, 0)
}

func calc(root *TreeNode, v int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val + v*10
	}

	v = root.Val + v*10
	return calc(root.Left, v) + calc(root.Right, v)
}

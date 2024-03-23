package leetcode101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return leftTreeSymmetricRightTree(root.Left, root.Right)
}

func leftTreeSymmetricRightTree(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	if !leftTreeSymmetricRightTree(left.Left, right.Right) {
		return false
	}
	if !leftTreeSymmetricRightTree(left.Right, right.Left) {
		return false
	}

	return true
}

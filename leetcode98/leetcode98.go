package leetcode98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	pre := 0
	preOk := false

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !dfs(node.Left) {
			return false
		}

		if preOk && pre >= node.Val {
			return false
		}

		pre = node.Val
		preOk = true

		if !dfs(node.Right) {
			return false
		}

		return true
	}

	return dfs(root)
}

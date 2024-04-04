package leetcode106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			root.Left = buildTree(inorder[0:i], postorder[0:i])
			root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
			break
		}
	}

	return root
}

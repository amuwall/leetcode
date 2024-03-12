package leetcode1261

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	vals map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	elements := FindElements{
		vals: map[int]bool{},
	}

	root.Val = 0
	dfsTreeNode(root, elements.vals)

	return elements
}

func dfsTreeNode(root *TreeNode, vals map[int]bool) {
	if root == nil {
		return
	}

	vals[root.Val] = true

	if root.Left != nil {
		root.Left.Val = root.Val*2 + 1
		dfsTreeNode(root.Left, vals)
	}
	if root.Right != nil {
		root.Right.Val = root.Val*2 + 2
		dfsTreeNode(root.Right, vals)
	}
}

func (this *FindElements) Find(target int) bool {
	_, ok := this.vals[target]
	return ok
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */

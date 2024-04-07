package leetcode114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	p := root
	for p != nil {
		if p.Left != nil {
			next := p.Left
			q := next
			for q.Right != nil {
				q = q.Right
			}
			q.Right = p.Right
			p.Left, p.Right = nil, next
		}
		p = p.Right
	}
}

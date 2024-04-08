package leetcode102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	currentNodes := []*TreeNode{root}
	nextNodes := []*TreeNode{}
	for len(currentNodes) > 0 {
		level := []int{}
		for _, node := range currentNodes {
			level = append(level, node.Val)
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}

		result = append(result, level)

		currentNodes = nextNodes
		nextNodes = []*TreeNode{}
	}

	return result
}

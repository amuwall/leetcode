package leetcode103

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	currentNodes := []*TreeNode{root}
	nextNodes := []*TreeNode{}
	asc := true
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

		if !asc {
			slices.Reverse(level)
		}

		result = append(result, level)

		currentNodes = nextNodes
		nextNodes = []*TreeNode{}
		asc = !asc
	}

	return result
}

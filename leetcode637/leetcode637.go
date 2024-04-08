package leetcode637

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	if root == nil {
		return result
	}

	currentNodes := []*TreeNode{root}
	nextNodes := []*TreeNode{}
	for len(currentNodes) > 0 {
		sum := 0
		for _, currentNode := range currentNodes {
			if currentNode.Left != nil {
				nextNodes = append(nextNodes, currentNode.Left)
			}
			if currentNode.Right != nil {
				nextNodes = append(nextNodes, currentNode.Right)
			}

			sum += currentNode.Val
		}

		result = append(result, float64(sum)/float64(len(currentNodes)))

		currentNodes = nextNodes
		nextNodes = []*TreeNode{}
	}

	return result
}

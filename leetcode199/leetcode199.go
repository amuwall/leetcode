package leetcode199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	var currentNodes []*TreeNode
	var nextNodes []*TreeNode

	currentNodes = append(currentNodes, root)
	for len(currentNodes) > 0 {
		for _, node := range currentNodes {
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}

		result = append(result, currentNodes[len(currentNodes)-1].Val)

		currentNodes = nextNodes
		nextNodes = []*TreeNode{}
	}

	return result
}

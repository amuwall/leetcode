package leetcode501

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var currentVal, currentCount, maxCount int
	var modes []int

	inOrderTraversal(root, &currentVal, &currentCount, &maxCount, &modes)

	return modes
}

func inOrderTraversal(node *TreeNode, currentVal, currentCount, maxCount *int, modes *[]int) {
	if node == nil {
		return
	}

	inOrderTraversal(node.Left, currentVal, currentCount, maxCount, modes)

	if node.Val == *currentVal {
		*currentCount++
	} else {
		*currentVal = node.Val
		*currentCount = 1
	}

	if *currentCount == *maxCount {
		*modes = append(*modes, *currentVal)
	} else if *currentCount > *maxCount {
		*maxCount = *currentCount
		*modes = []int{*currentVal}
	}

	inOrderTraversal(node.Right, currentVal, currentCount, maxCount, modes)
}

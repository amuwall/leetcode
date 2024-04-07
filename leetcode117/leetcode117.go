package leetcode117

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	currentNodes := []*Node{}
	nextNodes := []*Node{}

	currentNodes = append(currentNodes, root)
	for len(currentNodes) > 0 {
		for i := 0; i < len(currentNodes); i++ {
			if i != len(currentNodes)-1 {
				currentNodes[i].Next = currentNodes[i+1]
			}

			if currentNodes[i].Left != nil {
				nextNodes = append(nextNodes, currentNodes[i].Left)
			}
			if currentNodes[i].Right != nil {
				nextNodes = append(nextNodes, currentNodes[i].Right)
			}
		}

		currentNodes = nextNodes
		nextNodes = []*Node{}
	}

	return root
}

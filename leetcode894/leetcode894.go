package leetcode894

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{{}}
	}

	var result []*TreeNode
	for i := 1; i < n-1; i += 2 {
		leftFBTs := allPossibleFBT(i)
		rightFBTs := allPossibleFBT(n - 1 - i)

		for j := range leftFBTs {
			for k := range rightFBTs {
				result = append(result, &TreeNode{
					Left:  leftFBTs[j],
					Right: rightFBTs[k],
				})
			}
		}
	}

	return result
}

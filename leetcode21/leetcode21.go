package leetcode21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var next *ListNode

	p := list1
	q := list2
	for {
		if p == nil && q == nil {
			break
		}

		var node *ListNode
		if p == nil {
			node = &ListNode{Val: q.Val}
			q = q.Next
		} else if q == nil {
			node = &ListNode{Val: p.Val}
			p = p.Next
		} else if p.Val < q.Val {
			node = &ListNode{Val: p.Val}
			p = p.Next
		} else {
			node = &ListNode{Val: q.Val}
			q = q.Next
		}

		if head == nil {
			head = node
			next = node
		} else {
			next.Next = node
			next = node
		}
	}

	return head
}

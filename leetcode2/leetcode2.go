package leetcode2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	l1Node := l1
	l2Node := l2
	additionCarry := 0

	for {
		node := &ListNode{}

		if l1Node != nil {
			node.Val += l1Node.Val
			l1Node = l1Node.Next
		}
		if l2Node != nil {
			node.Val += l2Node.Val
			l2Node = l2Node.Next
		}
		node.Val += additionCarry
		additionCarry = 0

		if node.Val >= 10 {
			additionCarry = node.Val / 10
			node.Val = node.Val % 10
		}

		tail.Next = node
		tail = tail.Next

		if l1Node == nil && l2Node == nil && additionCarry == 0 {
			break
		}
	}

	return head.Next
}

package leetcode141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next
	slow := head

	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		if fast == slow {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}
}

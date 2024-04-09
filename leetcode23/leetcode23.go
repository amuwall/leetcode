package leetcode23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	p := lists[0]
	for i := 1; i < len(lists); i++ {
		p = mergeLists(p, lists[i])
	}

	return p
}

func mergeLists(list1, list2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}

	return head.Next
}

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists_digui(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists_digui(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists_digui(l1, l2.Next)
		return l2
	}
}

func mergeTwoLists_diedai(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 1}
	pre := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 == nil {
		pre.Next = l2
	} else if l2 == nil {
		pre.Next = l1
	}
	return head.Next
}

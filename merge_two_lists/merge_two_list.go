package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var min, max *ListNode
	if l1.Val < l2.Val {
		min, max = l1, l2
	} else {
		min, max = l2, l1
	}
	var pre *ListNode = min
	min = min.Next
	for min != nil {
		if min.Val <= max.Val {
			pre = min
			min = min.Next
			continue
		}
		pre.Next = max
		max = max.Next
	}
}

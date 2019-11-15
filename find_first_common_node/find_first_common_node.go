package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func findFirstCommonNode(headA *ListNode, headB *ListNode) *ListNode {
	a := headA
	b := headB
	aLen, bLen := 0, 0
	for a != nil {
		aLen++
		a = a.Next
	}
	for b != nil {
		bLen++
		b = b.Next
	}
	if aLen > bLen {
		for i := 0; i < aLen-bLen; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < aLen-bLen; i++ {
			headB = headB.Next
		}
	}
	for headB != nil && headA != nil {
		if headA.Val == headB.Val {
			return headA
		} else {
			headA = headA.Next
			headB = headB.Next
		}
	}
	return nil
}

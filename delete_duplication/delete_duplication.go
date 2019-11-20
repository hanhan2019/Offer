package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplication(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			for head != nil && head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			pre.Next = head.Next
			head = head.Next
		} else {
			pre = head
			head = head.Next
		}
	}
	return dummy.Next
}

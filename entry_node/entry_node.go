package main

import "fmt"

func main() {
	a := &ListNode{0, nil}
	b := &ListNode{1, a}
	head := &ListNode{2, b}
	c := head
	fmt.Println(head)
	c.Next = b.Next
	fmt.Println(head)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func entryNodeOfLoop(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	record := make(map[*ListNode]int)
// 	for head != nil {
// 		if _, ok := record[head]; ok {
// 			return head
// 		} else {
// 			record[head] = 1
// 		}
// 		head = head.Next
// 	}
// 	return nil
// }

func entryNodeOfLoop(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	first := head
	second := head
	for second.Next != nil && second.Next.Next != nil {
		second = second.Next.Next
		first = first.Next
		if second == first {
			first = head
			for first != second {
				first = first.Next
				second = second.Next
			}
			return first
		}
	}
	return nil
}

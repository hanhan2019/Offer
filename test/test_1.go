package main

import "fmt"

func main() {
	type ListNode struct {
		Value int64
		Next  *ListNode
	}
	a := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
			Next: &ListNode{
				Value: 3,
				Next: &ListNode{
					Value: 4,
					Next:  nil,
				},
			},
		},
	}
	p := a
	tempNode := p.Next
	// tempHead := tempNode
	p.Next = tempNode.Next
	p = p.Next
	fmt.Println(a.Value, a.Next)

}

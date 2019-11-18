package main

import (
	"fmt"
	"strings"
)

var PrefixTagSortKey = "v1:ts:"

func main() {
	// type ListNode struct {
	// 	Value int64
	// 	Next  *ListNode
	// }
	// a := &ListNode{
	// 	Value: 1,
	// 	Next: &ListNode{
	// 		Value: 2,
	// 		Next: &ListNode{
	// 			Value: 3,
	// 			Next: &ListNode{
	// 				Value: 4,
	// 				Next:  nil,
	// 			},
	// 		},
	// 	},
	// }
	// p := a
	// tempNode := p.Next
	// // tempHead := tempNode
	// p.Next = tempNode.Next
	// p = p.Next
	// fmt.Println(a.Value, a.Next)
	rewards := strings.Split("jifen888_道具_1_10000/jifen1888_道具_1_5000/jifen2888_道具_1_1000/jifen8888_道具_1_500/jifen18888_道具_1_50/jifen88888_道具_1_10", "/")
	fmt.Println(rewards)
}

//打印链表倒数第n个节点的信息
package main

import "fmt"

type NodeList struct {
	Val  int
	Next *NodeList
}

// just print it form tail to head and do not modify the original NodeList

func main() {

}

func print_list_nth_end_to_start(head *NodeList, n int) {
	p := head
	count := 0
	value := 0
	for p != nil {
		count++
		p = p.Next
	}
	if count < n {
		fmt.Printf("wrong")
	} else {
		for head != nil {
			value++
			if value == count-n+1 {
				fmt.Printf("%v", head.Val)
			} else {
				head = head.Next
			}
		}
	}
}

func better(head *NodeList, n int) *NodeList {
	if head == nil || n <= 0 {
		return nil
	}
	p := head
	r := head
	for n > 1 {
		if r.Next != nil {
			r = r.Next
		}
		n--
	}
	for r.Next != nil {
		p = p.Next
		r = r.Next
	}
	return p
}

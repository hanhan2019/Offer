package main

type NodeList struct {
	Val  int
	Next *NodeList
}

func main() {

}

func reverseList_digui(head *NodeList) *NodeList {
	if head == nil || head.Next == nil {
		return nil
	}
	p := reverseList_digui(head.Next)
	head.Next.Next = head.Next
	head.Next = nil
	return p
}

func reverseList_diedai(head *NodeList) *NodeList {
	var pre *NodeList = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

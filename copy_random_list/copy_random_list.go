package main

func main() {

}

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

func copyRandomList(head *ListNode) *ListNode {
	cloneNodes(head)
	setRandomNodes(head)
	return reconnectNodeList(head)

}

func cloneNodes(head *ListNode) {
	p := head
	for p != nil {
		tempNode := &ListNode{Val: p.Val, Next: p.Next, Random: nil}
		p.Next = tempNode
		p = tempNode.Next
	}
}

func setRandomNodes(head *ListNode) {
	p := head
	for p != nil {
		tempNode := p.Next
		if p.Random != nil {
			tempNode.Random = p.Random.Next
		}
		p = tempNode.Next
	}
}

func reconnectNodeList(head *ListNode) *ListNode {
	p := head
	var tempNode, tempHead *ListNode = nil, nil
	if p != nil {
		tempNode = p.Next
		tempHead = tempNode
		p.Next = tempNode.Next
		p = p.Next
	}
	for p != nil {
		tempNode.Next = p.Next
		tempNode = tempNode.Next
		p.Next = tempNode.Next
		p = p.Next
	}
	return tempHead
}

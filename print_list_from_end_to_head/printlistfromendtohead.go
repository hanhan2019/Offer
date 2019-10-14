//从尾到头打印链表
package main

type NodeList struct {
	Val  int
	Next *NodeList
}

// just print it form tail to head and do not modify the original NodeList

func main() {

}

//reverseList 递归
func reverseList_1(head *NodeList) *NodeList {
	if head == nil || head.Next == nil {
		return head
	}
	endNode := reverseList_1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return endNode
}

//reverseList 迭代
func reverseList_2(head *NodeList) *NodeList {
	var preNode *NodeList = nil
	for head != nil {
		nextNode := head.Next
		head.Next = preNode
		preNode = head
		head = nextNode
	}
	return preNode
}

//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//1<-2 3->4->5->NULL

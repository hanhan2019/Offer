//通过先序和中序推出二叉树
package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {

}

func ReconstructBinaryTree(preList, midList []int) *Node {
	head := &Node{
		Value: preList[0],
	}
	index := 0
	for i, v := range midList {
		if v == preList[0] {
			index = i
		}
	}
	head.Left = ReconstructBinaryTree(preList[1:index+1], midList[:index])
	head.Right = ReconstructBinaryTree(preList[index+1:], midList[index+1:])
	return head
}

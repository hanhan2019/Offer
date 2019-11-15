package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printFromTopToBottom(root *TreeNode) []int {
	a := []int{}
	b := []*TreeNode{}
	if root != nil {
		b = append(b, root)
	}
	for i := 0; i < len(b); i++ {
		a = append(a, b[i].Val)
		if b[i].Left != nil {
			b = append(b, b[i].Left)
		}
		if b[i].Right != nil {
			b = append(b, b[i].Right)
		}
	}
	return a
}

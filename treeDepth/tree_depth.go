package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeDepth(root *TreeNode) int {
	if root != nil {
		left := treeDepth(root.Left)
		right := treeDepth(root.Right)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
	return 0
}

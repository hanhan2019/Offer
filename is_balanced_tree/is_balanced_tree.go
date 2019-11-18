package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if treeDepth(root) >= 0 {
		return true
	}
	return false
}

func treeDepth(root *TreeNode) int {
	if root != nil {
		left := treeDepth(root.Left)
		right := treeDepth(root.Right)
		if left < 0 || right < 0 {
			return -1
		}
		if left-right > 1 || left-right < -1 {
			return -1
		}
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
	return 0
}

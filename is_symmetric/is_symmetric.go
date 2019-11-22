package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == nil && r == nil
	}
	if l.Val != r.Val {
		return false
	}
	return check(l.Left, r.Right) && check(l.Right, r.Left)
}

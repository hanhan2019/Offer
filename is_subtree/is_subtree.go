package main

import "fmt"

func main() {
	fmt.Println(true || false)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return isSubTree(s.Left, t) || isSubTree(s.Right, t)
	} else {
		return equelTwoTree(s, t) || isSubTree(s.Left, t) || isSubTree(s.Right, t)
	}
	return false
}

func equelTwoTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	}
	if s.Val == t.Val {
		return equelTwoTree(s.Right, t.Right) && equelTwoTree(s.Left, t.Left)
	}
	return false
}

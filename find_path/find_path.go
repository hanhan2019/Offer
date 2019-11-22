package main

import "fmt"

func main() {
	root := &TreeNode{}
	root.Val = 6
	root.Left = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	root.Right = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(findPath(root, 13))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func findPath(root *TreeNode, sum int) [][]int {
// 	result := [][]int{}
// 	if root == nil {
// 		return result
// 	}
// 	a := []int{}
// 	count := 0
// 	temp := [][]int{}
// 	tempList := binaryTreePaths(root, a, count, temp)
// 	for _, v := range tempList {
// 		if v[len(v)-1] == sum {
// 			result = append(result, v[:len(v)-1])
// 		}
// 	}
// 	return result
// }

// func binaryTreePaths(root *TreeNode, path []int, count int, pathList [][]int) [][]int {
// 	if root == nil {
// 		return pathList
// 	}
// 	path = append(path, root.Val)
// 	count += root.Val
// 	if root.Left == nil && root.Right == nil {
// 		path = append(path, count)
// 		pathList = append(pathList, path)
// 		return pathList
// 	}
// 	pathList = binaryTreePaths(root.Left, path, count, pathList)
// 	pathList = binaryTreePaths(root.Right, path, count, pathList)
// 	return pathList
// }

func findPath(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	a := []int{}
	temp := [][]int{}
	return binaryTreePaths(root, a, sum, temp)
}

func binaryTreePaths(root *TreeNode, path []int, sum int, pathList [][]int) [][]int {
	fmt.Println(sum)
	if root == nil {
		return pathList
	}
	sum -= root.Val
	// if sum >= 0 {

	// 	fmt.Println(path)
	// }
	if root.Left == nil && root.Right == nil && sum >= 0 {
		path = append(path, root.Val)
		pathList = append(pathList, path)
		return pathList
	}
	pathList = binaryTreePaths(root.Left, path, sum, pathList)
	pathList = binaryTreePaths(root.Right, path, sum, pathList)
	return pathList
}

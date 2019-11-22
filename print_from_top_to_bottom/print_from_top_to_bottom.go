package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	list := make([][]string, 0)
	list = append(list, a)
	fmt.Println("list:", list)
	// 修改a
	a[0] = "xxx"
	fmt.Println("list:", list)
	fmt.Printf("ptr(a)    = %p \n", a)
	fmt.Printf("ptr(list[0]) = %p \n", list[0])

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func printFromTopToBottom(root *TreeNode) [][]int {
// 	if root == nil {
// 		return [][]int{}
// 	}
// 	o := 1
// 	que := []*TreeNode{root}
// 	result := [][]int{}
// 	start, end := 0, 0
// 	for len(que) != 0 {
// 		if o == 1 {
// 			one := []int{}
// 			for _, v := range que {
// 				one = append(one, v.Val)
// 			}
// 			result = append(result, one)
// 			o = 0
// 		}

// 		if que[0].Left != nil {
// 			que = append(que, que[0].Left)
// 		}
// 		if que[0].Right != nil {
// 			que = append(que, que[0].Right)
// 		}
// 	}
// }

func printFromTopToEnd(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	que := []*TreeNode{root}
	result := []int{}
	for len(que) != 0 {
		if que[0].Left != nil {
			que = append(que, que[0].Left)
		}
		if que[0].Right != nil {
			que = append(que, que[0].Right)
		}
		result = append(result, que[0].Val)
		que = que[1:]
	}
	return result
}

func printFromTopToEndByEnter(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	que := []*TreeNode{root}
	left := 1
	count := 0
	result := [][]int{}
	one := []int{}
	for len(que) != 0 {
		if que[0].Left != nil {
			que = append(que, que[0].Left)
			count++
		}
		if que[0].Right != nil {
			que = append(que, que[0].Right)
			count++
		}
		one = append(one, que[0].Val)
		que = que[1:]
		left--
		if left == 0 {
			left = count
			count = 0
			result = append(result, one)
			one = []int{}
		}
	}
	return result
}

func printFromTopToEndByZSize(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	que1, que2 := []*TreeNode{root}, []*TreeNode{}
	result := [][]int{}
	one := []int{}
	reserver := false
	for len(que1) != 0 || len(que2) != 0 {
		if !reserver {
			if que1[len(que1)-1].Left != nil {
				que2 = append(que2, que1[len(que1)-1].Left)
			}
			if que1[len(que1)-1].Right != nil {
				que2 = append(que2, que1[len(que1)-1].Right)
			}
			one = append(one, que1[len(que1)-1].Val)
			que1 = que1[:len(que1)-1]
			if len(que1) == 0 {
				fmt.Println(result)
				one = []int{}
				reserver = !reserver
			}
		} else {
			if que2[len(que2)-1].Right != nil {
				que1 = append(que1, que2[len(que2)-1].Right)
			}
			if que2[len(que2)-1].Left != nil {
				que1 = append(que1, que2[len(que2)-1].Left)
			}
			one = append(one, que2[len(que2)-1].Val)
			que2 = que2[:len(que2)-1]
			if len(que2) == 0 {
				one = []int{}
				reserver = !reserver
			}
		}

	}
	return result
}

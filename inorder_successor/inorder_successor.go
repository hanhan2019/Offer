package main

func main() {

}

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Father *TreeNode
}

// func inorderSuccessor(p *TreeNode) *TreeNode {
// 	if p == nil {
// 		return nil
// 	}
// 	root := p
// 	for root.Father != nil {
// 		root = root.Father
// 	}
// 	list := getMidleSort(root)
// 	for i := 0; i < len(list); i++ {
// 		if list[i] == p && i+1 < len(list) {
// 			return list[i+1]
// 		}
// 	}
// 	return nil
// }

// func getMidleSort(p *TreeNode) []*TreeNode {
// 	if p == nil {
// 		return nil
// 	}
// 	a := []*TreeNode{}
// 	if p.Left != nil {
// 		a = append(a, getMidleSort(p.Left)...)
// 	}
// 	a = append(a, p)
// 	if p.Right != nil {
// 		a = append(a, getMidleSort(p.Right)...)
// 	}
// 	return a
// }

func inorderSuccessor(p *TreeNode) *TreeNode {
	if p == nil {
		return nil
	}
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}
	for p.Father != nil && p.Father.Right == p {
		p = p.Father
	}
	return p.Father
}

// func getMidleFirstSort(p *TreeNode) *TreeNode{
//     if p==nil{
//         return nil
//     }
//     if p.Left!=nil{
//         return getMidleFirstSort(p.Left)
//     }
//         return p
// }

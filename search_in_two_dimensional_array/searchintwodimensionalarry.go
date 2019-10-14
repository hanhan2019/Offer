//二维数组中的查找
package main

import "fmt"

func main() {
	a := 50
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	fmt.Println(BestSolution(matrix, a))
}

func Find(p int, vector [][]int) bool {
	hlen := len(vector)
	llen := len(vector[0])
	var temp int
	if hlen%2 != 0 {
		temp = hlen / 2
	} else {
		temp = hlen/2 - 1
	}
	for i := 0; i < llen; i++ {
		if vector[temp][i] != p {
			continue
		} else {
			return true
		}
	}
	if hlen == 1 {
		return false
	}
	if vector[temp][llen-1] > p {
		return Find(p, vector[0:temp])
	} else if vector[temp][llen-1] < p {
		return Find(p, vector[temp+1:])
	} else {
		return true
	}
}

func Find1(board [][]int, target int) bool {
	rlen := len(board)
	clen := len(board[0])

	//  我们从右上角的元素找起来
	for r, c := 0, clen-1; r < rlen && c >= 0; {
		if board[r][c] == target {
			return true
		}
		if board[r][c] > target {
			c--
			continue
		} else {
			r++
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	l, r := 0, len(matrix)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if matrix[mid][0] > target {
			r = mid - 1
		} else {
			l = mid
		}
	}
	if matrix[l][0] > target {
		return false
	}

	array := matrix[l]
	l, r = 0, len(array)-1
	for l < r {
		mid := (l + r) >> 1
		if array[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return array[l] == target
}

func BestSolution(matrix [][]int, target int) bool {
	hlen := len(matrix)
	if hlen == 0 {
		return false
	}
	llen := len(matrix[0])
	low, high := 0, hlen*llen-1
	for low <= high {
		temp := (low + high) / 2
		tempValue := matrix[temp/llen][temp%llen]
		if tempValue == target {
			return true
		} else if tempValue < target {
			low = temp + 1
		} else {
			high = temp - 1
		}
	}
	return false
}

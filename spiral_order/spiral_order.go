package main

func main() {

}

// func spiralOrder(matrix [][]int) []int {
// 	rowLen := len(matrix[0])
// 	lenLen := len(matrix)
// 	if rowLen == 0 || lenLen == 0 {
// 		return []int{}
// 	}
// 	list := []int{}
// 	for _, v := range matrix[0] {
// 		list = append(list, v)
// 	}
// 	for i := 1; i < lenLen; i++ {
// 		list = append(list, matrix[i][rowLen-1])
// 	}
// 	for i := rowLen - 2; i >= 0; i-- {
// 		list = append(list, matrix[lenLen-1][i])
// 	}
// 	for i := lenLen - 2; i > 0; i-- {
// 		list = append(list, matrix[i][0])
// 	}
// 	nextMatrix := [][]int{}
// 	for i := 1; i < lenLen-1; i++ {
// 		nextMatrix = append(nextMatrix, matrix[i][1:rowLen-1])
// 		// nextMatrix[i-1] = matrix[i][1 : rowLen-1]
// 	}
// 	list = append(list, spiralOrder(matrix)...)
// 	return list
// }

func spiralOrder(matrix [][]int) []int {
	rowLen := len(matrix)
	if rowLen == 0 {
		return []int{}
	}
	lenLen := len(matrix[0])
	if lenLen == 0 {
		return []int{}
	}

	list := []int{}
	rowIndex := 0 //控制行位置
	lenIndex := 0 //控制列位置
	for rowLen != 0 && lenLen != 0 {
		for j := lenIndex; j < lenLen; j++ {
			list = append(list, matrix[rowIndex][j])
		}
		rowIndex++
		for i := rowIndex; i < rowLen; i++ {
			list = append(list, matrix[i][lenLen-1])
		}
		lenLen--
		for j := lenLen - 1; j >= lenIndex; j-- {
			list = append(list, matrix[rowLen-1][j])
		}
		rowLen--
		for i := rowLen - 1; i >= rowIndex; i-- {
			list = append(list, matrix[i][lenIndex])
		}
		lenIndex++
	}
	return list
}

func PrintfMatrix(matrix [][]int, row, colum int) []int {
	if matrix == nil || row <= 0 || colum <= 0 {
		return []int{}
	}
	start := 0
	list := []int{}
	for colum > start*2 && row > start*2 {
		list = append(list, PrintfMatrixInCircle(matrix, row, colum, start)...)
		start++
	}
	return list
}

func PrintfMatrixInCircle(matrix [][]int, row, colum, start int) []int {
	list := []int{}
	endX := colum - 1 - start
	endY := row - 1 - start
	for i := start; i <= endX; i++ {
		list = append(list, matrix[start][i])
	}
	if start < endY {
		for i := start + 1; i <= endY; i++ {
			list = append(list, matrix[i][endX])
		}
	}
	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			list = append(list, matrix[endY][i])
		}
	}
	if start < endX && start < endY-1 {
		for i := endY - 1; i > start; i-- {
			list = append(list, matrix[i][start])
		}
	}
	return list
}

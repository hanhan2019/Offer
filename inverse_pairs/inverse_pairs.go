//在数组中的两个数字如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
// 输入一个数组，求出这个数组中的逆序对的总数。
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 0}
	fmt.Println(mergeSort(a))
}

func mergeSort(nums []int) ([]int, int) {
	if len(nums) < 2 {
		return nums, 0
	}
	i := len(nums) / 2
	left, leftRes := mergeSort(nums[0:i])
	right, rightRes := mergeSort(nums[i:])
	result, res := merge(left, right)
	fmt.Println(result, leftRes, rightRes, res)
	return result, leftRes + rightRes + res
}

func merge(left, right []int) ([]int, int) {
	res := 0
	result := []int{}
	m, n, l, r := 0, 0, len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			res += len(left[m:])
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result, res
}

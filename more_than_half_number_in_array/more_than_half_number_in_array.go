//数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。如果不存在则输出0。
package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := best(a)
	fmt.Println(b)
}

func moreThanHalfNum_Solution(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	max := 0
	result := 0
	for i, j := range m {
		if j > max {
			max = j
			result = i
		}
	}
	if !checkMoreThanHalf(nums, result) {
		result = 0
	}
	return result
}

func best(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	value := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			value = nums[i]
			count = 1
		} else if nums[i] == value {
			count++
		} else {
			count--
		}
	}
	if !checkMoreThanHalf(nums, value) {
		value = 0
	}
	return value
}

func checkMoreThanHalf(nums []int, value int) bool {
	count := uint(0)
	for _, v := range nums {
		if v == value {
			count++
		}
	}
	if (1 << count) > len(nums) {
		return true
	}
	return false
}

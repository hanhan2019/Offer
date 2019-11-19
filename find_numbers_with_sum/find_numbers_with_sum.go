package main

import "sort"

func main() {

}

func findNumbersWithSum(nums []int, target int) []int {
	a := []int{}
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] == target {
			a = append(a, nums[i], nums[j])
			return a
		}
		if nums[i]+nums[j] < target {
			i++
		}
		if nums[i]+nums[j] > target {
			j--
		}
	}
	return a
}

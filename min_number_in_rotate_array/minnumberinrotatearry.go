package main

func main() {

}

func minNumberInRotateArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return nums[i]
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			return nums[i]
		}
	}
	return nums[0]
}

//利用变形的二分查找
func BestAnwer(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	} else {
		midIndex := len(nums) / 2
		if nums[0] < nums[midIndex] {
			return BestAnwer(nums[midIndex+1:])
		} else {
			return BestAnwer(nums[0 : midIndex+1])
		}
	}
}

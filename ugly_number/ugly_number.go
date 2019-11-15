package main

import "fmt"

func main() {
	fmt.Println(getUglyNumber(5))
}

// 1\2\3\5
func getUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	uglyNumber := []int{1}
	a, b, c := 0, 0, 0
	for len(uglyNumber) < n {
		for uglyNumber[a]*2 <= uglyNumber[len(uglyNumber)-1] {
			a++
		}
		for uglyNumber[b]*3 <= uglyNumber[len(uglyNumber)-1] {
			b++
		}
		for uglyNumber[c]*5 <= uglyNumber[len(uglyNumber)-1] {
			c++
		}
		fmt.Println(a, b, c)
		min := Min(uglyNumber[a]*2, uglyNumber[b]*3, uglyNumber[c]*5)
		uglyNumber = append(uglyNumber, min)
	}
	return uglyNumber[len(uglyNumber)-1]
}
func Min(a, b, c int) (min int) {
	if a > b {
		min = b
	} else {
		min = a
	}
	if min > c {
		min = c
	}
	return
}
func CheckIsUgly(a int64) bool {
	for a%2 == 0 {
		a /= 2
	}
	for a%3 == 0 {
		a /= 3
	}
	for a%5 == 0 {
		a /= 5
	}
	if a == 1 {
		return true
	}
	return false
}

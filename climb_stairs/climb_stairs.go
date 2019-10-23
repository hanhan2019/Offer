//爬楼梯，一次1层或者2层
package main

import "fmt"

func main() {
	fmt.Println(climbStairs_n_times(4))
}

//最少次
func climbStairs_min_times(n int) int {
	if n <= 0 {
		return 0
	}
	min := n
	list := []int{1, 2}
	for _, v := range list {
		value := climbStairs_min_times(n-v) + 1
		if value < min {
			min = value
		}
	}
	return min
}

//多少种
func climbStairs_times(n int) int {
	record := make(map[int]int)
	record[1] = 1
	record[2] = 1 + 1
	for i := 1; i <= n; i++ {
		if _, ok := record[i]; !ok {
			record[i] = record[i-1] + record[i-2]
		}
	}
	return record[n]
}

//爬楼梯，一次n层
func climbStairs_n_times(n int) int {
	return 2 << (uint64(n) - 2)
}
